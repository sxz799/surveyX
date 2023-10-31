<template>
  <h2 style="text-align: center">{{ survey.title }}</h2>
  <h4 style="text-align: center">{{ survey.description }}</h4>
  <div style="text-align: left" v-for="(question,index) in survey.questions">
    <el-row>
      <el-col :span="24">
        <el-text>{{ question.text }}</el-text>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-input type="textarea" v-model="answers[index]" v-if="question.type==='text'"
                  placeholder="请输入答案"/>

        <el-radio-group v-model="answers[index]" v-if="question.type==='radio'">
          <el-radio v-for="option in question.options" :label=option.label>
            {{ option.label }} ：{{ option.value }}
            <el-input v-model="option['extMsg']" v-if="option.has_ext_msg==='Y'"
                      placeholder="请输入答案"/>
          </el-radio>
        </el-radio-group>

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

  <el-button @click="submitAnswer">提交</el-button>
</template>
<script setup>
import {ref} from "vue";
import {get} from "../../api/survey.js";
import {list} from "../../api/question.js";
import {add} from "../../api/answer.js";


const props = defineProps({
  surveyId: Number,
})

const survey = ref({
  title: "",
  description: "",
  questions: [],
})

const answers = ref([])


function initSurvey() {
  get(props.surveyId).then(res => {
    survey.value = res.data
  })
  list({pageNum: 1, pageSize: 999}, props.surveyId).then(res => {
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
              survey_id: props.surveyId,
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
              survey_id: props.surveyId,
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
                survey_id: props.surveyId,
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
:deep(.el-radio) {
  display: block;

}

.el-checkbox-group .el-checkbox {
  display: block;
}
</style>
