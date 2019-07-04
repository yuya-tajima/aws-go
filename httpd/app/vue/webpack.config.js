
const path = require('path');

const { VueLoaderPlugin } = require('vue-loader');

module.exports = {
	entry: './src/main.js',
	output: {
		path: path.resolve(__dirname, '../assets/js'),
		filename: 'bundle.js'
	},
	module: {
		rules: [
			{
				test: /\.vue$/,
				use: [
					{ loader: 'vue-loader' }
				]
			},
			{
				test: /\.css$/,
				use: [
					{ loader: "style-loader" },
					{ loader: "css-loader" }
				]
			},
			{
				test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
				use: [
					{ loader: 'url-loader?mimetype=image/svg+xml' }
				],
			}
		]
},
  plugins: [
	new VueLoaderPlugin()
  ],
  resolve: {
	alias: {
	  'vue$': 'vue/dist/vue.esm.js',
		'@': path.resolve(__dirname, 'src')
	}
  }
}


