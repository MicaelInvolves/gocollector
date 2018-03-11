/* global __dirname, require, module*/

const webpack = require('webpack');
const UglifyJsPlugin = webpack.optimize.UglifyJsPlugin;
const path = require('path');
const env = require('yargs').argv.env; // use --env with webpack 2
const pkg = require('./package.json');
const CopyWebpackPlugin = require('copy-webpack-plugin');

let libraryName = pkg.name;

let plugins = [],
  outputFile;

plugins.push(new webpack.DefinePlugin({
  'API_PATH': JSON.stringify('/api/access')
}));

if (env === 'build') {
  plugins.push(new UglifyJsPlugin({
    minimize: true
  }));
  plugins.push(new CopyWebpackPlugin([
    { from: 'lib/index.html', to: '../../static/examples/index.html', force: true},
    { from: 'lib/contact.html', to: '../../static/examples/contact.html', force: true},
    { from: 'lib/about.html', to: '../../static/examples/about.html', force: true},
    { from: 'lib/price.html', to: '../../static/examples/price.html', force: true},
    { from: 'lib/rd-collector.min.js', to: '../../static/examples/rd-collector.min.js', force: true}
  ]));
  outputFile = libraryName + '.min.js';
} else {
  outputFile = libraryName + '.js';
}

const config = {
  entry: __dirname + '/src/index.js',
  devtool: 'source-map',
  output: {
    path: __dirname + '/lib',
    filename: outputFile,
    library: libraryName,
    libraryTarget: 'umd',
    umdNamedDefine: true
  },
  module: {
    rules: [{
      test: /(\.jsx|\.js)$/,
      loader: 'babel-loader',
      exclude: /(node_modules|bower_components)/
    },
    {
      test: /(\.jsx|\.js)$/,
      loader: 'eslint-loader',
      exclude: /node_modules/
    }
    ]
  },
  resolve: {
    modules: [path.resolve('./node_modules'), path.resolve('./src')],
    extensions: ['.json', '.js']
  },
  plugins: plugins
};

module.exports = config;
