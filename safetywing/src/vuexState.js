import Vuex from "vuex";
import Vue from "vue";


Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        logined: false,
        user: {id: "0"},
        ws: {id: ''},
        todayMessage: {
            log: {
                new: false,
                type: "info",
                logs: [],
            },
            malf: {
                new: false,
                type: "info",
                malfs: [],
            },
            trainingList: {
                new: false,
                type: "info",
                list: [],
            }
        },
    },
    getters: {
        getNew: state => {
            return state.todayMessage.log.new || state.todayMessage.malf.new || state.todayMessage.trainingList.new;
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
                case "malf":
                    state.todayMessage.malf.new = true;
                    state.todayMessage.malf.type = "success";
                    break;
            }
        },
        refreshedLog(state, newlogs) {
            state.todayMessage.log.logs = newlogs;
            state.todayMessage.log.new = false;
            state.todayMessage.log.type = "info";
        },
        refreshedMalf(state, newMalfs) {
            state.todayMessage.malf.malfs = newMalfs;
            state.todayMessage.malf.new = false;
            state.todayMessage.malf.type = "info";
        },
        refreshedTrainingList(state, newtlist) {
            state.todayMessage.trainingList.list = newtlist;
            state.todayMessage.trainingList.new = false;
            state.todayMessage.trainingList.type = "info";
        },
    },
});

export default store;