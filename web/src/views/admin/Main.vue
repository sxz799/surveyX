<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <Header></Header>
      </el-header>
      <el-main>
        <el-row :gutter="4">
          <el-col :span="8" :xs="24">
              <SurveyList @updateSurveyId="updateSurveyId"></SurveyList>
          </el-col>
          <el-col :span="16" :xs="24">
            <el-card>
              <el-tabs>
                <el-tab-pane label="题目管理">
                  <QuestionList v-if="openDetails" :surveyId="surveyId" />
                </el-tab-pane>
                <el-tab-pane label="问卷分析">
                  <Analysis v-if="openDetails" :surveyId="surveyId" />
                </el-tab-pane>
              </el-tabs>
            </el-card>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
    <div style="position: static; bottom: 0; left: 0; right: 0; text-align: center; padding: 20px 0;">
      <el-link type="info" href="https://github.com/sxz799/surveyX">SurveyX 提供技术支持</el-link>
    </div>
  </div>
</template>

<script setup>

import QuestionList from "./QuestionList.vue";
import Analysis from "@/views/admin/Analysis.vue";
import Header from "@/views/admin/Header.vue";
import SurveyList from "@/views/admin/SurveyList.vue";
import {ref} from "vue";

const openDetails = ref(false)
const surveyId = ref('')

function updateSurveyId(newValue){
  surveyId.value = newValue
  openDetails.value = true
}

</script>

<style scoped>


.app-container {
  padding: 20px;
}





</style>
