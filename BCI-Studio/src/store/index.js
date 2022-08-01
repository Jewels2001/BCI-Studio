import {createStore} from "vuex";
 
// Vue.use(Vuex);

export default createStore({
    state: {
        data: {
            ch0: {
                // labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                data: [0, 0, 0, 0, 0],
                // hist: [[]]
            },
            ch1: {
                // labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                data: [0, 0, 0, 0, 0],
                // hist: [[]]
            },
            ch2: {
                // labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                data: [0, 0, 0, 0, 0],
                // hist: [[]]
            },
            ch3: {
                // labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                data: [0, 0, 0, 0, 0],
                // hist: [[]]
            }
        }
    },
    getters: {
        getData(state) {
            return state.data
        }
    },
    mutations: {
        updateData(state, payload) {
            state.data = payload
        }
    },
})