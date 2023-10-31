<template>
  <el-row>

    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <div class="app-container">
        <el-table :data="surveyList">
          <el-table-column type="selection" width="50" align="center"/>
          <el-table-column label="ID" align="center" key="id" prop="id"/>
          <el-table-column label="标题" align="center" key="title" prop="title" :show-overflow-tooltip="true"/>
          <el-table-column label="描述" align="center" key="description" prop="description" :show-overflow-tooltip="true"/>
          <el-table-column label="操作" align="center" width="150">
            <template #default="scope">
              <router-link :to="`/survey/${scope.row.id}`">答题</router-link>
            </template>
          </el-table-column>
        </el-table>
        <el-pagination
            v-show="total > 0"
            :total="total"
            v-model:page="queryParams.pageNum"
            v-model:limit="queryParams.pageSize"
            @pagination="getList"
        />
      </div>
    </el-col>
    <el-col :span="6" :xs="0"/>
  </el-row>

</template>

<script setup>
import {reactive, ref, toRefs} from 'vue'
import {list} from "../../api/survey.js";

const surveyList = ref([])
const total = ref(0)
const data = reactive({
  queryParams: {
    pageNum: 1,
    pageSize: 10,
  },
});
const {queryParams} = toRefs(data);

function getList() {
  list(queryParams.value).then(res => {
    surveyList.value = res.data.list
    total.value = res.data.total
  })
}




getList()

</script>

<style scoped>

</style>
