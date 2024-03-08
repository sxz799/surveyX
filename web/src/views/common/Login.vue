<template>

  <el-row style="padding-top: 100px">
    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <div class="login-container">
        <h2 class="login-title">欢迎使用 SurveyX</h2>
        <el-input :autofocus="true" v-model="username" placeholder="请输入用户名" @keyup.enter="Login"
                  class="login-input"></el-input>
        <el-input v-model="password" placeholder="请输入密码" @keyup.enter="Login" class="login-input"></el-input>
        <el-row>


          <el-col :span="12" :xs="24">
            <el-button type="primary" @click="Login" class="login-button">登录</el-button>

          </el-col>
          <el-col :span="12" :xs="24">

            <el-button type="primary" @click="LoginWithGithub" class="login-button">GITHUB登录</el-button>
          </el-col>
        </el-row>
      </div>
    </el-col>
    <el-col :span="6" :xs="0"/>
  </el-row>


</template>

<script setup>

import {login, getGithubLoginUrl} from '@/api/common/login.js'
import {ref,} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()
const username = ref('')
const password = ref('')


function Login() {
  login({username: username.value, password: password.value}).then(res => {
    if (res.success) {
      ElMessage.success(res.message)
      router.push({path: '/admin'})
    } else {
      ElMessage.error(res.message)
    }
  })
}

function LoginWithGithub() {
  getGithubLoginUrl().then(res => {
    if (res.success) {
      window.location.href = res.data
    } else {
      ElMessage.error(res.message)
    }
  })

}


</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  margin: 0 auto;
  padding: 50px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  border-radius: 4px;
}

.login-title {
  text-align: center;
  margin-bottom: 20px;
}

.login-input {
  margin-bottom: 20px;
}

.login-button {
  background-color: #24292e; /* GitHub 的主题颜色 */
  color: #ffffff; /* 文字颜色为白色 */
  border: none; /* 无边框 */
  font-size: 16px; /* 字体大小 */
  padding: 10px 20px; /* 内边距 */
  border-radius: 4px; /* 圆角边框 */
  transition: background-color 0.3s ease; /* 过渡效果 */
}

.login-button:hover {
  background-color: #444d56; /* 鼠标悬停时的背景颜色 */
}
</style>
