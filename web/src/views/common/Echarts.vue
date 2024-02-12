<template>
  <v-chart class="chart" :option="option"/>
</template>

<script setup>
import {use} from "echarts/core";
import {CanvasRenderer} from "echarts/renderers";
import {PieChart} from "echarts/charts";
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent
} from "echarts/components";
import VChart, {THEME_KEY} from "vue-echarts";
import {ref, provide} from "vue";
import {analysis as QuestionAnalysis} from "@/api/admin/question.js";
import {ElMessage} from "element-plus";

use([
  CanvasRenderer,
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent
]);

provide(THEME_KEY, "light");

const props = defineProps({
  answerId: Number,
  answerOptions: Array
})

const option = ref({
  // title: {
  //   text: "题目分析",
  //   left: "center"
  // },
  tooltip: {
    trigger: "item",
    formatter: "{a} <br/>{b} : {c} ({d}%)"
  },
  legend: {
    orient: "vertical",
    left: "left",
    data: [
      // "Direct", "Email", "Ad Networks", "Video Ads", "Search Engines"
    ]
  },
  series: [
    {
      name: "选项详情:",
      type: "pie",
      radius: "55%",
      center: ["50%", "60%"],
      data: [
        // {value: 335, name: "Direct"},
        // {value: 310, name: "Email"},
        // {value: 234, name: "Ad Networks"},
        // {value: 135, name: "Video Ads"},
        // {value: 1548, name: "Search Engines"}
      ],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: "rgba(0, 0, 0, 0.5)"
        }
      }
    }
  ]
});

// 接受父组件传来的data
function getQuestionAnalysis(question_id) {
  QuestionAnalysis(question_id).then(res => {
    if (res.success) {
      option.value.legend.data = props.answerOptions.map(obj => obj.label + ":" + obj.value)
      option.value.series[0].data = res.data.label_info.map(obj => {
        return {value: obj.count, name: obj.label + ":" + props.answerOptions.find(item => item.label === obj.label).value}
      })

      for (let datum of option.value.legend.data) {
        if (!option.value.series[0].data.find(item => item.name === datum)) {
          option.value.series[0].data.push({value: 0, name: datum})
        }
      }


    } else {
      ElMessage.error(res.message)
    }
  })
}

getQuestionAnalysis(props.answerId)


</script>

<style scoped>
.chart {
  height: 200px;
}
</style>
