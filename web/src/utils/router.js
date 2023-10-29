// router.js
import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/SurveyList.vue';
import Question from '../views/QuestionList.vue';

const routes = [
    { path: '/', component: Home },
    { path: '/question/:surveyID', name: 'Question',component: Question , props: true},
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
