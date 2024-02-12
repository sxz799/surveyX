<template>

  <div class="app-container">
    <h1>欢迎使用SurveyX</h1>
    <el-button type="danger" @click="Login">登录后台</el-button>
  </div>
</template>

<script setup>

import {useRouter} from 'vue-router'

const router = useRouter()

const Login = () => {
  router.push('/login')
}

</script>

<style scoped>

.app-container {
  padding: 20px;
  text-align: center;
  font-size: 20px;
  color: #333;
  background-color: #f5f5f5;
  border-radius: 4px;
  margin: 20px;
  border: 1px solid #ebeef5;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  transition: all 0.3s;

}

</style>
