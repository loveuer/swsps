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
    // training recorder
    {
        path: "/works/trainingRecorder/add",
        name: "works-training-add",
        component: () => import ("./components/trainingRecorder/TrainingAdd.vue"),
    },
    {
        path: "/works/trainingRecorder/check/:date",
        name: "works-training-check",
        component: () => import ("./components/trainingRecorder/TrainingCheck.vue"),
    },
    // log recorder
    {
        path: "/works/logRecorder/add",
        name: "works-log-add",
        component: () => import ("./components/logRecorder/LogAdd.vue"),
    },
    {
        path: "/works/logRecorder/check/:date",
        name: "works-log-check",
        component: () => import ("./components/logRecorder/LogCheck.vue"),
    },
    // sps recorder
    {
        path: "/works/spsRecorder/search",
        name: "works-sps-search",
        component: () => import ("./components/spsRecorder/SpsSearch.vue"),
    },
    {
        path: "/works/spsRecorder/search/:skey",
        name: "works-sps-search",
        component: () => import ("./components/spsRecorder/SpsSearch.vue"),
    },
    {
        path: "/works/spsRecorder/detail/:spid",
        name: "works-sps-detail",
        component: () => import ("./components/spsRecorder/SpsOne.vue"),
    },
    // malf recorder
    {
        path: "/works/malfRecorder/add",
        name: "works-malf-detail",
        component: () => import ("./components/malfRecorder/MalfAdd.vue"),
    },
    // 工作安排 ...
    {
        path: "/works/workRecorder/duty",
        name: "works-work-duty",
        component: () => import ("./components/workRecorder/WorkDuty.vue"),
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