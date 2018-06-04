//require
const path = require('path')
const {VueLoaderPlugin} = require('vue-loader')

module.exports = {
  mode: 'development',
  resolve: {
    // set alias
    alias: {
      'vue$': 'vue/dist/vue.common.js',
      'vue-chartjs$': 'vue-chartjs/dist/vue-chartjs.min.js'
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
      },
      {
        test: /\.less$/,
        exclude: /node_modules/,
        use: [
          'vue-style-loader',
          'css-loader',
          'less-loader'
        ]
      }
    ]
  },
  plugins: [
    new VueLoaderPlugin()
  ]
}

