// did not work

const createProxyMiddleware = require('http-proxy-middleware');

module.exports = function(app) {
    app.use(createProxyMiddleware({
            target: 'http://18.162.60.222:8080',
            changeOrigin: true,
        })
    );
};