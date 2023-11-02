<template>
  <el-row class="survey-container">
    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <h2 class="survey-title">{{ survey.title }}</h2>
      <h4 class="survey-description">{{ survey.description }}</h4>
      <el-form ref="answersRef" :inline-message="true" :model="form" :rules="rules" label-width="0">
        <el-form-item label="" prop="contact">
          <el-input v-model="form.contact" placeholder="请输入联系方式"></el-input>
        </el-form-item>
        <el-form-item :inline="false" v-for="(question , index) in survey.questions" :prop="'answers.' + question.id ">
          <div class="question-container">

            <!--题目-->
            <el-tag v-if="question.type === 'radio'">{{ index + 1 }} [单选]</el-tag>
            <el-tag type="warning" v-if="question.type === 'checkbox'">{{ index + 1 }} [多选]</el-tag>
            <el-tag type="success" v-if="question.type === 'text'">{{ index + 1 }} [简答题]</el-tag>
            <el-text>{{ question.text }}[{{ question.id }}]</el-text>
            <br>
            <!--简答题-->
            <el-input type="textarea" :autosize="{ minRows: 2, maxRows: 5 }" v-model="form.answers[question.id]" v-if="question.type==='text'"
                      placeholder=""/>
            <!--单选题-->
            <el-radio-group v-model="form.answers[question.id]" v-if="question.type==='radio'">
              <div style="display: flex;align-items: center" v-for="option in question.options">
                <div style="max-width: 70%;">
                  <el-radio :label=option.label>
                    {{ option.label }} {{ option.value }}
                  </el-radio>
                </div>
                <div>
                  <el-input v-model="option['extMsg']" size="small" type="textarea" :autosize="{ minRows: 1, maxRows: 3 }"
                            v-if="form.answers[question.id] && form.answers[question.id][0] === option.label &&  option.has_ext_msg === 'Y'"
                            placeholder="请在此处填写补充信息！"/>
                </div>
              </div>
            </el-radio-group>
            <!--多选题-->
            <el-checkbox-group v-model="form.answers[question.id]" v-if="question.type==='checkbox'">
              <div style="display: flex;align-items: center" v-for="option in question.options">
                <div style="max-width: 70%;">
                  <el-checkbox :label=option.label>
                    {{ option.label }} {{ option.value }}
                  </el-checkbox>
                </div>
                <div>
                  <el-input v-model="option['extMsg']" size="small" type="textarea" :autosize="{ minRows: 1, maxRows: 3 }"
                            v-if="form.answers[question.id] && form.answers[question.id].toString().indexOf(option.label)>-1 && option.has_ext_msg==='Y'"
                            placeholder="请在此处填写补充信息！"/>
                </div>
              </div>
            </el-checkbox-group>

          </div>
        </el-form-item>
      </el-form>

      <el-button class="submit-button" @click="submitAnswer(answersRef)">提交</el-button>
    </el-col>
    <el-col :span="6" :xs="0"/>
  </el-row>

</template>
<script setup>
import {reactive, ref} from "vue";
import {get} from "../../api/survey.js";
import {list} from "../../api/question.js";
import {add} from "../../api/answer.js";
import {useRoute} from "vue-router";
import {ElMessage} from 'element-plus'

const route = useRoute()

const surveyId = Number(route.params.id)

const survey = ref({
  title: "",
  description: "",
  questions: [],
})
const answersRef = ref()
const rules = ({
  contact: [{required: true, message: '请输入联系方式', trigger: 'blur'}],

})

const form = reactive({
  contact: '',
  answers: {},
})

function initSurvey() {
  get(surveyId).then(res => {
    survey.value = res.data
  })
  list({pageNum: 1, pageSize: 999}, surveyId).then(res => {
    survey.value.questions = res.data.list
    survey.value.questions.forEach(q => {
      rules['answers.' + q.id] = [{required: true, message: '请填写', trigger: 'blur'}]
    })
  })

}

function submitAnswer(elForm) {
  elForm.validate((valid, fields) => {
    if (valid) {
      console.log('通过校验:')
      let answerResult = []

      for (const index in survey.value.questions) {
        const question = survey.value.questions[index]
        let options = question.options
        switch (question.type) {
          case 'text':
            answerResult.push(
                {
                  id: 0,
                  survey_id: surveyId,
                  question_id: question.id,
                  content: form.answers[question.id],
                  contact: form.contact,
                }
            );
            break;
          case 'radio':
            let extMsg = ''
            for (let o of options) {
              if (o.has_ext_msg === 'Y' && (o.extMsg === undefined || o.extMsg === '')) {
                ElMessage.error('请填写完整[第' + (Number(index) + 1) + '题]的补充信息')
                return
              }
              if (o.label === form.answers[question.id][0]) {
                extMsg = o.extMsg
              }
            }
            answerResult.push(
                {
                  id: 0,
                  survey_id: surveyId,
                  question_id: question.id,
                  label: form.answers[question.id].toString(),
                  ext_msg: extMsg,
                  contact: form.contact,
                }
            );
            break;
          case 'checkbox':
            const arr = form.answers[question.id]
            for (let str of arr) {
              let extMsg = ''
              for (let o of options) {
                if (o.has_ext_msg === 'Y' && o.label === str && (o.extMsg === undefined || o.extMsg === '')) {
                  ElMessage.error('请填写完整[第' + (Number(index) + 1) + '题,选项' + o.label + ']的补充信息')
                  return
                }
                if (o.label === str) {
                  extMsg = o.extMsg
                }
              }
              answerResult.push(
                  {
                    id: 0,
                    survey_id: surveyId,
                    question_id: question.id,
                    label: str,
                    ext_msg: extMsg,
                    contact: form.contact,
                  }
              );

            }
        }
      }
      add(answerResult).then(res => {
        if (res.success) {
          ElMessage.success(res.message)
          elForm.resetFields()
        } else {
          ElMessage.error(res.message)
        }
      })
    } else {
      ElMessage.error('请填写完整')
    }
  })
}


initSurvey()
</script>

<style scoped>

:deep(.el-checkbox) {
  height: auto;
  min-height: 40px;
  width: auto;
  padding: 5px;
}

:deep(.el-radio) {
  height: auto;
  min-height: 40px;
  width: auto;
  padding: 5px;
}

:deep(.el-radio__label) {
  white-space: pre-line;
}

:deep(.el-checkbox__label) {
  white-space: pre-line;
}

:deep(.el-radio-group) {
  display: grid;
}

:deep(.el-checkbox-group) {
  display: grid;
}

:deep(.el-form-item__content) {
  display: inline;
}


.survey-container {
  padding: 10px;
  background-color: #ffffff;
}

.survey-title {
  text-align: center;
  font-size: 24px;
  margin-bottom: 10px;
}

.survey-description {
  text-align: center;
  font-size: 16px;
  color: #666;
  margin-bottom: 20px;
}

.question-container {
  background-color: #fff;
  padding: 20px;
  margin-bottom: 1px;
  border: 1px solid #e0e0e0;
  border-radius: 20px;
}

.submit-button {
  display: block;
  margin: 0 auto;
  background-color: #007BFF;
  color: #fff;
  border: none;
  border-radius: 5px;
  padding: 10px 20px;
  cursor: pointer;
}

.submit-button:hover {
  background-color: #0056b3;
}


</style>
