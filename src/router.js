import { createRouter, createWebHistory } from 'vue-router'

import Login from './views/Login.vue'
import Management from './views/Management.vue'
import ManagementId from './views/ManagementId.vue'
import NotFound from './views/NotFound.vue'
import News from './views/News.vue'
import NewsId from './views/NewsId.vue'
import Ads from './views/Ads.vue'
import AdId from './views/AdId.vue'
import Units from './views/Units.vue'
import UnitId from './views/UnitId.vue'
import Council from './views/Council.vue'
import Partners from './views/Partners.vue'
import About from './views/About.vue'
import Documents from './views/Documents.vue'
import Home from './views/Home.vue'

const routes = [
    {
        path: '/',
        redirect: '/home'
    },
    {
        path: '/home',
        name: 'home',
        component: Home,
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
    },
    {
        path: '/about',
        name: 'About',
        component: About,
    },
    {
        path: '/management',
        name: 'Management',
        component: Management,
    },
    {
        path: '/management/:id',
        name: 'ManagementId',
        component: ManagementId,
    },
    {
        path: '/:catchAll(.*)*',
        name: 'NotFound',
        component: NotFound,
    },
    {
        path: '/news',
        name: 'News',
        component: News,
    },
    {
        path: '/documents',
        name: 'Documents',
        component: Documents,
    },
    {
        path: '/news/:id',
        name: 'NewsId',
        component: NewsId,
    },
    {
        path: '/ads',
        name: 'Ads',
        component: Ads,
    },
    {
        path: '/ads/:id',
        name: 'AdId',
        component: AdId,
    },
    {
        path: '/units',
        name: 'Units',
        component: Units,
    },
    {
        path: '/units/:id',
        name: 'UnitId',
        component: UnitId,
    },
    {
        path: '/council',
        name: 'Council',
        component: Council,
    },
    {
        path: '/partners',
        name: 'Partners',
        component: Partners,
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router