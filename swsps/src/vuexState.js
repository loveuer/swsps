import Vuex from "vuex";
import Vue from "vue";

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {},
    getters: {}, // just like computed
    mutations: {}, // commit to state
});

export default store;
