
<template>
	<v-app>
		<v-navigation-drawer
			fixed
			clipped app
			v-model="navBar"
			width=150
		>
			<v-list dense class="pt-2">
				<router-link to="/">
					<v-list-tile>
						<v-list-tile-action>
							<v-icon>dashboard</v-icon>
						</v-list-tile-action>
						<v-list-tile-content>
							<v-list-tile-title>Index</v-list-tile-title>
						</v-list-tile-content>
					</v-list-tile>
				</router-link>
				<router-link to="/ec2">
					<v-list-tile>
						<v-list-tile-action>
							<v-icon>computer</v-icon>
						</v-list-tile-action>
						<v-list-tile-content>
							<v-list-tile-title>EC2</v-list-tile-title>
						</v-list-tile-content>
					</v-list-tile>
				</router-link>
			</v-list>
		</v-navigation-drawer>
		<v-toolbar
			dark
			color="primary"
			clipped-left
			fixed
			app
		>
			<v-toolbar-side-icon v-on:click.stop="navBar =! navBar">
			</v-toolbar-side-icon>
			<v-toolbar-title class="white--text">
				Photoruction Development Admin
			</v-toolbar-title>
			<v-spacer></v-spacer>
			<v-btn icon>
				<v-icon>search</v-icon>
			</v-btn>
			<v-btn icon>
				<v-icon>apps</v-icon>
			</v-btn>
			<v-btn icon v-on:click="reload">
				<v-icon>refresh</v-icon>
			</v-btn>
			<v-btn icon>
				<v-icon>more_vert</v-icon>
			</v-btn>
		</v-toolbar>
		<v-content>
			<v-fade-transition mode="out-in">
				<router-view></router-view>
			</v-fade-transition>
		</v-content>
	</v-app>
</template>

<script>
import axios from 'axios';
export default {
	name: 'app',
	data () {
		return {
			navBar:null
		}
	},
	created() {
		let self = this;
		axios.get('http://localhost:8080/cred', {
			headers: {
				'Content-Type': 'application/json',
			}
		})
		.then(function (response) {
			let data = response.data;
			console.log('created');
		})
		.catch(function (error) {
			console.log(response);
		})
		.finally(function () {
			//
		});
	},
	mounted() {
		console.log('app load');
	},
	methods: {
		reload: function (event) {
			this.$router.go({path: this.$router.currentRoute.path, force: true});
		}
	}
}
</script>

<style>
	a {
		text-decoration: none;
	}
</style>

<style scoped>
	.v-list__tile__action {
		min-width:30px;
	}
</style>
