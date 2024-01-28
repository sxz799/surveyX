<template>

  <el-row class="mb8">
    <el-col :span="6">
      <el-statistic title="题目数量" group-separator=","  :value="analysisSurveyData.QuestionCount" />
    </el-col>
    <el-col :span="6">
      <el-statistic title="填写次数" group-separator=","   :value="analysisSurveyData.KeyCount" />
    </el-col>
    <el-col :span="6">
      <el-statistic title="浏览器数量" group-separator=","   :value="analysisSurveyData.FingerCount" />
    </el-col>
    <el-col :span="6">
      <el-statistic title="联系方式数量" group-separator=","   :value="analysisSurveyData.ContactCount" />
    </el-col>
    <el-col :span="12">
      <el-statistic title="第一次填写时间" group-separator=","   :value="analysisSurveyData.FirstCreateAt" />
    </el-col>
    <el-col :span="12">
      <el-statistic title="最后一次填写时间" group-separator=","   :value="analysisSurveyData.LastCreateAt" />
    </el-col>
  </el-row>


  <el-table border :data="questionList" @expand-change="ExpandChange">
    <el-table-column type="expand">
      <template #default="props">
        <div v-if="props.row.type!=='text'" style="padding-left: 50px">
          <Echarts :answerId=props.row.id :answerOptions=props.row.options></Echarts>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="题目" align="left" key="text" prop="text" :show-overflow-tooltip="true">
      <template #default="scope">
        <el-tag v-if="scope.row.type === 'radio'">单选</el-tag>
        <el-tag type="warning" v-if="scope.row.type === 'checkbox'">多选</el-tag>
        <el-tag type="success" v-if="scope.row.type === 'text'">简答</el-tag>
        <span>{{ scope.row.text }}</span>
        <el-button size="small" type="success" plain
                   @click="handleAnswerDetails(scope.row.type,scope.row.id)">查看答案详情
        </el-button>
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
      :page-sizes="[5, 10, 30, 50]"
      :total="total"
      layout=" sizes, prev, pager, next"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
  />

  <el-dialog title="答案详情" v-model="showDetails" width="70%" append-to-body>
    <AnswerDetails :question-type="currType" :question-id="currQuestionId" ></AnswerDetails>
  </el-dialog>


</template>

<script setup>


import {analysis as SurveyAnalysis} from "@/api/admin/survey.js";
import {list as listQuestion,analysis as QuestionAnalysis} from "@/api/admin/question.js";

import {onMounted, reactive, ref, watch} from "vue";
import AnswerDetails from "@/views/admin/AnswerDetails.vue";
import Echarts from "@/views/common/Echarts.vue";
import {ElMessage} from "element-plus";

const props = defineProps({
  surveyId: String,
})

const showDetails = ref(false)
const currType = ref('')
const currQuestionId = ref(0)


const total = ref(0)
const questionList = ref([])


const analysisSurveyData = ref({
  KeyCount: 0,
  ContactCount: 0,
  FingerCount: 0,
  LastCreateAt: '',
  FirstCreateAt: '',
  QuestionCount: 0,
});

const queryParams = reactive({
  pageNum: 1,
  pageSize: 10,
  survey_id: props.surveyId
});


watch(() => props.surveyId, (newValue) => {
  if (newValue==='') {
    questionList.value = []
    analysisSurveyData.value={
      KeyCount: 0,
      ContactCount: 0,
      FingerCount: 0,
      LastCreateAt: '',
      FirstCreateAt: '',
      QuestionCount: 0,
    }
    return
  }
  queryParams.survey_id = newValue
  getSurveyAnalysis()
});

onMounted(() => {
  getSurveyAnalysis()
})


function ExpandChange(row, expandedRows) {
  //expandedRows是一个数组，里面包含了所有展开的行的数据,要过滤expandedRows，只保留当前行的数据
  expandedRows = expandedRows.filter(item => item.id === row.id)
  if (expandedRows.length > 0) {
    currQuestionId.value=row.id
  }

}

function getSurveyAnalysis() {
  SurveyAnalysis(props.surveyId).then(res => {
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


function handleSizeChange(val) {
  queryParams.pageSize = val
  ListQuestion()
}

function handleCurrentChange(val) {
  queryParams.pageNum = val
  ListQuestion()
}

function handleAnswerDetails(type,qid) {
  currType.value = type
  currQuestionId.value=0
  currQuestionId.value=qid
  showDetails.value = true
}


</script>

<style scoped>
.mb8 {
  margin-bottom: 8px;
}
.el-tag {
  margin-bottom: 8px;
  margin-right: 8px;
}
.el-col {
  text-align: center;
}

</style>
