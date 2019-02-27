import Vue from 'vue';
import App from './App.vue';
import VueRouter from "vue-router";
import axios from "axios";


import router from "./router";

Vue.use(VueRouter)

Vue.prototype.$data = {}
Vue.prototype.$http = axios
Vue.prototype.$cookies = getCookie


Vue.config.productionTip = false


new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

function getCookie(name) {
  var match = document.cookie.match(new RegExp('(^| )' + name + '=([^;]+)'));
  if (match) {
    return match[2];
  } else {
    return "";
  }
}
