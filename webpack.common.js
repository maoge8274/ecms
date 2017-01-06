var webpack = require('webpack');


module.exports = {
    entry: {
        'polyfills': './static/js/polyfills.ts',
        // 'vendor': './static/js/vendor.ts',
        'admin': './static/js/admin/main.ts'
    },

    resolve: {
        extensions: ['', '.ts', '.js']
    },
    module : {
        loaders: [
            {
                test   : /\.ts$/,
                loaders: ['awesome-typescript-loader', 'angular2-template-loader']
            },
            {
                test: /\.html$/,
                loader: 'html'
            },
            {
                test: /\.scss$/,
                exclude: /node_modules/,
                loaders: ['raw-loader', 'sass-loader']
            }
        ]
    }
}