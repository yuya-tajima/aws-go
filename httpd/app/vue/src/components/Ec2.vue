<template>
	<v-container fluid>
		<v-data-table
		:headers="headers"
		:items="items"
		hide-actions
		class="ec2-desc-table"
		>
			<template v-slot:items="info">
				<td class="text-xs-left">{{ info.item.name }}</td>
				<td class="text-xs-left">{{ info.item.ins_id }}</td>
				<td class="text-xs-left">{{ info.item.image_id }}</td>
				<td class="text-xs-left">{{ info.item.public_ipv4 }}</td>
				<td class="text-xs-left">{{ info.item.private_ipv4 }}</td>
				<td class="text-xs-left">{{ info.item.ins_type }}</td>
				<td class="text-xs-left">
				<v-menu bottom offset-y>
					<template v-slot:activator="{ on }">
						<a v-bind:class="actionClass(info.item.status_name)" v-on="on">{{ info.item.status_name }}</a>
					</template>
					<v-list dense>
						<v-list-tile
							v-for="(item, i) in actions"
							:key="i"
							v-on:click=""
						>
							<v-icon small left>
								{{ item.icon }}
							</v-icon>
							<v-list-tile-title>
								<a v-on:click.stop="action(item.action, info.item.ins_id)">
									{{ item.name }}
								</a>
							</v-list-tile-title>
						</v-list-tile> </v-list>
				</v-menu> 

				</td> </template>
		</v-data-table>
	</v-container>
</template>

<script>
import api from '@/utils/api.js'
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
					name: 'start',
					action: 'start',
					icon: 'power'
				},
				{
					name: 'stop',
					action: 'stop',
					icon: 'power_off'
				},
			]
		}
	},
	methods: {
		action(action, ins_id) {
			api.post(`/ec2/${action}`,
				{
					ins_id: ins_id
				},
				{
					headers: {
						'Content-Type': 'application/json',
					}
				}
			)
			.then( response => {
				console.log(response);
			})
			.catch( error => {
				console.log(error);
			})
			.finally( () => {
				this.fetch()
			});
		},
		actionClass(status) {
			return {
				'red--text' : (status === 'stopped'),
				'amber--text' : (status === 'stopping'),
				'orange--text' : (status === 'pending'),
				'green--text' : (status === 'running')
			}
		},
		reload() {
			this.fetch()
		},
		fetch () {
			api.get('/ec2/desc', {
				headers: {
					'Content-Type': 'application/json',
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
	},
	mounted: function () {
		this.fetch()
	}
}
</script>
