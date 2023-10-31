<template>
  <el-row class="survey-container">
    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <h2 class="survey-title">{{ survey.title }}</h2>
      <h4 class="survey-description">{{ survey.description }}</h4>
      <div class="question-container" v-for="(question , index) in survey.questions">
        <el-row :gutter="10">
          <el-col :span="24">
              <el-tag  v-if="question.type === 'radio'">{{index+1}} [单选]</el-tag>
              <el-tag type="warning" v-if="question.type === 'checkbox'">{{index+1}} [多选]</el-tag>
              <el-tag type="success" v-if="question.type === 'text'">{{index+1}} [简答题]</el-tag>
            <el-text>{{ question.text }}</el-text>

          </el-col>
        </el-row>
        <br>
        <el-row :gutter="10">
          <el-col :span="24">
            <el-input type="textarea" autosize v-model="answers[question.id]" v-if="question.type==='text'"
                      placeholder=""/>

            <el-checkbox-group :max="1" v-model="answers[question.id]" v-if="question.type==='radio'">
              <el-checkbox class="wrap-checkbox" v-for="option in question.options" :label=option.label>
                {{ option.label }} ：{{ option.value }}
                <el-input v-model="option['extMsg']" size="small"
                          v-if="answers[question.id]&& answers[question.id][0] === option.label &&  option.has_ext_msg === 'Y'"
                          placeholder=""/>
              </el-checkbox>
            </el-checkbox-group>

            <el-checkbox-group v-model="answers[question.id]" v-if="question.type==='checkbox'">
              <el-checkbox class="wrap-checkbox" v-for="option in question.options" :label="option.label">
                {{ option.label }} ：{{ option.value }}
                <el-input v-model="option['extMsg']" size="small"
                          v-if="answers[question.id] && answers[question.id].toString().indexOf(option.label)>-1 && option.has_ext_msg==='Y'"
                          placeholder=""/>
              </el-checkbox>
            </el-checkbox-group>

          </el-col>
        </el-row>
      </div>

      <el-button class="submit-button" @click="submitAnswer">提交</el-button>
    </el-col>
    <el-col :span="6" :xs="0"/>
  </el-row>

</template>
<script setup>
import {ref} from "vue";
import {get} from "../../api/survey.js";
import {list} from "../../api/question.js";
import {add} from "../../api/answer.js";
import {useRoute} from "vue-router";
import { ElMessage } from 'element-plus'

const route = useRoute()

const surveyId = Number(route.params.id)

const survey = ref({
  title: "",
  description: "",
  questions: [],
})

const answers = ref({})


function initSurvey() {
  get(surveyId).then(res => {
    survey.value = res.data
  })
  list({pageNum: 1, pageSize: 999}, surveyId).then(res => {
    survey.value.questions = res.data.list
  })

}

function submitAnswer() {
  let answerResult = []
  for (const question of survey.value.questions) {
    let options = question.options
    switch (question.type) {
      case 'text':
        answerResult.push(
            {
              id: 0,
              survey_id: surveyId,
              question_id: question.id,
              content: answers.value[question.id],
            }
        );
        break;
      case 'radio':
        let extMsg = ''
        for (let o of options) {
          if (o.label === answers.value[question.id][0]) {
            extMsg = o.extMsg
          }
        }
        answerResult.push(
            {
              id: 0,
              survey_id: surveyId,
              question_id: question.id,
              label: answers.value[question.id].toString(),
              ext_msg: extMsg,
            }
        );
        break;
      case 'checkbox':
        const arr = answers.value[question.id]
        for (let str of arr) {
          let extMsg = ''
          for (let o of options) {
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
              }
          );

        }
    }
  }
  add(answerResult).then(res => {
    if(res.success){
      ElMessage.success(res.message)
    }else {
      ElMessage.error(res.message)
    }
  })
}

initSurvey()
</script>

<style scoped>
:deep(.el-checkbox){
  height:auto;
  padding: 5px;
}
:deep(.el-checkbox__label){
  white-space:pre-line;
}

.survey-container {
  padding: 20px;
  background-color: #f5f5f5;
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
  margin-bottom: 10px;
  border: 1px solid #e0e0e0;
  border-radius: 10px;
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
