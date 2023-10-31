<template>
  <el-row class="survey-container">
    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <h2 class="survey-title">{{ survey.title }}</h2>
      <h4 class="survey-description">{{ survey.description }}</h4>
      <div class="question-container" v-for="(question,index) in survey.questions">
        <el-row>
          <el-col :span="24">
            <el-text>{{ question.text }}</el-text>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="24">
            <el-input type="textarea" v-model="answers[index]" v-if="question.type==='text'"
                      placeholder="请输入答案"/>

            <el-checkbox-group :max="1" v-model="answers[index]" v-if="question.type==='radio'">
              <el-checkbox v-for="option in question.options"  :label=option.label @change="()=>{console.log(answers[index].toString())}">
                {{ option.label }} ：{{ option.value }}
                <el-input v-model="option['extMsg']"  v-if="option.has_ext_msg === 'Y'"
                          placeholder="请输入答案"/>
              </el-checkbox>
            </el-checkbox-group>

            <el-checkbox-group v-model="answers[index]" v-if="question.type==='checkbox'">
              <el-checkbox v-for="option in question.options" :label="option.label">
                {{ option.label }} ：{{ option.value }}
                <el-input v-model="option['extMsg']" v-if="option.has_ext_msg==='Y'"
                          placeholder="请输入答案"/>
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


const route = useRoute()

const surveyId = route.params.id

const survey = ref({
  title: "",
  description: "",
  questions: [],
})

const answers = ref([])


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
  for (let i = 0; i < survey.value.questions.length; i++) {
    let question = survey.value.questions[i]
    let options = question.options
    switch (survey.value.questions[i].type) {
      case 'text':
        answerResult.push(
            {
              id: 0,
              survey_id: surveyId,
              question_id: survey.value.questions[i].id,
              content: answers.value[i],
            }
        );
        break;
      case 'radio':
        let extMsg = ''
        for (let o of options) {
          if (o.label === answers.value[i]) {
            extMsg = o.extMsg
          }
        }
        answerResult.push(
            {
              id: 0,
              survey_id: surveyId,
              question_id: question.id,
              label: answers.value[i],
              ext_msg: extMsg,
            }
        );
        break;
      case 'checkbox':
        const arr = answers.value[i]
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
                question_id: survey.value.questions[i].id,
                label: str,
                ext_msg: extMsg,
              }
          );
        }
    }


  }


  add(answerResult).then(res => {
    console.log(res.message)
  })
}

initSurvey()
</script>

<style scoped>
.el-checkbox-group .el-checkbox {
  display: block;
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
  margin-bottom: 20px;
  border: 1px solid #e0e0e0;
  border-radius: 5px;
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
