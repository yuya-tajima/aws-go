import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
	state: {
		cred: {}
	},
	getters: {
		id: state => {
			return state.cred.id;
		},
		secret: state => {
			return state.cred.secret_key;
		}
	},
	mutations: {
		setCred: (state, cred) => {
			state.cred = cred;
		}
	}
});
