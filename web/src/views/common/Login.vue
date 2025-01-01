<template>
  <div class="login-page">
    <div class="login-container">
      <div class="login-box">
        <h2 class="title">欢迎使用 SurveyX</h2>

        <el-form :model="model" :rules="rules" ref="formRef" label-position="top" size="large">
          <el-form-item prop="username">
            <el-input v-model="model.username" placeholder="请输入用户名" :prefix-icon="User"
                      @keyup.enter="handleLogin"></el-input>
          </el-form-item>

          <el-form-item prop="password">
            <el-input v-model="model.password" type="password" placeholder="请输入密码" :prefix-icon="Lock"
                      show-password
                      @keyup.enter="handleLogin"></el-input>
          </el-form-item>

          <div class="options">
            <el-checkbox v-model="rememberMe">记住我</el-checkbox>
            <el-link type="primary" :underline="false" @click="handleRestPasswrd">忘记密码？</el-link>
          </div>

          <div class="buttons">
            <el-button type="primary" :loading="loading" @click="handleLogin" class="login-btn">登 录</el-button>

            <el-button @click="handleRegister" class="register-btn">注 册</el-button>
          </div>

          <div class="divider">
            <span>其他登录方式</span>
          </div>

          <div class="social-login">
            <el-button type="primary" @click="loginWithGithub" class="github-btn" :icon="Platform">
              Github 登录
            </el-button>
          </div>
        </el-form>
      </div>
    </div>

    <Register ref="registerRef"></Register>
    <RestPasswrd ref="restPasswrdRef"></RestPasswrd>
  </div>
</template>

<script setup>
import {ref} from "vue"
import {ElMessage} from "element-plus"
import {User, Lock, Platform} from '@element-plus/icons-vue'
import {useRouter} from 'vue-router'
import {login, getGithubLoginUrl} from '@/api/common/common.js'
import Register from './Register.vue'
import RestPasswrd from './RestPasswrd.vue'

const router = useRouter()
const formRef = ref(null)
const registerRef = ref(null)
const restPasswrdRef = ref(null)
const loading = ref(false)
const rememberMe = ref(false)

const model = ref({
  username: '',
  password: ''
})

const rules = {
  username: [
    {required: true, message: '请输入用户名', trigger: 'blur'},
  ],
  password: [
    {required: true, message: '请输入密码', trigger: 'blur'},
  ]
}

async function handleLogin() {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        loading.value = true
        const res = await login({
          ...model.value,
          remember_me: rememberMe.value
        })
        if (res.success) {
          ElMessage.success(res.message || '登录成功')
          localStorage.setItem('token', res.data.token)
          localStorage.setItem('userInfo', JSON.stringify(res.data.userInfo))
          router.push('/admin')
        } else {
          ElMessage.error(res.message || '登录失败')
        }
      } catch (error) {
        ElMessage.error('登录失败，请稍后重试')
      } finally {
        loading.value = false
      }
    }
  })
}

function handleRegister() {
  registerRef.value?.Open()
}

function handleRestPasswrd() {
  restPasswrdRef.value?.Open()
}

async function loginWithGithub() {
  try {
    const res = await getGithubLoginUrl()
    if (res.success) {
      window.location.href = res.data
    } else {
      ElMessage.error(res.message)
    }
  } catch (error) {
    ElMessage.error('获取Github登录链接失败')
  }
}
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--el-bg-color-page);
}

.login-container {
  width: 100%;
  padding: 20px;
}

.login-box {
  max-width: 400px;
  margin: 0 auto;
  padding: 40px;
  background: var(--el-bg-color);
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.title {
  text-align: center;
  color: var(--el-text-color-primary);
  margin-bottom: 30px;
  font-size: 24px;
}

.options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.buttons {
  display: flex;
  gap: 12px;
  margin-bottom: 24px;
}

.login-btn,
.register-btn,
.github-btn {
  flex: 1;
  height: 40px;
}

.divider {
  display: flex;
  align-items: center;
  margin: 24px 0;
  color: var(--el-text-color-secondary);
}

.divider::before,
.divider::after {
  content: '';
  flex: 1;
  height: 1px;
  background: var(--el-border-color-lighter);
}

.divider span {
  padding: 0 12px;
  font-size: 14px;
}

.social-login {
  display: flex;
  justify-content: center;
}

.github-btn {
  width: 100%;
  background-color: #24292e;
  border-color: #24292e;
}

.github-btn:hover {
  background-color: #2c3238;
  border-color: #2c3238;
}

:deep(.el-input) {
  --el-input-hover-border-color: var(--el-primary-color);
}

:deep(.el-form-item) {
  margin-bottom: 24px;
}

@media (max-width: 768px) {
  .login-box {
    padding: 30px 20px;
  }

  .buttons {
    flex-direction: column;
  }
}
</style>
