module.exports = {
  entry: './entry.js',
  output: {
    path: __dirname,
    filename: 'bundle.js'
  },
  modules: {
    rules: [{
      test: /test.js$/,
      use: 'mocha-loader',
      exclude: /node_modules/,
    }]
  }
}
