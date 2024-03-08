<template>

  <h1>Github 登录中,请耐心等待跳转。。。若长时间无反应,请刷新页面！</h1>


</template>

<script setup>

import {loginByGithub} from '@/api/common/login.js'
import { onMounted} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()


onMounted(() => {
  if (router.currentRoute.value.query.code) {
    LoginByGithub(router.currentRoute.value.query.code)
  }
})


function LoginByGithub(code) {
  loginByGithub(code).then(res => {
    if (res.success) {
      ElMessage.success(res.message)
      router.push({path: '/admin'})
    } else {
      ElMessage.error(res.message)
    }
  })
}


</script>

<style scoped>
h2 {
  display: flex;
  color: #0056b3;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
</style>
