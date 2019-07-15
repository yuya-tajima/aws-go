<template>
	<v-content>
		<v-container
			fluid
			fill-height
		>
			<v-layout
				align-center
				justify-center
			>
				<v-flex sm8 xs12 md4>
					<v-card>
						<v-toolbar
							color="primary"
							dark
						>
							<v-toolbar-title>
								Login Form
							</v-toolbar-title>
						</v-toolbar>
						<v-form @submit.prevent="exec">
							<v-card-text>
									<v-text-field
										v-model="login"
										prepend-icon="person"
										:rules="[() => login.length > 0 || 'Login ID is required']"
										label="Login ID"
										class="mt-1"
										required
									></v-text-field>
									<v-text-field
										v-model="password"
										prepend-icon="lock"
										:rules="[() => login.length > 0 || 'Login Password is required']"
										:append-icon="vf ? 'visibility' : 'visibility_off'"
										@click:append="vf = !vf"
										:type="vf ? 'password' : 'text'"
										label="Enter your password"
										required
									></v-text-field>
							</v-card-text>
							<v-card-actions>
								<v-spacer></v-spacer>
								<v-btn type="submit" depressed color="primary">Submit
									<v-icon dark right>check_circle</v-icon>	
								</v-btn>
							</v-card-actions>
						</v-form>
					</v-card>
				</v-flex>
			</v-layout>
		</v-container>
	</v-content>
</template>

<script>
import api from '@/utils/api.js'
export default {
	name: 'login',
	data () {
		return {
			login:'',
			password:'',
			vf: false
		}
	},
	methods: {
		exec: function () {
			api.post( '/auth', {
				"login": this.login,
				"password": this.password
				})
			.then(res => {
				console.log(res)
				this.$store.commit('setCred', true)
			})
			.catch(err => {
			
			});
		}
	}
}
</script>
