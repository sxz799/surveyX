<template>
  <div class="app-container">
    <el-row>
      <el-col :span="6" :xs="0"/>
      <el-col :span="12" :xs="24">

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
                type="primary"
                plain
                :icon="Plus"
                @click="handleAdd"
            >新增
            </el-button>
          </el-col>
        </el-row>
        <el-table border :data="surveyList">
          <el-table-column type="selection" width="50" align="center"/>
          <el-table-column label="ID" width="50" align="center" key="id" prop="id"/>
          <el-table-column label="标题" align="center" key="title" prop="title" :show-overflow-tooltip="true">
            <template #default="scope">
              <el-link :icon="Edit" type="warning" @click="handleQuestion(scope.row.id)">{{ scope.row.title }}
              </el-link>
            </template>
          </el-table-column>
          <el-table-column label="描述" align="center" key="description" prop="description"
                           :show-overflow-tooltip="true"/>
          <el-table-column width="100" label="状态" align="center" key="status" prop="status"
                           :show-overflow-tooltip="true">
            <template #default="scope">
              <span v-if="scope.row.status === 'Y'">启用</span>
              <span v-if="scope.row.status === 'N'">禁用</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="150">
            <template #default="scope">
              <el-button link type="primary" @click="handleEdit(scope.row)" :icon="Edit">修改</el-button>
              <el-button link type="success" @click="handleQuestion(scope.row.id)" :icon="Tools">配置题目</el-button>
              <el-button link type="danger" @click="handleDelete(scope.row)" :icon="Delete">删除</el-button>
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


        <el-dialog :title="title" v-model="open" :width="dialogWidth" append-to-body>
          <el-form ref="surveyRef" :model="form" label-width="100px">
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
        <el-dialog title="配置题目" v-model="openDetails" :width="dialogWidth" append-to-body>
          <template #footer>
            <QuestionList :surveyId="surveyId"/>
            <div class="dialog-footer">
              <el-button @click="openDetails = false">完成</el-button>
            </div>
          </template>
        </el-dialog>

      </el-col>
      <el-col :span="6" :xs="0"/>
    </el-row>
  </div>
</template>

<script setup>
import {computed, reactive, ref, toRefs} from 'vue'
import {list, add, del, update, get} from "../../api/survey.js";
import QuestionList from "./QuestionList.vue";
import {Delete, Edit, Plus, Tools} from "@element-plus/icons";


const surveyId = ref(0)
const open = ref(false)
const openDetails = ref(false)

const title = ref('')
const surveyList = ref([])
const total = ref(0)
const dialogWidth = computed(() => {
  return window.innerWidth > 768 ? '60%' : '90%'
})


const data = reactive({
  form: {},
  queryParams: {
    pageNum: 1,
    pageSize: 10,
  },
});

const {queryParams, form} = toRefs(data);


function getList() {
  list(queryParams.value).then(res => {
    surveyList.value = res.data.list
    total.value = res.data.total
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

function handleAdd() {
  reset();
  open.value = true
  title.value = '新增'

}

function reset() {
  form.value = {
    id: undefined,
    title: "",
    description: "",
    status: "Y",
  };

}

function handleEdit(row) {
  reset();
  get(row.id).then(res => {
    form.value = res.data
    open.value = true
    title.value = '修改'
  })
}

function handleQuestion(id) {
  surveyId.value = id
  openDetails.value = true
}

function handleDelete(row) {
  del(row.id).then(res => {
    getList()
  })
}

function submitForm() {
  if (form.value.id) {
    update(form.value).then(res => {
      open.value = false
      getList()
    })
  } else {
    add(form.value).then(res => {
      open.value = false
      getList()
    })
  }
}


getList()

</script>

<style scoped>

.app-container {
  padding: 20px;
}

.mb8 {
  margin-bottom: 8px;
}

</style>
