<template>


  <div class="analysis-container">
    <el-tag>题目数量 : {{ analysisSurveyData.QuestionCount }}</el-tag>
    <el-tag>浏览器数量 : {{ analysisSurveyData.FingerCount }}</el-tag>
    <el-tag>联系方式数量 : {{ analysisSurveyData.ContactCount }}</el-tag>
    <el-tag>第一次答题时间 : {{ analysisSurveyData.MinCreateAt }}</el-tag>
    <el-tag>最后一次答题时间 : {{ analysisSurveyData.MaxCreateAt }}</el-tag>
  </div>


  <el-table border :data="questionList" @expand-change="ExpandChange">
    <el-table-column type="expand">
      <template #default="props">
        <div style="padding-left: 10px">
          <div v-if="props.row.type!=='text'">
            <el-tag v-for="op in props.row.options">{{ op.label }} : {{ op.value }}</el-tag>
          </div>
          <div v-if="questionAnalysisResults[props.row.id]"
               v-for="result in questionAnalysisResults[props.row.id].split('###')">
            <el-tag type="success">{{ result }}</el-tag>
          </div>

          <el-button size="small" type="info" plain
                     @click="handleAnswerDetails(props.row.type,questionAnalysisDetails[props.row.id])">查看答案详情
          </el-button>

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
      :page-sizes="[5, 10, 30, 50]"
      :total="total"
      layout=" sizes, prev, pager, next"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
  />

  <el-dialog title="答案详情" v-model="showDetails" width="50%" append-to-body>
    <AnswerDetails :question-type="currType" :answer-details="currDetails"></AnswerDetails>
  </el-dialog>


</template>

<script setup>

import {ElMessage} from "element-plus";
import {analysis} from "@/api/admin/survey.js";
import {list as listQuestion} from "@/api/admin/question.js";
import {list as listAnswer} from "@/api/admin/answer.js";
import {onMounted, reactive, ref, watch} from "vue";
import AnswerDetails from "@/views/admin/AnswerDetails.vue";

const props = defineProps({
  surveyId: String,
})

const showDetails = ref(false)
const currType = ref('')
const currDetails = ref([])


const total = ref(0)
const questionList = ref([])

const questionAnalysisDetails = ref({})
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
  //expandedRows是一个数组，里面包含了所有展开的行的数据,要过滤expandedRows，只保留当前行的数据
  expandedRows = expandedRows.filter(item => item.id === row.id)
  if (expandedRows.length > 0) {
    getQuestionAnalysis(row.id)
  }

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
      const labelCount = {};
      const uniqueFinger = new Set();
      const uniqueContact = new Set();
      let data = res.data
      questionAnalysisDetails.value[question_id] = data
      data.forEach(function (item) {
        uniqueFinger.add(item.finger)
        uniqueContact.add(item.contact)
        const label = item.label;
        labelCount[label] = (labelCount[label] || 0) + 1;
      });
      const totalCount = data.length;
      const labelPercentage = {};
      for (const label in labelCount) {
        const count = labelCount[label];
        const percentage = (count / totalCount) * 100;
        labelPercentage[label] = {
          count: count,
          percentage: percentage.toFixed(2) + "%",
        };
      }
      const result = "共有 " + uniqueFinger.size + " 个浏览器参与了本题调查,共留下 " + uniqueContact.size + " 个联系方式! ###";
      let result2 = "";
      for (const item in labelPercentage) {
        if (item === "") break
        result2 += "  选项 " + item + " 被选择了 " + labelPercentage[item].count + " 次,占比为: " + labelPercentage[item].percentage + " ###";
      }
      questionAnalysisResults.value[question_id] = result + result2 + "点击按钮查看答案详情"
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

function handleAnswerDetails(type, details) {
  console.log(type, details)
  currType.value = type
  currDetails.value = details
  showDetails.value = true
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
