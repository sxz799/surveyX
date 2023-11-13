

<template>
  <el-table fit border size="small" :data="answerDetails">
    <el-table-column width="50" align="center" v-if="questionType!=='text'" label="选项" prop="label"/>
    <el-table-column v-if="questionType!=='text'" label="备注" prop="ext_msg"/>
    <el-table-column v-if="questionType==='text'" label="答案" prop="content"/>
    <el-table-column label="浏览器指纹" prop="finger"/>
    <el-table-column label="联系方式" prop="contact"/>
  </el-table>
</template>


<script setup>


const props = defineProps({
  answerDetails: Array,
  questionType: String,
})

</script>

<style scoped>



</style>
