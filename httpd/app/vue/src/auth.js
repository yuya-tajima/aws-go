import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

import axios from 'axios';

export default {
	init(f) {
		let cred = null;
		axios.get('/cred', {
			headers: {
				'Content-Type': 'application/json',
			}
		})
		.then(function (response) {
			let data = response.data;
			cred = {
				'id': data.access_key,
				'secret_key': data.secret_key
			}
			f(cred);
		})
		.catch(function (error) {
			f(cred);
			console.log(error);
		})
		.finally(function () {
			//
		});
	}
};
