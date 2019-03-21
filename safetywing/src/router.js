import VueRouter from 'vue-router'


const routes = [
    {
        path: "/",
        name: "home",
        component: () => import ("./components/Home.vue"),
    },
    {
        path: "/works",
        name: "works",
        component: () => import ("./components/Works.vue"),
    },
    {
        path: "/works/trainingRecorder/add",
        name: "works-training-add",
        component: () => import ("./components/trainingRecorder/Tadd.vue"),
    },
    {
        path: "/login",
        name: "login",
        component: () => import ("./components/Login.vue"),
    },
    {
        path: "*",
        name: "404",
        component: () => import ("./components/PageNotFound.vue"),
    }
]

export default new VueRouter({
    mode: "history",
    routes: routes,
});