package service

import (
	"errors"
	"github.com/sxz799/surveyX/model/common/response"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"strings"
	"time"
)

type SurveyService struct {
}

var questionService QuestionService

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

	db.Debug().Count(&total)
	db = db.Limit(limit).Offset(offset)
	err := db.Debug().Order("id DESC").Find(&surveys).Error
	return response.PageResult{
		List:     surveys,
		Total:    total,
		PageNum:  pi.PageNum,
		PageSize: pi.PageSize}, err
}

func (ts *SurveyService) Add(s entity.Survey) (err error) {
	err = utils.DB.Debug().Create(&s).Error
	return
}

func (ts *SurveyService) Update(s entity.Survey) error {
	err := utils.DB.Debug().Where("id=?", s.Id).Updates(&s).Error
	return err
}

func (ts *SurveyService) Del(id int) (err error) {
	s := entity.Survey{
		Id: id,
	}
	_ = questionService.DelBySurveyId(id)
	err = utils.DB.Delete(&s).Error
	return
}

func (ts *SurveyService) Get(id int) (s entity.Survey, err error) {
	s.Id = id
	err = utils.DB.First(&s).Error
	return
}

func (ts *SurveyService) Import(file *multipart.FileHeader) (err error) {
	// 1. 读取文件
	r, err := file.Open()
	if err != nil {
		return
	}
	// 2. 解析文件
	filename := file.Filename
	if len(filename) < 4 || filename[len(filename)-4:] != "xlsx" {
		err = errors.New("文件格式错误")
		return
	}
	excelFile, err := excelize.OpenReader(r)
	if err != nil {
		err = errors.New("文件读取失败")
		return
	}
	sheets := excelFile.GetSheetList()
	rows, _ := excelFile.GetRows(sheets[0])
	if len(rows) < 5 {
		err = errors.New("问卷格式错误")
		return
	}
	rows = rows[1:]
	surveyRow := rows[1]
	var dict = map[string]string{
		"是":     "yes",
		"否":     "no",
		"更新":    "yes_but_update",
		"浏览器指纹": "finger",
		"联系方式":  "contact",
		"无":     "",
		"单选题":   "radio",
		"多选题":   "checkbox",
		"简答题":   "text",
	}
	// 2.1 读取问卷信息
	startTime, _ := time.Parse("20060102150405", surveyRow[3])
	endTime, _ := time.Parse("20060102150405", surveyRow[4])
	survey := entity.Survey{
		Title:       surveyRow[0],
		Description: surveyRow[1],
		Status:      dict[surveyRow[2]],
		StartTime:   startTime,
		EndTime:     endTime,
		NeedContact: dict[surveyRow[5]],
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
			var value string
			var hasExtMsg string
			if strings.Contains(row[i], "[填写备注]") {
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
	err = utils.DB.Debug().Create(&survey).Error
	for _, question := range questions {
		question.SurveyId = survey.Id
		_ = questionService.Add(question)
	}

	return
}
