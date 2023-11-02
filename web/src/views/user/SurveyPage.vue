<template>
  <el-row class="survey-container">
    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <h2 class="survey-title">{{ survey.title }}</h2>
      <h4 class="survey-description">{{ survey.description }}</h4>
      <el-form ref="answersRef" :model="form" :rules="rules" label-width="0">
        <el-form-item v-for="(question , index) in survey.questions" :prop="'answers.' + question.id ">
          <div class="question-container">
            <!--题目-->
            <el-tag v-if="question.type === 'radio'">{{ index + 1 }} [单选]</el-tag>
            <el-tag type="warning" v-if="question.type === 'checkbox'">{{ index + 1 }} [多选]</el-tag>
            <el-tag type="success" v-if="question.type === 'text'">{{ index + 1 }} [简答题]</el-tag>
            <el-text>{{ question.text }}[{{ question.id }}]</el-text>
            <br>
            <!--简答题-->
            <el-input type="textarea" :autosize="{ minRows: 2, maxRows: 5 }" v-model="form.answers[question.id]"
                      v-if="question.type==='text'"
                      placeholder=""/>
            <!--单选题-->
            <el-radio-group v-model="form.answers[question.id]" v-if="question.type==='radio'">
              <div style="display: flex;align-items: center" v-for="option in question.options">
                <div style="max-width: 100%;min-width: 65%;">
                  <el-radio :label=option.label>
                    {{ option.label }} {{ option.value }}
                  </el-radio>
                </div>
                <div style="width: 35%;display: contents">
                  <el-input v-model="option['extMsg']" size="small" type="textarea"
                            :autosize="{ minRows: 1, maxRows: 3 }"
                            v-if="form.answers[question.id] && form.answers[question.id][0] === option.label &&  option.has_ext_msg === 'yes'"
                            placeholder="请在此处填写补充信息！"/>
                </div>
              </div>
            </el-radio-group>
            <!--多选题-->
            <el-checkbox-group v-model="form.answers[question.id]" v-if="question.type==='checkbox'">
              <div style="display: flex;align-items: center" v-for="option in question.options">
                <div style="max-width: 100%;min-width: 65%;">
                  <el-checkbox :label=option.label>
                    {{ option.label }} {{ option.value }}
                  </el-checkbox>
                </div>
                <div style="width: 35%;display: contents">
                  <el-input v-model="option['extMsg']" size="small" type="textarea"
                            :autosize="{ minRows: 1, maxRows: 3 }"
                            v-if="form.answers[question.id] && form.answers[question.id].toString().indexOf(option.label)>-1 && option.has_ext_msg==='yes'"
                            placeholder="请在此处填写补充信息！"/>
                </div>
              </div>
            </el-checkbox-group>
          </div>
        </el-form-item>
        <el-form-item style="width: 80%;" v-if="survey.need_contact==='yes'" label="联系方式:" label-width="30%"
                      prop="contact">
          <el-input v-model="form.contact" placeholder="请填写联系方式"></el-input>
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
import Fingerprint2 from 'fingerprintjs2';

const route = useRoute()
const surveyId = Number(route.params.id)
const finger = ref('')
const survey = ref({
  title: '',
  description: '',
  questions: [],
  need_contact: '',
  repeat: '',
  repeat_check: '',
})
const answersRef = ref()
const rules = ({
  contact: [{required: true, message: '请输入联系方式', trigger: 'blur'}],
})

const form = reactive({
  contact: undefined,
  answers: {},
})

function initSurvey() {
  get(surveyId).then(res => {
    survey.value = res.data
  })
  list({pageNum: 1, pageSize: 99999}, surveyId).then(res => {
    survey.value.questions = res.data.list
    survey.value.questions.forEach(q => {
      rules['answers.' + q.id] = [{required: true, message: '请填写', trigger: 'blur'}]
    })
  })

}

function submitAnswer(elForm) {
  elForm.validate((valid) => {
    if (!valid) {
      ElMessage.error('请填写完整')
      return
    }
    let answerResult = []
    // 遍历问题
    for (const index in survey.value.questions) {
      const question = survey.value.questions[index] // 问题
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
              ElMessage.error('请填写完整[第' + (Number(index) + 1) + '题]的补充信息')
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
                ElMessage.error('请填写完整[第' + (Number(index) + 1) + '题,选项' + option.label + ']的补充信息')
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
        ElMessage.success(res.message)
        if (survey.value.repeat === 'yes') {
          elForm.resetFields()
        }
      } else {
        ElMessage.error(res.message)
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

initSurvey()
getFinger()
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
