

<template>


  <div class="analysis-container">
    <el-tag>题目数量:{{analysisData.QuestionCount}}</el-tag>
    <el-tag>浏览器数量:{{analysisData.FingerCount}}</el-tag>
    <el-tag>联系方式数量:{{analysisData.ContactCount}}</el-tag>
    <el-tag>最早答题时间:{{analysisData.MinCreateAt}}</el-tag>
    <el-tag>最晚答题时间:{{analysisData.MaxCreateAt}}</el-tag>
  </div>



</template>

<script setup>

import {ElMessage} from "element-plus";
import {analysis} from "@/api/admin/survey.js";
import {onMounted, reactive, ref, watch} from "vue";
const props = defineProps({
  surveyId: String,
})

const analysisData = ref({
  ContactCount: 0,
  FingerCount: 0,
  MaxCreateAt: '',
  MinCreateAt: '',
  QuestionCount: 0,
});

watch(() => props.surveyId, (newValue) => {
  getAnalysis()
});

onMounted(() => {
  getAnalysis()
})

function getAnalysis() {

  analysis(props.surveyId).then(res => {
    if (res.success) {
      ElMessage.success(res.message)
      analysisData.value= res.data
    } else {
      ElMessage.error(res.message)
    }
  })

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
