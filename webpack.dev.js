var webpackMerge = require('webpack-merge');
var commonConfig = require('./webpack.common');
var path = require('path');

module.exports = webpackMerge(commonConfig, {
    devtool: 'cheap-module-eval-source-map',

    output   : {
        path         : path.resolve('./static/dist'),
        publicPath: "/static/dist/",
        filename     : '[name].js',
        chunkFilename: '[id].chunk.js'
    },
    devServer: {
        hot: true,
        port:8181
    }
});