<template>
  <el-watermark :gap="[120,120]" :content="survey.water_mark">
    <el-row class="survey-container">

      <el-col :span="6" :xs="0"/>
      <el-col :span="12" :xs="24">
        <h2 class="survey-title">{{ survey.title }}</h2>
        <h4 class="survey-description">{{ survey.description }}</h4>
        <el-form :disabled="disabled" ref="answersRef" :model="form" :rules="rules" label-width="0">
          <div class="question-container" v-for="(question , index) in survey.questions">
            <el-form-item :prop="'answers.' + question.id ">
              <!--题目-->
              <el-tag v-if="question.type === 'radio'">{{ index + 1 }} [单选]</el-tag>
              <el-tag type="warning" v-if="question.type === 'checkbox'">{{ index + 1 }} [多选]</el-tag>
              <el-tag type="success" v-if="question.type === 'text'">{{ index + 1 }} [简答]</el-tag>
              <el-text>{{ question.text }}</el-text>
              <br>
              <!--简答-->
              <el-input type="textarea" :autosize="{ minRows: 2, maxRows: 5 }" v-model="form.answers[question.id]"
                        v-if="question.type==='text'"
                        placeholder=""/>
              <!--单选题-->
              <el-radio-group v-model="form.answers[question.id]" v-if="question.type==='radio'">
                <div v-for="option in question.options">
                  <el-row>
                    <el-col :span="16" :xs="24">
                      <el-radio :label=option.label>
                        {{ option.label }} {{ option.value }}
                      </el-radio>
                    </el-col>
                    <el-col :span="8" :xs="24">
                      <div style="margin: auto;padding-left: 5%">
                        <el-input v-model="option['extMsg']" size="small" type="textarea"
                                  :autosize="{ minRows: 1, maxRows: 3 }"
                                  v-if="form.answers[question.id] && form.answers[question.id][0] === option.label &&  option.has_ext_msg === 'yes'"
                                  placeholder="请在此处填写补充信息！"/>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-radio-group>
              <!--多选题-->
              <el-checkbox-group v-model="form.answers[question.id]" v-if="question.type==='checkbox'">
                <div v-for="option in question.options">
                  <el-row>
                    <el-col :span="16" :xs="24">
                      <el-checkbox :label=option.label>
                        {{ option.label }} {{ option.value }}
                      </el-checkbox>
                    </el-col>
                    <el-col :span="8" :xs="24">
                      <div style="margin: auto;padding-left: 5%">
                        <el-input v-model="option['extMsg']" size="small" type="textarea"
                                  :autosize="{ minRows: 1, maxRows: 3 }"
                                  v-if="form.answers[question.id] && form.answers[question.id].toString().indexOf(option.label)>-1 && option.has_ext_msg==='yes'"
                                  placeholder="请在此处填写补充信息！"/>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-checkbox-group>

            </el-form-item>
          </div>
          <el-form-item style="width: 80%;" v-if="survey.need_contact==='yes'" label="联系方式:" label-width="35%"
                        prop="contact">
            <el-input v-model="form.contact" placeholder="请填写联系方式"></el-input>
          </el-form-item>
        </el-form>
        <el-button class="submit-button" v-if="allowSubmit" @click="submitAnswer(answersRef)">提交</el-button>
        <div style="position: relative; bottom: 0; left: 0; right: 0; text-align: center; padding: 20px 0;">
          <el-link type="info" href="https://github.com/sxz799/surveyX">SurveyX 提供技术支持</el-link>
        </div>
      </el-col>
      <el-col :span="6" :xs="0"/>

    </el-row>
  </el-watermark>
</template>
<script setup>

import {onMounted, reactive, ref} from "vue";
import {get} from "@/api/user/survey.js";
import {list} from "@/api/user/question.js";
import {add} from "@/api/user/answer.js";
import {useRoute} from "vue-router";
import {ElNotification} from 'element-plus'
import Fingerprint2 from 'fingerprintjs2';

const disabled = ref(false)
const allowSubmit = ref(true)
const route = useRoute()
const surveyId = route.params.id
const finger = ref('')
const survey = reactive({
  title: '',
  description: '',
  need_contact: '',
  water_mark: [],
  questions: [],
})
const answersRef = ref()
const rules = ({
  contact: [{required: true, message: '请填写联系方式', trigger: 'blur'}],
})

const form = reactive({
  contact: undefined,
  answers: {},
})

onMounted(() => {
  initSurvey();
  getFinger();
});

async function initSurvey() {
  const surveyData = await get(surveyId);
  const datetime = new Date();
  if (datetime < new Date(surveyData.data.start_time) || datetime > new Date(surveyData.data.end_time)) {
    ElNotification({
      title: '抱歉',
      message: '当前问卷不在答题时间内!',
      type: 'warning',
    })
    allowSubmit.value = false
    return
  }
  survey.title = surveyData.data.title;
  survey.description = surveyData.data.description;
  survey.need_contact = surveyData.data.need_contact;
  survey.water_mark = surveyData.data.water_mark.split('\n');
  survey.questions = (await list({pageNum: 1, pageSize: 99999, survey_id: surveyId})).data.list;
  survey.questions.forEach((q) => {
    rules[`answers.${q.id}`] = [{required: true, message: "请填写", trigger: q.type === 'text' ? "blur" : "change"}];
  });
}

function submitAnswer(elForm) {
  elForm.validate((valid) => {
    if (!valid) {
      ElNotification({
        title: '问卷信息不完整！',
        type: 'info',
      })
      return
    }
    let answerResult = []
    // 遍历问题
    for (const index in survey.questions) {
      const question = survey.questions[index] // 问题
      let options = question.options// 选项
      const answer = form.answers[question.id].toString()// 答案
      switch (question.type) {// 答案类型
        case 'text':
          answerResult.push({question_id: question.id, content: answer});// 答案
          break;
        case 'radio':
          let extMsg = ''
          for (let option of options) {// 遍历选项
            // 如果有补充信息，但是没有填写
            if (option.has_ext_msg === 'yes' && option.label === answer && (option.extMsg === undefined || option.extMsg === '')) {
              ElNotification({
                title: '信息不全',
                message: '请填写[第' + (Number(index) + 1) + '题]的补充信息',
                type: 'warning',
              })
              return
            }
            // 如果选项等于答案，就把补充信息赋值给extMsg
            if (option.label === answer) {
              extMsg = option.extMsg
            }
          }
          // 把答案添加到答案数组中
          answerResult.push({question_id: question.id, label: answer, ext_msg: extMsg,});
          break;
        case 'checkbox':
          // 选项是一个数组，遍历选项
          const labels = answer.split(',')
          for (let label of labels) {
            let extMsg = ''
            // 如果有补充信息，但是没有填写
            for (let option of options) {
              if (option.has_ext_msg === 'yes' && option.label === label && (option.extMsg === undefined || option.extMsg === '')) {
                ElNotification({
                  title: '信息不全',
                  message: '请填写[第' + (Number(index) + 1) + '题,选项' + option.label + ']的补充信息',
                  type: 'warning',
                })
                return
              }
              // 如果选项等于答案，就把补充信息赋值给extMsg
              if (option.label === label) {
                extMsg = option.extMsg
              }
            }
            // 把答案添加到答案数组中
            answerResult.push({question_id: question.id, label: label, ext_msg: extMsg,});
          }
      }
    }
    // 遍历答案数组，添加指纹、问卷id、联系方式
    answerResult.forEach(a => {
      a.finger = finger.value
      a.survey_id = surveyId
      a.contact = form.contact
    })
    // 提交答案
    add(answerResult).then(res => {
      if (res.success) {
        ElNotification({
          title: '提交成功',
          message: res.message,
          type: 'success',
        })
        allowSubmit.value = false
        disabled.value = true
      } else {
        ElNotification({
          title: '提交失败',
          message: res.message,
          type: 'error',
        })
      }
    })
  })
}

// 获取浏览器指纹
function getFinger() {
  Fingerprint2.get((components) => {
    // components数组包含了浏览器指纹的各个组成部分
    const values = components.map((component) => component.value);
    finger.value = Fingerprint2.x64hash128(values.join(''), 31);
  });
}


</script>

<style scoped>

:deep(.el-checkbox) {
  height: auto;
  min-height: 40px;
  width: auto;
  padding: 3px;
}

:deep(.el-radio) {
  height: auto;
  min-height: 40px;
  width: auto;
  padding: 3px;
}

:deep(.el-radio__label) {
  white-space: pre-line;
  line-height: normal;
}

:deep(.el-checkbox__label) {
  white-space: pre-line;
  line-height: normal;
}

:deep(.el-radio-group) {
  display: grid;
}

:deep(.el-checkbox-group) {
  display: grid;
}

:deep(.el-form-item__content) {
  display: unset;
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
  border: 1px solid #e0e0e0;
  border-radius: 20px;
  margin-top: 20px;
  margin-bottom: 20px;
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
