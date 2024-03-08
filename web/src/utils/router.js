// router.js
import { createRouter, createWebHistory } from 'vue-router';

import SurveyListAdmin from '../views/admin/SurveyList.vue';
import Index from '../views/user/Index.vue';
import SurveyPage from '../views/user/SurveyPage.vue';
import Login from '../views/common/Login.vue';
import LoginOauth from '../views/common/LoginOauth.vue';
const routes = [
    { path: '/', name: 'Index',component: Index},
    { path: '/login', name: 'Login',component: Login},
    { path: '/oauth-github', name: 'LoginOauth',component: LoginOauth},
    { path: '/admin', name: 'SurveyListAdmin',component: SurveyListAdmin},
    { path: '/survey/:id', name: 'SurveyPage',component: SurveyPage},
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
