import VueRouter from 'vue-router'


const routes = [
    {
        path: "/",
        name: "home",
        component: () => import ("./components/Home.vue"),
    },
    {
        path: "/ydd/salary",
        name: "salary",
        component: () => import ("./components/ydd-tools/salary.vue"),
    },
    {
        path: "*",
        name: "404",
        component: () => import ("./components/u-cmpnts/404.vue"),
    }
]

export default new VueRouter({
    mode: "history",
    routes: routes,
});