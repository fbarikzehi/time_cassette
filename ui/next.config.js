/** @type {import('next').NextConfig} */

const isProduction = process.env.NODE_ENV === 'production'

const nextConfig = {
    distDir: 'build',
    async redirects() {
        return [
            {
                source: '/',
                destination: '/auth/signin',
                permanent: true,
            },
        ]
    },
    assetPrefix: isProduction ? 'https://timecassette.com/' : undefined
}

module.exports = nextConfig
