<template>
  <div class="app-container">
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
            type="primary"
            plain
            icon="Plus"
            @click="handleAdd"
        >新增
        </el-button>
      </el-col>
    </el-row>

    <el-table :data="surveyList">
      <el-table-column type="selection" width="50" align="center" />
      <el-table-column label="ID" align="center" key="id" prop="id" />
      <el-table-column label="标题" align="center" key="title" prop="title" :show-overflow-tooltip="true" />
      <el-table-column label="描述" align="center" key="description" prop="description" :show-overflow-tooltip="true" />
      <el-table-column label="状态" align="center" key="status" prop="status" :show-overflow-tooltip="true" />
      <el-table-column label="操作" align="center" width="150">
        <template #default="scope">
          <el-tooltip content="修改" placement="top">
            <el-button link type="primary" icon="Edit">Edit</el-button>
          </el-tooltip>
          <el-tooltip content="删除" placement="top" >
            <el-button link type="primary" icon="Delete">Delete</el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>
    <pagination
        v-show="total > 0"
        :total="total"
        v-model:page="queryParams.pageNum"
        v-model:limit="queryParams.pageSize"
        @pagination="getList"
    />
  </div>

</template>

<script setup>
import {ref} from 'vue'
import {list,add,del,update,get} from "../../api/survey.js";

defineProps({
  msg: String,
})

const surveyList = ref([])
const total = ref(0)
const queryParams = ref({
  pageNum: 1,
  pageSize: 10,
})


function getList() {
  list(queryParams.value).then(res => {
    surveyList.value = res.data.list
    total.value = res.data.total
  })
}

function handleAdd() {

}


getList()

</script>

<style scoped>

</style>
