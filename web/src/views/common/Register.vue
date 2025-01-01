<template>
  <el-dialog
      v-model="open"
      title="用户注册"
      width="400px"
      :close-on-click-modal="false"
      :show-close="false"
      center
  >
    <el-form
        :model="model"
        :rules="rules"
        ref="formRef"
        label-width="80px"
        status-icon
    >
      <el-form-item label="用户名" prop="username">
        <el-input
            v-model="model.username"
            placeholder="请输入用户名"
            prefix-icon="User"
        ></el-input>
      </el-form-item>

      <el-form-item label="密码" prop="password">
        <el-input
            v-model="model.password"
            type="password"
            placeholder="请输入密码"
            prefix-icon="Lock"
            show-password
        ></el-input>
      </el-form-item>

      <el-form-item label="昵称" prop="nickname">
        <el-input
            v-model="model.nickname"
            placeholder="请输入昵称"
            prefix-icon="User"
        ></el-input>
      </el-form-item>

      <el-form-item label="邮箱" prop="email">
        <el-input
            v-model="model.email"
            placeholder="请输入邮箱"
            prefix-icon="Message"
        ></el-input>
      </el-form-item>

      <el-form-item label="手机号" prop="phone">
        <el-input
            v-model="model.phone"
            placeholder="请输入手机号"
            prefix-icon="Phone"
        ></el-input>
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="open = false">取 消</el-button>
        <el-button type="primary" @click="handleRegister">注 册</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import {register} from '@/api/user/user.js'
import {ref} from "vue"
import {ElMessage} from "element-plus"
import {User, Lock, Message, Phone} from '@element-plus/icons-vue'

const open = ref(false)
const formRef = ref(null)

const model = ref({
  username: '',
  password: '',
  nickname: '',
  email: '',
  phone: ''
})

// 表单验证规则
const rules = {
  username: [
    {required: true, message: '请输入用户名', trigger: 'blur'},
    {min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur'}
  ],
  password: [
    {required: true, message: '请输入密码', trigger: 'blur'},
    {min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur'}
  ],
  nickname: [
    {required: true, message: '请输入昵称', trigger: 'blur'}
  ],
  email: [
    {required: true, message: '请输入邮箱地址', trigger: 'blur'},
    {type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change']}
  ],
  phone: [
    {required: true, message: '请输入手机号', trigger: 'blur'},
    {pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号码', trigger: 'blur'}
  ]
}

function Open() {
  open.value = true
  // 打开对话框时重置表单
  if (formRef.value) {
    formRef.value.resetFields()
  }
}

async function handleRegister() {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const res = await register(model.value)
        if (res.success) {
          ElMessage.success(res.message)
          open.value = false
        } else {
          ElMessage.error(res.message)
        }
      } catch (error) {
        ElMessage.error('注册失败，请稍后重试')
      }
    }
  })
}

defineExpose({
  Open
})
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 20px;
}

:deep(.el-form-item) {
  margin-bottom: 22px;
}

:deep(.el-input) {
  --el-input-hover-border-color: var(--el-primary-color);
}
</style>
