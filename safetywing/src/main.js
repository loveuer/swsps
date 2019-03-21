import Vue from 'vue';
import App from './App.vue';
import VueRouter from "vue-router";
import axios from "axios";
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

import store from "./vuexState";
import router from "./router";

Vue.use(VueRouter);
Vue.use(ElementUI);

Vue.prototype.$http = axios;
Vue.prototype.$getcookie = getCookie;
Vue.prototype.$setcookie = setCookie;
// Vue.prototype.$route = router;

Vue.config.productionTip = false;


new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app');

function getCookie(name) {
  var match = document.cookie.match(new RegExp('(^| )' + name + '=([^;]+)'));
  if (match) {
    return match[2];
  } else {
    return "";
  }
};

function setCookie(name,value,mins) {
  var expires = "";
  if (mins) {
      var date = new Date();
      date.setTime(date.getTime() + (mins*60*1000));
      expires = "; expires=" + date.toUTCString();
  }
  document.cookie = name + "=" + (value || "")  + expires + "; path=/";
};
