import VueRouter from 'vue-router'


const routes = [
    {
        path: "/",
        name: "home",
        component: () => import ("./components/home/Home.vue"),
    },
    {
        path: "/login",
        name: "login",
        component: () => import ("./components/login/Login.vue"),
    },
    {
        path: "/search",
        name: "search",
        component: () => import ("./components/search/Search.vue"),
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