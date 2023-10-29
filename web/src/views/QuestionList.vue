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

    <el-table :data="questionList">
      <el-table-column type="selection" width="50" align="center"/>
      <el-table-column label="ID" align="center" key="id" prop="id"/>
      <el-table-column label="问卷ID" align="center" key="survey_id" prop="survey_id" :show-overflow-tooltip="true"/>
      <el-table-column label="题目" align="center" key="text" prop="text" :show-overflow-tooltip="true"/>
      <el-table-column label="类型" align="center" key="type" prop="type" :show-overflow-tooltip="true"/>
      <el-table-column label="选项" align="center" key="options" prop="options" :show-overflow-tooltip="true"/>
      <el-table-column label="排序" align="center" key="order" prop="order" :show-overflow-tooltip="true"/>
      <el-table-column label="操作" align="center" width="150">
        <template #default="scope">
          <el-tooltip content="修改" placement="top">
            <el-button link type="primary" @click="handleUpdate(scope.row)" icon="Edit">Edit</el-button>
          </el-tooltip>
          <el-tooltip content="删除" placement="top">
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
      <el-form ref="surveyRef" :model="form" label-width="100px">
        <el-row>
          <el-col :span="24">
            <el-form-item label="问卷ID" prop="surveyId">
              <el-input disabled v-model="form.survey_id" placeholder="请输入标题"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="题目内容" prop="text">
              <el-input v-model="form.text" placeholder="请输入描述"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="题目类型" prop="type">
              <el-input v-model="form.type" placeholder="请输入描述"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item
                v-for="(option, index) in dynamicOptions.options"
                :label="'选项' + index"
                :key="option.key"
                :prop="'option.' + index + '.value'">
              <el-input v-model="option.value"></el-input>
              <el-button @click.prevent="removeOption(option)">删除</el-button>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="排序" prop="order">
              <el-input-number v-model="form.order" placeholder="请输入排序"></el-input-number>
            </el-form-item>
            <el-button @click="addOption">新增选项</el-button>
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


  </div>

</template>

<script setup>
import {reactive, ref, toRefs, watch} from 'vue'
import {list, add, del, get} from "../api/question.js";
import {update} from "../api/survey.js";

const props = defineProps({
  surveyId: Number,
})

const open = ref(false)
const title = ref('')
const questionList = ref([])
const total = ref(0)
const dynamicOptions = ref({
  options: [{
    id: 0,
    question_id: 0,
    label: '',
    value: ''
  }]
})

const data = reactive({
  form: {},
  queryParams: {
    pageNum: 1,
    pageSize: 10,
  },
});

const {queryParams, form} = toRefs(data);

// 监控props.surveyId的变化
watch(() => props.surveyId, (newValue, oldValue) => {
  getList()
  // 在这里可以执行你的逻辑
});

function addOption() {
  dynamicOptions.value.options.push({
        value: '',
        key: Date.now()
      }
  )
}

function removeOption(op) {
  const index = dynamicOptions.value.options.indexOf(op);
  if (index !== -1) {
    dynamicOptions.value.options.splice(index, 1)
  }
}

function getList() {
  list(queryParams.value, props.surveyId).then(res => {
    questionList.value = res.data.list
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
    survey_id: props.surveyId,
    text: '',
    type: '',
    options: [],
    order: '',
  }

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
  form.value.options = dynamicOptions.value.options
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

</style>
