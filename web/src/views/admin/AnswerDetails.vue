

<template>
  <el-table fit border size="small" :data="answerList">
    <el-table-column width="50" align="center" v-if="questionType!=='text'" label="选项" prop="label"/>
    <el-table-column v-if="questionType!=='text'" label="备注" prop="ext_msg"/>
    <el-table-column v-if="questionType==='text'" label="答案" prop="content"/>
    <el-table-column label="填写时间" prop="create_at">
      <template #default="scope">
        <span>{{ scope.row.create_at.slice(0,19).replace("T"," ") }}</span>
      </template>
    </el-table-column>
    <el-table-column label="联系方式" prop="contact"/>
    <el-table-column label="浏览器指纹" prop="finger"/>
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
</template>


<script setup>

import {list} from "@/api/admin/answer.js";
import {onMounted, reactive, ref, toRefs, watch} from "vue";
import {ElMessage} from "element-plus";

const props = defineProps({
  questionType: String,
  questionId: Number,
})

const data = reactive({
  queryParams: {
    pageNum: 1,
    pageSize: 10,
    question_id: props.questionId
  },
});

const answerList = ref([])
const total = ref(0)

const {queryParams} = toRefs(data);

watch(() => props.questionId, (newValue) => {
  console.log('questionId change', newValue)
  console.log(props.questionType)
  queryParams.value.question_id= newValue
  getList()
});

onMounted(() => {
  getList()
});

function getList() {
  list(queryParams.value).then(res => {
    if(res.success){
      ElMessage.success(res.message)
      answerList.value = res.data.list
      total.value = res.data.total
    }else{
      ElMessage.error(res.message)
    }

  })
}

function handleSizeChange(val) {
  queryParams.value.pageSize = val
  getList()
}

function handleCurrentChange(val) {
  queryParams.value.pageNum = val
  getList()
}

</script>

<style scoped>



</style>
