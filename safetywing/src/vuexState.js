import Vuex from "vuex";
import Vue from "vue";

import mockdata from "./mockdatas"

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
        spsOne: {
            id: '0',
            name: 'null',
            pn: 'null',
            sn: 'null',
            sim: 'null',
        },
        mockhis: mockdata.mockhis,
    },
    mutations: {
        chgColor(state, newColor) {
            state.customColor = newColor;
        },
        chgUser(state, newUser) {
            state.user = newUser;
        },
        updateSpsOne(state, newsp) {
            state.spsOne = newsp;
        },
    },
});

export default store;