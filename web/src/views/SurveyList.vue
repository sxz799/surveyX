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
      <el-table-column label="标题" align="center" key="title" prop="title" :show-overflow-tooltip="true" >
        <template #default="scope">
          <el-button @click="()=>{
            surveyId=scope.row.id
            openDetails=true
          }">{{scope.row.title}}</el-button>
        </template>
      </el-table-column>
      <el-table-column label="描述" align="center" key="description" prop="description" :show-overflow-tooltip="true" />
      <el-table-column label="状态" align="center" key="status" prop="status" :show-overflow-tooltip="true" />
      <el-table-column label="操作" align="center" width="150">
        <template #default="scope">
          <el-tooltip content="修改" placement="top">
            <el-button link type="primary" @click="handleUpdate(scope.row)" icon="Edit">Edit</el-button>
          </el-tooltip>
          <el-tooltip content="删除" placement="top" >
            <el-button link type="primary" @click="handleDelete(scope.row)" icon="Delete">Delete</el-button>
          </el-tooltip>
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

    <el-dialog :title="title" v-model="open" width="50%" append-to-body>
      <el-form ref="surveyRef" :model="form"  label-width="100px">
        <el-row>
          <el-col :span="24">
            <el-form-item label="标题" prop="title">
              <el-input v-model="form.title" placeholder="请输入标题"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="描述" prop="description">
              <el-input v-model="form.description" placeholder="请输入描述"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="状态" prop="status">
              <el-select v-model="form.status" placeholder="请选择状态">
                <el-option label="启用" value="Y"></el-option>
                <el-option label="禁用" value="N"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>


      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitForm">确 定</el-button>
          <el-button @click="open = false">取 消</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog title="配置题目" v-model="openDetails" width="50%" append-to-body>
      <template #footer>
        <QuestionList :surveyId="surveyId"/>
        <div class="dialog-footer">
          <el-button @click="open = false">完成</el-button>
        </div>
      </template>
    </el-dialog>





  </div>

</template>

<script setup>
import {reactive, ref, toRefs, watch} from 'vue'
import {list,add,del,update,get} from "../api/survey.js";
import QuestionList from "./QuestionList.vue";

const surveyId = ref(0)
const open = ref(false)
const openDetails = ref(false)
const title = ref('')
const surveyList = ref([])
const total = ref(0)



const data = reactive({
  form: {},
  queryParams: {
    pageNum: 1,
    pageSize: 10,
  },
});

const { queryParams, form } = toRefs(data);


function getList() {
  list(queryParams.value).then(res => {
    surveyList.value = res.data.list
    total.value = res.data.total
  })
}

function handleAdd() {
  reset();
  open.value = true
  title.value = '新增'

}

function reset() {
  form.value = {
    id:undefined,
    title: "",
    description: "",
    status: "Y",
  };

}

function handleUpdate(row) {
  reset();
  get(row.id).then(res => {
    form.value = res.data
    open.value = true
    title.value = '修改'
  })
}
function handleDelete(row) {
  del(row.id).then(res => {
    getList()
  })
}
function submitForm() {
  if(form.value.id){
    update(form.value).then(res => {
      open.value = false
      getList()
    })
  }else{
    add(form.value).then(res => {
      open.value = false
      getList()
    })
  }
}


getList()

</script>

<style scoped>

</style>
