<template>


  <div class="analysis-container">
    <el-tag>题目数量:{{ analysisSurveyData.QuestionCount }}</el-tag>
    <el-tag>浏览器数量:{{ analysisSurveyData.FingerCount }}</el-tag>
    <el-tag>联系方式数量:{{ analysisSurveyData.ContactCount }}</el-tag>
    <el-tag>最早答题时间:{{ analysisSurveyData.MinCreateAt }}</el-tag>
    <el-tag>最晚答题时间:{{ analysisSurveyData.MaxCreateAt }}</el-tag>
  </div>


  <el-table border :data="questionList" @expand-change="ExpandChange">
    <el-table-column type="expand">
      <template #default="props">
        <div style="width: 80%;margin: auto">
          <el-table v-if="props.row.type!=='text'" border size="small" :data="props.row.options">
            <el-table-column label="选项" prop="label"/>
            <el-table-column label="内容" prop="value"/>
          </el-table>
          <el-table border size="small" :data="questionAnalysisResults[props.row.id]">
            <el-table-column v-if="props.row.type!=='text'" label="选项" prop="label"/>
            <el-table-column v-if="props.row.type!=='text'" label="备注" prop="ext_msg"/>
            <el-table-column v-if="props.row.type==='text'" label="简答题答案" prop="content"/>
            <el-table-column label="浏览器指纹" prop="finger"/>
            <el-table-column label="联系方式" prop="contact"/>
          </el-table>
        </div>
      </template>
    </el-table-column>

    <el-table-column label="题目" align="left" key="text" prop="text" :show-overflow-tooltip="true">
      <template #default="scope">
        <el-tag v-if="scope.row.type === 'radio'">单选</el-tag>
        <el-tag type="warning" v-if="scope.row.type === 'checkbox'">多选</el-tag>
        <el-tag type="success" v-if="scope.row.type === 'text'">简答</el-tag>
        <span>{{ scope.row.text }}</span>
      </template>
    </el-table-column>

  </el-table>
  <el-pagination
      style="padding-top: 20px"
      small
      :style="{'justify-content':'center'}"
      :background="true"
      :hide-on-single-page="false"
      :current-page="queryParams.pageNum"
      :page-size="queryParams.pageSize"
      :page-sizes="[2,5, 10, 30, 50]"
      :total="total"
      layout=" sizes, prev, pager, next"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
  />


</template>

<script setup>

import {ElMessage} from "element-plus";
import {analysis} from "@/api/admin/survey.js";
import {list as listQuestion} from "@/api/admin/question.js";
import {list as listAnswer} from "@/api/admin/answer.js";
import {onMounted, reactive, ref, watch} from "vue";


const props = defineProps({
  surveyId: String,
})

const total = ref(0)
const questionList = ref([])

const questionAnalysisResults = ref({})

const analysisSurveyData = ref({
  ContactCount: 0,
  FingerCount: 0,
  MaxCreateAt: '',
  MinCreateAt: '',
  QuestionCount: 0,
});

const queryParams = reactive({
  pageNum: 1,
  pageSize: 10,
  survey_id: props.surveyId
});


watch(() => props.surveyId, (newValue) => {
  queryParams.survey_id = newValue
  getSurveyAnalysis()
});

onMounted(() => {
  getSurveyAnalysis()
})


function ExpandChange(row, expandedRows) {
  getQuestionAnalysis(row.id)
}

function getSurveyAnalysis() {
  analysis(props.surveyId).then(res => {
    if (res.success) {
      analysisSurveyData.value = res.data
    } else {
      ElMessage.error(res.message)
    }
  })
  ListQuestion()
}

function ListQuestion() {
  listQuestion(queryParams).then(res => {
    if (res.success) {
      ElMessage.success(res.message)
      questionList.value = res.data.list
      total.value = res.data.total
    } else {
      ElMessage.error(res.message)
    }
  })
}

function getQuestionAnalysis(question_id) {
  listAnswer({question_id: question_id}).then(res => {
    if (res.success) {
      questionAnalysisResults.value[question_id] = res.data
    } else {
      ElMessage.error(res.message)
    }
  })
}

function handleSizeChange(val) {
  queryParams.pageSize = val
  ListQuestion()
}

function handleCurrentChange(val) {
  queryParams.pageNum = val
  ListQuestion()
}


</script>

<style scoped>
.analysis-container {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 20px;
}

.el-tag {
  margin-bottom: 8px;
  margin-right: 8px;
}

</style>
