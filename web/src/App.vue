

<template>
  <el-config-provider :locale="zhCn">
    <router-link to="/"></router-link>
    <router-view></router-view>
  </el-config-provider>
</template>
<script setup>
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
</script>


