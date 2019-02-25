import VueRouter from 'vue-router'

import Home from "./components/Home.vue"
import Sps from "./components/Sps.vue"
import Login from "./components/Login.vue"
import PageNotFound from "./components/PageNotFound.vue"

const routes = [
    {
        path: "/",
        name: "home",
        component: Home
    },
    {
        path: "/sps",
        name: "sps",
        component: Sps
    },
    {
        path: "/login",
        name: "login",
        component: Login
    },
    {
        path: "*",
        name: "404",
        component: PageNotFound,
    }
]

export default new VueRouter({
    mode: "history",
    routes: routes,
});