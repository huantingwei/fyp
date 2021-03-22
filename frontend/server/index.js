const express = require('express')
const createProxyMiddleware = require('http-proxy-middleware')

const options = {
    target:'http://18.162.60.222:8080',
    changeOrigin:true,
}
const proxy = createProxyMiddleware(options)

const app = express()

app.use(proxy)
app.listen(3001)
