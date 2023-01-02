const { defineConfig } = require('@vue/cli-service')
// module.exports = defineConfig({
//   transpileDependencies: true
// })

module.exports={
  devServer:{
  proxy:{
  ["/dev-api"]:{
    target:'http://localhost:8082/getinfo',
    changeOrigin:true,
    pathRewrite: {
      ['^' + "/dev-ap"]: ''
    }
  }
  }
}
}