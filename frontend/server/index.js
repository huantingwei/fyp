// const express = require('express')
// const bodyParser = require('body-parser')
// const pino = require('express-pino-logger')()

// const app = express()
// app.use(bodyParser.urlencoded({ extended: false }))
// app.use(pino)

// app.get('/api/greeting', (req, res) => {
//     const name = req.query.name || 'World'
//     res.setHeader('Content-Type', 'application/json')
//     res.send(JSON.stringify({ greeting: `Hello ${name}!` }))
// })

const express = require('express')
const { createProxyMiddleware } = require('http-proxy-middleware')

const app = express()

app.use(
    createProxyMiddleware({
        target: 'http://18.162.60.222:8080/',
        changeOrigin: true,
    })
)
app.listen(3001)
