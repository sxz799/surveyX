// router.js
import { createRouter, createWebHistory } from 'vue-router';

import Main from '../views/admin/Main.vue';
import Index from '../views/user/Index.vue';
import SurveyPage from '../views/user/SurveyPage.vue';
import Login from '../views/common/Login.vue';
import LoginOauth from '../views/common/LoginOauth.vue';
const routes = [
    { path: '/', name: 'Index',component: Index},
    { path: '/login', name: 'Login',component: Login},
    { path: '/oauth-github', name: 'LoginOauth',component: LoginOauth},
    { path: '/admin', name: 'Main',component: Main},
    { path: '/survey/:id', name: 'SurveyPage',component: SurveyPage},
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
