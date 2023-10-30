<template>
  <h2 style="text-align: center">{{ survey.title }}</h2>
  <h4 style="text-align: center">{{ survey.description }}</h4>
  <div style="text-align: left" v-for="(question) in survey.questions">
    <el-row>
      <el-col :span="24">
        <el-text>{{ question.text }}</el-text>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-input type="textarea" v-model="answers[question.id]" v-if="question.type==='text'"
                  placeholder="请输入答案"/>
        <el-radio-group v-model="answers[question.id]" v-if="question.type==='radio'">
          <el-radio v-for="op in question.options" :label=op.value size="large">{{ op.value }}</el-radio>
        </el-radio-group>
        <el-checkbox-group v-model="answers[question.id]" v-if="question.type==='checkbox'">
          <el-checkbox v-for="(option) in question.options" :label="option.value">{{ option.value }}</el-checkbox>
        </el-checkbox-group>
      </el-col>
    </el-row>
  </div>

  <el-button @click="submitAnswer">提交</el-button>
</template>
<script setup>
import {ref} from "vue";
import {get} from "../api/survey.js";
import {list} from "../api/question.js";
import {add} from "../api/answer.js";


const props = defineProps({
  surveyId: Number,
})

const survey = ref({
  title: "",
  description: "",
  questions: [],
})

const answers = ref({})


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
  let entries = Object.entries(answers.value);
  entries.forEach(([key, value]) => {
    answerResult.push(
        {
          id: 0,
          survey_id: props.surveyId,
          question_id: Number(key),
          content: value.toString()
        }
    )
  });

  add(answerResult).then(res => {
    console.log(res.message)
  })
}

initSurvey()
</script>
