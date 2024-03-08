<template>

  <el-row style="padding-top: 100px">
    <el-col :span="6" :xs="0"/>
    <el-col :span="12" :xs="24">
      <div class="login-container">
        <h2 class="login-title">欢迎登录</h2>
        <el-input :autofocus="true" v-model="username" placeholder="请输入用户名" @keyup.enter="Login"
                  class="login-input"></el-input>
        <el-input v-model="password" placeholder="请输入密码" @keyup.enter="Login" class="login-input"></el-input>
        <el-button type="primary" @click="Login" class="login-button">登录</el-button>
      </div>
    </el-col>
    <el-col :span="6" :xs="0"/>
    <el-button type="primary" @click="LoginWithGithub" class="login-button">GITHUB登录</el-button>
  </el-row>


</template>

<script setup>

import {login, } from '@/api/common/login.js'
import {ref, } from "vue";
import {useRouter} from "vue-router";

const router = useRouter()
const username = ref('admin')
const password = ref('admin')



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
  window.location.href = 'https://github.com/login/oauth/authorize?client_id=ff3f2aa615877bd961e7'
}



</script>

<style scoped>
.login-container {

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
  margin: 0 auto;
  width: 100%;
}
</style>
