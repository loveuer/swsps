import VueRouter from 'vue-router'


const routes = [
    { path: "/", name: "home", component: () => import ("./components/Home.vue") },
]

export default new VueRouter({
    mode: "history",
    routes: routes,
});