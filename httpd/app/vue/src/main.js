import Vue from 'vue'

import VueRouter from 'vue-router'
import Vuetify from 'vuetify';

import 'vuetify/dist/vuetify.min.css'

Vue.use(VueRouter);
Vue.use(Vuetify);

import auth from '@/auth.js'
import store from '@/store.js'
import index from '@/components/Index.vue'
import ec2 from '@/components/Ec2.vue'
import App from '@/App.vue'

auth.init( function (cred) {

	store.commit('setCred', cred);

	const router = new VueRouter({
		routes: [
			{
				path:'/'
				,component: index
			},
			{
				path:'/ec2',
				component: ec2
			}
		]
	});

	let vm = new Vue({
		el: '#app',
		render: h => h(App),
		router:router,
		store,
		mounted () {
			console.log(this.getCred);
		}
	})
});

