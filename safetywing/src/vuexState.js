import Vuex from "vuex";
import Vue from "vue";

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        customColor: "#1890ff",
        logined: false,
        user: {
            id: "0",
            username: "",
            realname: "",
            profile_icon: "",
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

export default store;