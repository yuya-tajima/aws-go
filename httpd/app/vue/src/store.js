import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
	state: {
		cred: null
	},
	getters: {
	},
	mutations: {
		setCred: (state, val) => {
			state.cred = val;
		}
	}
});
