import VueRouter from 'vue-router'


const routes = [
    {
        path: "/",
        name: "home",
        component: () => import ("./components/Home.vue"),
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