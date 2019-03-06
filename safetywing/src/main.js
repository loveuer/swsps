import Vue from 'vue';
import App from './App.vue';
import VueRouter from "vue-router";
import axios from "axios";
import Vuex from "vuex";


import router from "./router";

Vue.use(VueRouter)
Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        customColor: "#1890ff",
        user: {
            id: "0",
            username: "",
            realname: "",
            profileIcon: "",
        },
    },
    mutations: {
        chgColor(state, newColor) {
            state.customColor = newColor;
        },
        chgUser(state, newUser) {
            state.user = newUser;
        }
    },
});

Vue.prototype.$http = axios
Vue.prototype.$getcookie = getCookie
Vue.prototype.$setcookie = setCookie

Vue.config.productionTip = false


new Vue({
  store,
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

function setCookie(name,value,mins) {
  var expires = "";
  if (mins) {
      var date = new Date();
      date.setTime(date.getTime() + (mins*60*1000));
      expires = "; expires=" + date.toUTCString();
  }
  document.cookie = name + "=" + (value || "")  + expires + "; path=/";
}
