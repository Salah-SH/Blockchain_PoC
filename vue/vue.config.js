module.exports = {
    devServer: {
        proxy: {
            '^/api': {
                target: 'http://localhost:1317',
                changeOrigin: true,
                logLevel: 'debug',
                pathRewrite: { '^/api': '' },
            },
            '^/backend': {
                target: 'http://localhost:8000',
                changeOrigin: true,
                logLevel: 'debug',
                pathRewrite: { '^/backend': '/' },
            },
            
        },
    },
}