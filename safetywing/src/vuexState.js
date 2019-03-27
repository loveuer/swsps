import Vuex from "vuex";
import Vue from "vue";


Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        getNew: false,
        logined: false,
        user: {id: "0"},
        ws: {id: ''},
        todayMessage: {
            log: {
                new: false,
                type: "info",
                logs: [
                    {time:"06:32:21", auth:"李永翔", sim: "5978", log: "Preflight is OK"},
                    {time:"06:49:07", auth:"袁超", sim: "5008", log: "Preflight is OK"},
                    {time:"07:32:43", auth:"宋继朋", sim: "5015", log: "Preflight is OK"},
                ],
            },
        },
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
        setwsid(state, newid) {
            state.ws.id = newid; 
        },
        getSomeNew(state, which) {
            state.getNew = true;
            switch (which) {
                case "log":
                    state.todayMessage.log.new = true;
                    state.todayMessage.log.type = "success";
                    break;
            };
        },
        refreshedLog(state, logs) {
            state.todayMessage.log.logs = logs;
            state.todayMessage.log.new = false;
            state.todayMessage.log.type = "info";
        },
    },
});

export default store;