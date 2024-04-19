

<template>
  <div class="logo">
    <img src="@/assets/favicon.png" class="logo-img" alt="SurveyX logo">
    <span class="main-title">SurveyX</span>
    <span class="subtitle">——免费的问卷调查系统</span>
  </div>
  <div class="logout-button">
    <el-button @click="Logout" >退出登录</el-button>
  </div>
</template>

<script setup>


import {logout} from "@/api/common/login.js";
import {ElMessage} from "element-plus";
import router from "@/utils/router.js";

function Logout() {
  logout().then(res => {
    if (res.success) {
      ElMessage.success(res.message);
      localStorage.removeItem('token')
      router.push('/login')
    } else {
      ElMessage.error('退出失败');
    }
  })
}

</script>

<style scoped>

.logo {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
}

.logout-button {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
  top: 0;
  right: 0;
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