// router.js
import { createRouter, createWebHistory } from 'vue-router';

import SurveyPage from '../views/SurveyPage.vue';
const routes = [
    { path: '/index', component: SurveyPage },
    { path: '/survey', name: 'SurveyPage',component: SurveyPage},
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
