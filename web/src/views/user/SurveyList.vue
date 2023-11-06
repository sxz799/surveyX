<template>
  <el-row>

    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <div class="app-container">
        <el-table :data="surveyList">
<!--          <el-table-column type="selection" width="50" align="center"/>-->
          <el-table-column label="序号" width="70" align="center" type="index"/>
          <el-table-column label="标题" align="center" key="title" prop="title" :show-overflow-tooltip="true"/>
          <el-table-column label="描述" align="center" key="description" prop="description"
                           :show-overflow-tooltip="true"/>
          <el-table-column label="操作" align="center" width="150">
            <template #default="scope">
              <el-button type="text" :icon="Edit" @click="router.push('/survey/'+scope.row.id)">答题</el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-divider/>
        <el-pagination
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
      </div>
    </el-col>
    <el-col :span="6" :xs="0"/>
  </el-row>

</template>

<script setup>
import {onMounted, reactive, ref} from 'vue'
import {list} from "@/api/user/survey.js";
import {useRouter} from "vue-router";
import {Edit} from "@element-plus/icons";

const router = useRouter()
const surveyList = ref([])
const total = ref(0)
const queryParams = reactive({
  pageNum: 1,
  pageSize: 10,
  status:'yes',
  start_time: new Date(),
  end_time: new Date(),
})

onMounted(() => {
  getList()
})


function getList() {
  list(queryParams).then(res => {
    surveyList.value = res.data.list
    total.value = res.data.total
  })
}

function handleSizeChange(val) {
  getList()
}

function handleCurrentChange(val) {
  getList()
}




</script>

<style scoped>

.app-container {
  padding: 20px;
}

</style>
