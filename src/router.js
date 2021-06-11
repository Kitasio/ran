import { createRouter, createWebHistory } from 'vue-router'

import Login from './views/Login.vue'
import Management from './views/Management.vue'
import NotFound from './views/NotFound.vue'
import News from './views/News.vue'
import NewsId from './views/NewsId.vue'

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: Login,
    },
    {
        path: '/management',
        name: 'Management',
        component: Management,
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
        path: '/news/:id',
        name: 'NewsId',
        component: NewsId,
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router