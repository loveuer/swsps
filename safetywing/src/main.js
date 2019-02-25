import Vue from 'vue'
import App from './App.vue'
import VueRouter from "vue-router"
import axios from "axios"


import router from "./router"
import AntForm from "ant-design-vue/lib/form";
import "ant-design-vue/dist/antd.css";

Vue.use(VueRouter)
Vue.prototype.$http = axios
Vue.prototype.$cookies = getCookie
Vue.prototype.$form = AntForm


Vue.config.productionTip = false



Vue.component(AntForm.name, AntForm);

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
