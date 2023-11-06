<template>
  <div class="app-container">
    <el-row>
      <el-col :span="3" :xs="0"/>
      <el-col :span="18" :xs="24">
        <el-form :model="queryParams" size="small" :inline="true" label-width="68px">
          <el-form-item label="标题" prop="roleName">
            <el-input
                v-model="queryParams.title"
                placeholder="请输入标题"
                clearable
                style="width: 240px"
                @keyup.enter="handleQuery"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" :icon="Search" @click="handleQuery">搜索</el-button>
            <el-button :icon="Refresh" @click="resetQuery">重置</el-button>
          </el-form-item>
        </el-form>
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
                type="primary"
                plain
                :icon="Plus"
                @click="handleAdd"
            >新建问卷
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-upload
                :on-success="function (response) {
                  if (response.success) {
                    ElMessage.success('上传成功！');
                    getList()
                  } else {
                    ElMessage.error('上传失败！');
                  }
                }"
                :on-error="function (response) {
                  ElMessage.error('上传失败！');
                }"
                accept=".xlsx"
                :show-file-list="false"
                action="/api/admin/survey/import">
              <el-button :icon="UploadFilled" type="primary">上传问卷</el-button>
            </el-upload>
          </el-col>
        </el-row>

        <el-table border fit :data="surveyList">
          <!--          <el-table-column type="selection" align="center"/>-->
          <el-table-column label="序号" width="70" align="center" type="index"/>
          <el-table-column label="标题" align="center" key="title" prop="title" :show-overflow-tooltip="false">

          </el-table-column>

          <el-table-column label="状态" align="center" key="status" prop="status"
                           :show-overflow-tooltip="true">
            <template #default="scope">
              <el-tag type="success" v-if="scope.row.status === 'yes'">启用</el-tag>
              <el-tag type="danger" v-if="scope.row.status === 'no'">禁用</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="开始时间" width="110" :editable="false" align="center" :formatter="dateTimeFormat" key="start_time"
                           prop="start_time"/>
          <el-table-column label="结束时间" width="110" :editable="false" align="center" :formatter="dateTimeFormat" key="end_time"
                           prop="end_time"/>
          <el-table-column label="填写联系方式" width="90" align="center" key="need_contact" prop="need_contact">
            <template #default="scope">
              <el-tag v-if="scope.row.need_contact === 'yes'">是</el-tag>
              <el-tag type="danger" v-if="scope.row.need_contact === 'no'">否</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="重复提交" width="90" align="center" key="repeat" prop="repeat">
            <template #default="scope">
              <el-tag v-if="scope.row.repeat === 'yes'">是</el-tag>
              <el-tag type="danger" v-if="scope.row.repeat === 'no'">否</el-tag>
              <el-tag type="success" v-if="scope.row.repeat === 'yes_but_update'">更新</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="重复提交检查方式" width="105" align="center" key="repeat_check" prop="repeat_check">
            <template #default="scope">
              <el-tag type="warning" v-if="scope.row.repeat_check === 'contact'">联系方式</el-tag>
              <el-tag type="info" v-if="scope.row.repeat_check === 'finger'">浏览器指纹</el-tag>
              <el-tag type="info" v-if="scope.row.repeat_check === ''">无</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="150">
            <template #default="scope">
              <el-button link type="primary" @click="handleEdit(scope.row)" :icon="Edit">修改</el-button>
              <el-button link type="success" @click="handleQuestion(scope.row.id)" :icon="Tools">配置题目</el-button>
              <el-popconfirm
                  confirm-button-text="确定"
                  cancel-button-text="取消"
                  @confirm="handleDelete(scope.row)"
                  title="确定要删除吗?">
                <template #reference>
                  <el-button link type="danger" :icon="Delete">删除</el-button>
                </template>
              </el-popconfirm>
            </template>
          </el-table-column>
          <el-table-column label="更多操作" align="center" width="150">
            <template #default="scope">
              <el-button link type="warning" @click="copySurveyLink(scope.row)" :icon="Link">问卷地址</el-button>
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
          <el-form ref="surveyRef" :model="form" :rules="rules" label-width="25%">
            <el-row :gutter="10">
              <el-col :span="24">
                <el-form-item label="标题" prop="title">
                  <el-input type="textarea" :autosize="{ minRows: 1, maxRows: 3 }" v-model="form.title"
                            placeholder="请输入标题"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="描述" prop="description">
                  <el-input type="textarea" :autosize="{ minRows: 1, maxRows: 3 }" v-model="form.description"
                            placeholder="请输入描述"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="状态" prop="status">
                  <el-select v-model="form.status" placeholder="请选择状态">
                    <el-option label="启用" value="yes"></el-option>
                    <el-option label="禁用" value="no"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="填写联系方式" prop="need_contact">
                  <el-select v-model="form.need_contact" placeholder="请选择是否需要填写联系方式">
                    <el-option label="是" value="yes"></el-option>
                    <el-option label="否" value="no"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="可重复提交" prop="repeat">
                  <el-select v-model="form.repeat" placeholder="请选择是否可重复提交">
                    <el-option label="是" value="yes"></el-option>
                    <el-option label="否" value="no"></el-option>
                    <el-option label="是(更新)" value="yes_but_update"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item v-if="form.repeat!=='yes'" label="重复提交检查方式" prop="repeat_check">
                  <el-select v-model="form.repeat_check" placeholder="请选择重复提交检查方式">
                    <el-option v-if="form.need_contact==='yes'" label="联系方式" value="contact"></el-option>
                    <el-option label="浏览器指纹" value="finger"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="开始时间" prop="start_time">
                  <el-date-picker
                      v-model="form.start_time"
                      type="datetime"
                      placeholder="选择开始日期时间"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="结束时间" prop="end_time">
                  <el-date-picker
                      v-model="form.end_time"
                      type="datetime"
                      placeholder="选择结束日期时间"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="水印" prop="water_mark">
                  <el-input type="textarea" :autosize="{ minRows: 1, maxRows: 3 }" v-model="form.water_mark"
                            placeholder="请输入水印"></el-input>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
          <template #footer>
            <div class="dialog-footer">
              <el-button type="primary" @click="submitForm(surveyRef)">确 定</el-button>
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
      <el-col :span="3" :xs="0"/>
    </el-row>
  </div>
</template>

<script setup>
import {computed, onMounted, reactive, ref, toRefs} from 'vue'
import {list, add, del, update, get} from "@/api/admin/survey.js";
import QuestionList from "./QuestionList.vue";
import {Delete, Edit, Plus, Tools, Search, Refresh, Link, UploadFilled} from "@element-plus/icons";
import useClipboard from 'vue-clipboard3';
import {ElMessage} from "element-plus";
import {useRouter} from "vue-router";

const router = useRouter()

const {toClipboard} = useClipboard();
const surveyRef = ref()
const surveyId = ref('')
const open = ref(false)
const openDetails = ref(false)

const title = ref('')
const surveyList = ref([])
const total = ref(0)
const dialogWidth = computed(() => {
  return window.innerWidth > 768 ? '60%' : '90%'
})

const data = reactive({
  form: {
    options: []
  },
  queryParams: {
    pageNum: 1,
    pageSize: 10,
    title: '',
  },
});

const {queryParams, form} = toRefs(data);


const rules = ({
  title: [{required: true, message: '请填写', trigger: 'blur'}],
  description: [{required: true, message: '请填写', trigger: 'blur'}],
  status: [{required: true, message: '请填写', trigger: 'change'}],
  need_contact: [{required: true, message: '请填写', trigger: 'change'}],
  repeat: [{required: true, message: '请填写', trigger: 'change'}],
  repeat_check: [{required: true, message: '请填写', trigger: 'change'}],
  start_time: [{required: true, message: '请填写', trigger: 'blur'}],
  end_time: [{required: true, message: '请填写', trigger: 'blur'}],
})

onMounted(() => {
  getList()
})

function handleQuery() {
  queryParams.value.pageNum = 1;
  getList();
}

function resetQuery() {
  queryParams.value.title = '';
  queryParams.value.pageNum = 1;
  queryParams.value.pageSize = 10;
  getList();
}

function getList() {
  list(queryParams.value).then(res => {
    if (res.success) {
      surveyList.value = res.data.list
      total.value = res.data.total
    } else {
      //路由到登录页面
      router.push({path: '/login'})
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

function handleAdd() {
  reset();
  open.value = true
  title.value = '新增'

}

function reset() {
  form.value = {};
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
    if (res.success) {
      ElMessage.success('删除成功！');
      getList()
    } else {
      ElMessage.error('删除失败！');
    }

  })
}

function handleImport() {

}

function copySurveyLink(row) {
  try {
    //获取url
    let url = 'http://' + window.location.host + '/survey/' + row.id
    toClipboard(url)
    ElMessage.success('复制成功！');
  } catch (e) {
    ElMessage.error('复制失败！');
  }

}

function submitForm(elForm) {
  elForm.validate((valid) => {
    if (valid) {
      if (form.value.id) {
        update(form.value).then(res => {
          if (res.success) {
            ElMessage.success('修改成功！');
            open.value = false
            getList()
          } else {
            ElMessage.error('修改失败！');
          }

        })
      } else {
        add(form.value).then(res => {
          if (res.success) {
            ElMessage.success('新增成功！');
            open.value = false
            getList()
          } else {
            ElMessage.error('新增失败！');
          }
        })
      }
    }
  })
}

function dateTimeFormat(row, column, cellValue) {
  if (cellValue === undefined || cellValue === null || cellValue === '') {
    return ''
  }
  const date = new Date(cellValue)
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  const seconds = String(date.getSeconds()).padStart(2, '0');
  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}


</script>

<style scoped>

.app-container {
  padding: 20px;
}

.mb8 {
  margin-bottom: 8px;
}

</style>
