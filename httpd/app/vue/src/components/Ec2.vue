<template>
	<v-container fluid>
		<v-data-table
		:headers="headers"
		:items="items"
		hide-actions
		class="ec2-desc-table"
		>
			<template v-slot:items="info">
				<td class="text-xs-left">{{ info.item.name }}</td> <td class="text-xs-left">{{ info.item.ins_id }}</td>
				<td class="text-xs-left">{{ info.item.image_id }}</td>
				<td class="text-xs-left">{{ info.item.public_ipv4 }}</td>
				<td class="text-xs-left">{{ info.item.private_ipv4 }}</td>
				<td class="text-xs-left">{{ info.item.ins_type }}</td>
				<td class="text-xs-left">
				<v-menu bottom offset-y>
					<template v-slot:activator="{ on }">
						<a v-bind:class="[ info.item.status_name === 'stopped' ? 'red--text' : 'teal--text'  ]" v-on="on">{{ info.item.status_name }}</a>
					</template>
					<v-list dense>
						<v-list-tile v-for="(item, i) in actions" :key="i" v-on:click="">
							<v-icon small left>{{ item.icon }}</v-icon>
							<v-list-tile-title>
								{{ item.title }}
							</v-list-tile-title>
						</v-list-tile>
					</v-list>
				</v-menu> 

				</td>
			</template>
		</v-data-table>
	</v-container>
</template>

<script>
import axios from 'axios';
export default {
	name: 'ec2',
	data () {
		return {
			headers: [
				{
					text: 'Name',
					value: 'name',
					align: 'left',
					sortable: false
				},
				{
					text: 'Instance ID',
					value: 'ins_id',
					align: 'left',
					sortable: false
				},
				{
					text: 'Image ID',
					value: 'image_id',
					align: 'left',
					sortable: false
				},
				{
					text: 'Public IPv4',
					value: 'public_ipv4',
					align: 'left',
					sortable: false
				},
				{
					text: 'Private IpV4',
					value: 'private_ipv4',
					align: 'left',
					sortable: false
				},
				{
					text: 'Instance Type',
					value: 'ins_type',
					align: 'left',
					sortable: false
				},
				{
					text: 'State',
					value: 'ins_type',
					align: 'left',
					sortable: false
				}
        	],
			items: [],
			actions:[
				{
					title: 'start',
					icon: 'power'
				},
				{
					title: 'stop',
					icon: 'power_off'
				},
			]
		}
	},
	mounted: function () {
		axios.get('//localhost:8081/ec2/desc', {
			headers: {
				'Content-Type': 'application/json',
				'Secret_Id': this.$store.getters.id,
				'Secret_key': this.$store.getters.secret,
				'Region': 'ap-northeast-1',
			}
		})
		.then( response => {
			const NAME_TAG = 'Name';
			this.items = [];
			let res = response.data.items;
			let items = [];
			res.forEach( item => {
				item.name = '';
				if (!item.public_ipv4) {
					item.public_ipv4 = 'IP address is not attached';
				}
				if (item.tags.length !== 0) {
					item.tags.forEach(tag =>{
						if (tag.key === NAME_TAG) {
							item.name = tag.value;
						}
					});
				}
				this.items.push(item);
			});
		})
		.catch( error => {
			console.log(error);
		})
		.finally( () => {
			//
		});
	}
}
</script>
