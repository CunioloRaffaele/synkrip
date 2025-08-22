import { createRouter, createWebHistory } from 'vue-router'
import mainPage from '../pages/mainPage.vue'
import LandingPage from '../pages/landing.vue'
import OsLibs from '../pages/osLibs.vue'
import Settings from '../pages/settings.vue'
import Eula from '../pages/eula.vue'
import PlaylistDetailPage from '../pages/playlist.vue'
import { Transition } from 'vue'
import { EventsOn } from '../../wailsjs/runtime/runtime'

const routes = [
    {
        path: '/',
        name: 'LandingPage',
        component: LandingPage
    },
    {
        path: '/mainPage',
        name: 'mainPage',
        component: mainPage
    },
    {
        path: '/playlist/:id',
        name: 'PlaylistDetailPage',
        component: PlaylistDetailPage,
        props: true,
    },
    {
        path: '/OsLibs',
        name: 'OsLibs',
        component: OsLibs
    },
    {
        path: '/settings',
        name: 'Settings',
        component: Settings
    },
    {
        path: '/eula',
        name: 'Eula',
        component: Eula
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

EventsOn("osLib", () => {
    console.log("osLib")
    router.push({ name: 'OsLibs' })
})

EventsOn("settings", () => {
    console.log("settings")
    router.push({ name: 'Settings' })
})

EventsOn("showEula", () => {
    console.log("eula")
    router.push({ name: 'Eula' })
})

export default router