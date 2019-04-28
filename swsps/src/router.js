import VueRouter from 'vue-router'


const routes = [
    {
        path: "/",
        name: "home",
        component: () => import ("./components/home/Home.vue"),
    },
    {
        path: "*",
        name: "404",
        component: () => import ("./components/404.vue"),
    }
]

export default new VueRouter({
    mode: "history",
    routes: routes,
});