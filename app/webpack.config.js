//require
const path = require('path')
const {VueLoaderPlugin} = require('vue-loader')

module.exports = {
  mode: 'development',
  resolve: {
    alias: {
      'vue$': 'vue/dist/vue.common.js' // set alias
    }
  },
  entry: path.resolve(__dirname, 'assets/js/index.js'),
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'app.js'
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        exclude: /node_modules/,
        use: [
          'vue-loader'
        ]
      },
      {
        test: /\.css$/,
        exclude: /node_modules/,
        use: [
          'vue-style-loader',
          'css-loader'
        ]
      }
    ]
  },
  plugins: [
    new VueLoaderPlugin()
  ]
}

