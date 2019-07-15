<template>
  <v-container fluid>
    <h1>This is dash board</h1>
  </v-container>
</template>

<script>
export default {
	name: 'index',
	methods: {
		reload() {
			this.fetch()
		},
		fetch () {
			this.$router.go({path: this.$router.currentRoute.path, force: true});
		}
	}
}

</script>
