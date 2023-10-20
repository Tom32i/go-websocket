const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = (env, argv) => ({
  target: 'web',
  entry: './client/index.js',
  output: {
    filename: 'client.js',
    path: `${__dirname}/public`,
  },
  devtool: argv.mode === 'production' ? false : 'source-map',
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: { presets: ['@babel/preset-env'] }
        }
      },
      {
        test: /\.scss$/,
        use: [
          'style-loader',
          'css-loader',
          'sass-loader'
        ],
      },
    ]
  },
  resolve: {
    alias: {
      '@client': `${__dirname}/client`,
    }
  },
  plugins: [
    new HtmlWebpackPlugin({ template: './client/index.html' })
  ]
});
