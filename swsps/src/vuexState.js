import Vuex from "vuex";
import Vue from "vue";

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        user: {
            id: 0,
        },
    },
    getters: {}, // just like computed
    mutations: {
        setUser: function(state, newUser) {
            state.user = newUser;
        },
        
    }, // commit to state
});

export default store;
