package service

import (
	"errors"
	"mime/multipart"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
	"github.com/xuri/excelize/v2"
)

type SurveyService struct{}

func (ts *SurveyService) List(s entity.SurveySearch) (response.PageResult, error) {
	var surveys []entity.Survey
	var total int64
	pi := s.PageInfo
	limit := pi.PageSize
	offset := pi.PageSize * (pi.PageNum - 1)
	db := utils.DB.Model(&entity.Survey{})

	survey := s.Survey

	if survey.Title != "" {
		db = db.Where("title like ?", "%"+survey.Title+"%")
	}
	if survey.Status != "" {
		db = db.Where("status = ?", survey.Status)
	}
	//问卷开始时间不为空
	if !survey.StartTime.IsZero() {
		db = db.Where("start_time < ?", survey.StartTime)
	}
	//问卷结束时间不为空
	if !survey.EndTime.IsZero() {
		db = db.Where("end_time > ?", survey.EndTime)
	}

	db = db.Where("user_id = ?", survey.UserId)

	db.Count(&total)
	db = db.Limit(limit).Offset(offset)
	err := db.Order("id DESC").Find(&surveys).Error
	return response.PageResult{
		List:     surveys,
		Total:    total,
		PageNum:  pi.PageNum,
		PageSize: pi.PageSize}, err
}

func (ts *SurveyService) Add(s entity.Survey) (err error) {
	s.Id = uuid.New().String()
	s.Status = "new"
	err = utils.DB.Create(&s).Error
	return
}

func (ts *SurveyService) Update(s entity.Survey) error {
	err := utils.DB.Where("id=?", s.Id).Updates(&s).Error
	return err
}

func (ts *SurveyService) Del(id string) (err error) {
	s := entity.Survey{
		Id: id,
	}
	_ = questionService.DelBySurveyId(id)
	_ = answerService.DelBySurveyId(id)
	err = utils.DB.Delete(&s).Error
	return
}

func (ts *SurveyService) Get(id string) (s entity.Survey, err error) {
	s.Id = id
	err = utils.DB.First(&s).Error
	return
}

func (ts *SurveyService) Analysis(id string) (any, error) {
	var result struct {
		KeyCount      int
		QuestionCount int
		FingerCount   int
		ContactCount  int
		FirstCreateAt string
		LastCreateAt  string
	}
	err := utils.DB.Table("answers").
		Select("COUNT(DISTINCT key) as key_count, COUNT(DISTINCT question_id) as question_count, COUNT(DISTINCT finger) as finger_count, COUNT(DISTINCT contact) as contact_count, min(create_at) as first_create_at, max(create_at) as last_create_at").
		Where("survey_id = ?", id).
		Scan(&result).Error
	//时间转为字符串
	if result.FirstCreateAt != "" {
		result.FirstCreateAt = result.FirstCreateAt[:19]
	}
	if result.LastCreateAt != "" {
		result.LastCreateAt = result.LastCreateAt[:19]
	}
	return result, err
}

func (ts *SurveyService) Import(userId int, file *multipart.FileHeader) (err error) {
	// 1. 读取文件
	r, err := file.Open()
	if err != nil {
		return
	}
	defer r.Close()
	// 2. 解析文件
	filename := file.Filename
	if len(filename) < 5 || strings.Contains(filename, ".xlsx") == false {
		err = errors.New("文件格式错误")
		return
	}
	excelFile, err := excelize.OpenReader(r)
	if err != nil {
		err = errors.New("文件读取失败")
		return
	}
	defer excelFile.Close()
	sheets := excelFile.GetSheetList()
	rows, _ := excelFile.GetRows(sheets[0])
	if len(rows) < 5 {
		err = errors.New("问卷格式错误")
		return
	}
	rows = rows[1:]
	surveyRow := rows[1]
	var dict = map[string]string{
		"手机号":    "phone",
		"邮箱":     "email",
		"手机号或邮箱": "phone|email",
		"不限制":    "other",
		"是":      "yes",
		"否":      "no",
		"更新":     "update",
		"浏览器指纹":  "finger",
		"联系方式":   "contact",
		"无":      "",
		"单选":     "radio",
		"多选":     "checkbox",
		"简答":     "text",
	}
	// 2.1 读取问卷信息
	startTime, _ := time.Parse("20060102150405", surveyRow[2])
	endTime, _ := time.Parse("20060102150405", surveyRow[3])
	survey := entity.Survey{
		Id:          uuid.New().String(),
		Title:       surveyRow[0],
		Description: surveyRow[1],
		Status:      "new",
		StartTime:   startTime,
		EndTime:     endTime,
		NeedContact: dict[surveyRow[4]],
		ContactType: dict[surveyRow[5]],
		Repeat:      dict[surveyRow[6]],
		RepeatCheck: dict[surveyRow[7]],
		WaterMark:   surveyRow[8],
	}
	var questions []entity.Question
	questionRows := rows[3:]
	for k, row := range questionRows {
		var ops []entity.Option
		for i := 3; i < len(row); i++ {
			if row[i] == "" {
				break
			}
			value := row[i]
			var hasExtMsg string
			if strings.Contains(value, "[填写备注]") {
				value = strings.ReplaceAll(row[i], "[填写备注]", "")
				hasExtMsg = "yes"
			} else {
				hasExtMsg = "no"
			}
			op := entity.Option{
				Label:     string(rune('A' + i - 3)),
				Value:     value,
				HasExtMsg: hasExtMsg,
			}
			ops = append(ops, op)
		}
		question := entity.Question{
			Text:    row[0],
			Type:    dict[row[1]],
			Order:   k + 1,
			Options: ops,
		}
		questions = append(questions, question)
	}

	survey.UserId = userId
	utils.DB.Create(&survey)
	for _, question := range questions {
		question.SurveyId = survey.Id
		if err := questionService.Add(question); err != nil {
			return err
		}
	}

	return
}
