import Vue from 'vue';
import App from './App.vue';
import VueRouter from "vue-router";
import axios from "axios";
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

import router from "./router";

Vue.config.productionTip = false;

Vue.use(VueRouter);
Vue.use(ElementUI);
Vue.prototype.$http = axios;

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
