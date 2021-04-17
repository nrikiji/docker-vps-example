const path = require("path");
const outputPath = path.resolve(__dirname, "dist");

module.exports = {
  mode: "production",
  entry: "./src/main.ts",
  output: {
    filename: "main.js",
    path: outputPath,
  },
  devServer: {
    contentBase: outputPath,
    port: 3000,
  },
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: "ts-loader",
      },
    ],
  },
  resolve: {
    extensions: [".ts", ".js"],
  },
};
