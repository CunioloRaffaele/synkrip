import { createRouter, createWebHistory } from 'vue-router'
import mainPage from '../pages/mainPage.vue'
import LandingPage from '../pages/landing.vue'
import OsLibs from '../pages/osLibs.vue'
import { Transition } from 'vue'
import { EventsOn } from '../../wailsjs/runtime/runtime'

const routes = [
    {
        path: '/',
        name: 'LandingPage',
        component: LandingPage,
    },
    {
        path: '/',
        name: 'mainPage',
        component: mainPage,
    },
    {
        path: '/',
        name: 'OsLibs',
        component: OsLibs,
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

EventsOn("osLib", () => {
    console.log("osLib")
    router.push({ name: 'OsLibs' })
})

export default router