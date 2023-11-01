<template>
  <el-row class="survey-container">
    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <h2 class="survey-title">{{ survey.title }}</h2>
      <h4 class="survey-description">{{ survey.description }}</h4>
      <div class="question-container" v-for="(question , index) in survey.questions">
        <el-row :gutter="10">
          <el-col :span="24">
            <el-tag v-if="question.type === 'radio'">{{ index + 1 }} [单选]</el-tag>
            <el-tag type="warning" v-if="question.type === 'checkbox'">{{ index + 1 }} [多选]</el-tag>
            <el-tag type="success" v-if="question.type === 'text'">{{ index + 1 }} [简答题]</el-tag>
            <el-text>{{ question.text }}</el-text>

          </el-col>
        </el-row>
        <br>
        <el-row :gutter="10">
          <el-col :span="24">
            <el-input type="textarea" autosize v-model="answers[question.id]" v-if="question.type==='text'"
                      placeholder=""/>

            <el-radio-group v-model="answers[question.id]" v-if="question.type==='radio'">
              <div v-for="option in question.options">
                <el-radio :label=option.label>
                  {{ option.label }} {{ option.value }}
                </el-radio>
                <div style="width: 65%;padding-left: 25px">
                  <el-input v-model="option['extMsg']" size="small" type="textarea" autosize
                            v-if="answers[question.id]&& answers[question.id][0] === option.label &&  option.has_ext_msg === 'Y'"
                            placeholder="请在此处填写补充信息！"/>
                </div>
              </div>
            </el-radio-group>

            <el-checkbox-group v-model="answers[question.id]" v-if="question.type==='checkbox'">
              <div v-for="option in question.options">

                <el-checkbox :label=option.label>
                  {{ option.label }} {{ option.value }}

                </el-checkbox>
                <div style="width: 94%;margin: auto">
                  <el-input v-model="option['extMsg']" size="small" type="textarea" autosize
                            v-if="answers[question.id] && answers[question.id].toString().indexOf(option.label)>-1 && option.has_ext_msg==='Y'"
                            placeholder="请在此处填写补充信息！"/>
                </div>
              </div>
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
import {ElMessage} from 'element-plus'

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
    if (res.success) {
      ElMessage.success(res.message)
    } else {
      ElMessage.error(res.message)
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
