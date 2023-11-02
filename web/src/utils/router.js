// router.js
import { createRouter, createWebHashHistory } from 'vue-router';

import SurveyListAdmin from '../views/admin/SurveyList.vue';
import SurveyListUser from '../views/user/SurveyList.vue';
import SurveyPage from '../views/user/SurveyPage.vue';
const routes = [
    { path: '/', name: 'SurveyListUser',component: SurveyListUser},
    { path: '/admin', name: 'SurveyListAdmin',component: SurveyListAdmin},
    { path: '/survey/:id', name: 'SurveyPage',component: SurveyPage},
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;
