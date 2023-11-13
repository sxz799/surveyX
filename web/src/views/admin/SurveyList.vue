<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <div class="logo">
          <img src="@/assets/favicon.png" class="logo-img" alt="SurveyX logo">
          <span class="main-title">SurveyX</span>
          <span class="subtitle">——免费的问卷调查系统</span>
        </div>
      </el-header>
      <el-main>
        <el-row :gutter="4">
          <el-col :span="8" :xs="24">
            <div style="margin-bottom: 4px;">
              <el-card>
                <el-form :model="queryParams" size="small" :inline="true">
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
                <el-row :gutter="5" class="mb8">
                  <el-col :span="1.5">
                    <el-button type="primary" plain size="small" :icon="Plus" @click="handleAdd">新建问卷</el-button>
                  </el-col>
                  <el-col :span="1.5">
                    <el-button type="success" plain size="small" :icon="VideoPlay" :disabled="selectedRows.length===0"
                               @click="handleStartCollect">开始收集
                    </el-button>
                  </el-col>
                  <el-col :span="1.5">
                    <el-button type="danger" plain size="small" :disabled="selectedRows.length===0" :icon="VideoPause"
                               @click="handleStopCollect">停止收集
                    </el-button>
                  </el-col>
                  <el-col :span="1.5">
                    <el-button type="success" plain size="small" :disabled="selectedRows.length!==1" :icon="Link"
                               @click="copySurveyLink">复制链接
                    </el-button>
                  </el-col>
                  <el-col :span="1.5">
                    <el-upload
                        :on-success="handleUploadSuccess"
                        :on-error="handleUploadFail"
                        accept=".xlsx"
                        :show-file-list="false"
                        action="/api/admin/survey/import">
                      <el-button color="#555555" size="small" :icon="UploadFilled" plain>上传问卷</el-button>
                    </el-upload>
                  </el-col>
                </el-row>
                <el-table border fit :highlight-current-row="true" :data="surveyList" @row-click="handleClickRow"
                          @selection-change="handleSelectionChange">
                  <el-table-column type="selection" align="center"/>
                  <el-table-column label="序号" width="70" align="center" type="index"/>
                  <el-table-column label="标题" align="center" key="title" prop="title" :show-overflow-tooltip="false"/>
                  <el-table-column label="状态" align="center" width="100" key="title" prop="status"
                                   :show-overflow-tooltip="false">
                    <template #default="scope">
                      <el-tag v-if="scope.row.status === 'new'">初始</el-tag>
                      <el-tag type="success" v-if="scope.row.status === 'collecting'">收集中</el-tag>
                      <el-tag type="danger" v-if="scope.row.status === 'stop'">停止</el-tag>
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" align="center" width="150">
                    <template #default="scope">
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
              </el-card>
            </div>
            <div style="margin-bottom: 4px;">
              <el-card v-if="open">
                <template #header>
                  <span>{{ title }}</span>
                </template>
                <el-form :disabled="title==='详 情'" @click="title='修 改'" ref="surveyRef" :model="form" :inline="true"
                         size="small" :rules="rules">
                  <el-form-item label="标题" prop="title">
                    <el-input type="textarea" :autosize="{ minRows: 1, maxRows: 3 }" v-model="form.title"
                              placeholder="请输入标题"></el-input>
                  </el-form-item>
                  <el-form-item label="描述" prop="description">
                    <el-input type="textarea" :autosize="{ minRows: 1, maxRows: 3 }" v-model="form.description"
                              placeholder="请输入描述"></el-input>
                  </el-form-item>
                  <el-form-item label="开始时间" prop="start_time">
                    <el-date-picker
                        v-model="form.start_time"
                        type="datetime"
                        placeholder="选择开始日期时间"
                    />
                  </el-form-item>
                  <el-form-item label="结束时间" prop="end_time">
                    <el-date-picker
                        v-model="form.end_time"
                        type="datetime"
                        placeholder="选择结束日期时间"
                    />
                  </el-form-item>
                  <el-form-item label="填写联系方式" prop="need_contact">
                    <el-select v-model="form.need_contact" placeholder="请选择">
                      <el-option label="是" value="yes"></el-option>
                      <el-option label="否" value="no"></el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="可重复提交" prop="repeat">
                    <el-select v-model="form.repeat" placeholder="请选择">
                      <el-option label="是" value="yes"></el-option>
                      <el-option label="否" value="no"></el-option>
                      <el-option label="更新" value="update"></el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="重复提交检查方式" prop="repeat_check">
                    <el-select v-model="form.repeat_check" placeholder="请选择">
                      <el-option v-if="form.need_contact==='yes'" label="联系方式" value="contact"></el-option>
                      <el-option label="浏览器指纹" value="finger"></el-option>
                    </el-select>
                  </el-form-item>
                  <el-form-item label="水印" prop="water_mark">
                    <el-input type="textarea" :autosize="{ minRows: 1, maxRows: 3 }" v-model="form.water_mark"
                              placeholder="请输入水印"></el-input>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" @click="submitForm(surveyRef)">确定</el-button>
                  </el-form-item>
                </el-form>
              </el-card>
            </div>
          </el-col>
          <el-col :span="16" :xs="24">
            <el-card>
              <el-tabs>
                <el-tab-pane label="题目管理">
                  <QuestionList v-if="openDetails" :surveyId="surveyId"/>
                </el-tab-pane>
                <el-tab-pane label="问卷分析">
                  <Analysis v-if="openDetails" :surveyId="surveyId"/>
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
import {onMounted, reactive, ref, toRefs} from 'vue'
import {list, add, del, update, get} from "@/api/admin/survey.js";
import QuestionList from "./QuestionList.vue";
import Analysis from "@/views/admin/Analysis.vue";
import {
  Delete,
  Plus,
  Search,
  Refresh,
  Link,
  UploadFilled,
  VideoPlay,
  VideoPause
} from "@element-plus/icons";
import useClipboard from 'vue-clipboard3';
import {ElMessage} from "element-plus";

const {toClipboard} = useClipboard();
const surveyRef = ref()
const surveyId = ref('')
const open = ref(false)
const openDetails = ref(false)
const surveyList = ref([])
const total = ref(0)
const title = ref('新 增')
const selectedRows = ref([])
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
  need_contact: [{required: true, message: '请填写', trigger: 'blur'}],
  repeat: [{required: true, message: '请填写', trigger: 'blur'}],
  repeat_check: [{required: true, message: '请填写', trigger: 'blur'}],
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

function reset() {
  form.value = {options: []};
}

function handleAdd() {
  title.value = '新 增'
  reset();
  open.value = true
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

function handleStartCollect() {
  selectedRows.value.forEach(row => {
    row.status = 'collecting'
    update(row).then(res => {
      if (res.success) {
        ElMessage.success(row.title + '发布成功！');
        getList()
      } else {
        ElMessage.error(row.title + '发布失败！');
      }

    })
  })
}

function handleStopCollect() {
  selectedRows.value.forEach(row => {
    row.status = 'stop'
    update(row).then(res => {
      if (res.success) {
        ElMessage.success(row.title + '停止成功！');
        getList()
      } else {
        ElMessage.error(row.title + '停止失败！');
      }

    })
  })

}

function handleSelectionChange(val) {
  selectedRows.value = val
}

function handleClickRow(row) {
  title.value = '详 情'
  surveyId.value = row.id
  openDetails.value = true

  get(row.id).then(res => {
    open.value = true
    form.value = res.data
  })
}

function handleUploadSuccess(response) {
  if (response.success) {
    ElMessage.success('上传成功！');
    getList()
  } else {
    ElMessage.error('上传失败！');
  }
}

function handleUploadFail() {
  ElMessage.error('上传失败！')
}

function copySurveyLink() {
  try {
    //获取url
    let url = 'http://' + window.location.host + '/survey/' + selectedRows.value[0].id
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
            getList()
          } else {
            ElMessage.error('修改失败！');
          }
        })
      } else {
        add(form.value).then(res => {
          if (res.success) {
            ElMessage.success('新增成功！');
            reset()
            getList()
          } else {
            ElMessage.error('新增失败！');
          }
        })
      }
      open.value = false
    }
  })
}

</script>

<style scoped>

.app-container {
  padding: 20px;
}

.mb8 {
  margin-bottom: 8px;
}

.el-header {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  padding: 10px;
  background-color: #f8f8f9;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, .1);
}

.logo {
  display: flex;
  align-items: center;
}

.logo-img {
  height: 80px;
  width: auto;
  margin-right: 0;
}

.main-title {
  color: #303133;
  font-size: 20px;
  font-weight: 600;
}

.subtitle {
  font-size: 13px;
  margin-left: 5px;
  color: #606266;
}

</style>
