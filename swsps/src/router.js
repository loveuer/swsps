import VueRouter from 'vue-router'


const routes = [
    { path: "/", name: "home", component: () => import ("./components/home/Home.vue") },
    { path: "/login", name: "login", component: () => import ("./components/login/Login.vue") },
    { path: "/search", name: "search", component: () => import ("./components/search/Search.vue") },
    { path: "/search/:key", name: "search", component: () => import ("./components/search/Search.vue") },
    { path: '/onesp/:spid', name: 'onesp', component: () => import ('./components/onesp/OneSP.vue') },
    { path: "/addsp", name: "addsp", component: () => import ("./components/addsp/AddSP.vue") },
    { path: "/sphis", name: "sphis", component: () => import ("./components/sphis/SPHis.vue") },
    { path: "/status", name: "sphis", component: () => import ("./components/status/Status.vue") },
    { path: '/test', name: 'test', component: () => import ('./components/LightBox.vue') },
    { path: "*", name: "404", component: () => import ("./components/404.vue") }
]

export default new VueRouter({
    mode: "history",
    routes: routes,
});