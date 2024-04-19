<template>
  <div class="logo">
    <img src="@/assets/favicon.png" class="logo-img" alt="SurveyX logo">
    <span class="main-title">SurveyX</span>
    <span class="subtitle">——免费的问卷调查系统</span>
  </div>
  <div class="logout-button">


    <el-dropdown>
      <span class="el-dropdown-link">
        欢迎您,{{ userInfo.nickname }}
        <el-icon>
          <arrow-down />
        </el-icon>
      </span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item @click="visable = true">修改密码</el-dropdown-item>
          <el-dropdown-item @click="Logout">退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>


  </div>
  <el-dialog v-model="visable" title="修改密码" width="20%" append-to-body>
    <el-form label-width="auto">
      <el-form-item label="新密码" prop="newPwd">
        <el-input v-model="newPwd" placeholder="请输入新密码"></el-input>
      </el-form-item>
      <el-form-item label="重复密码" prop="newPwd2">
        <el-input v-model="newPwd2" placeholder="请重复输入新密码"></el-input>
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button type="primary" @click="changePwd">确 定</el-button>
        <el-button @click="visable = false">取 消</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>


import { logout } from "@/api/common/common.js";
import { ElMessage } from "element-plus";
import router from "@/utils/router.js";
import { changPwd } from "@/api/common/common.js";
import { onMounted, ref } from "vue";
import { ArrowDown } from "@element-plus/icons";

const userInfo = ref({})

const visable = ref(false)
const newPwd = ref('')
const newPwd2 = ref('')

onMounted(() => {
  userInfo.value = JSON.parse(localStorage.getItem("userInfo"))
})

function changePwd() {
  changPwd({ "id": userInfo.value.id, "password": newPwd.value }).then(res => {
    if (res.success) {
      ElMessage.success(res.message)
      visable.value = false
    } else {
      ElMessage.error(res.message)
    }

  })
}

function Logout() {
  logout().then(res => {
    if (res.success) {
      ElMessage.success(res.message);
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
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

.el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
</style>