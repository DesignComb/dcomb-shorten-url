// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    modules: [
        '@nuxt/image-edge',
        '@pinia/nuxt',
    ],
    runtimeConfig: { // auto get .env setting
        public: {
            appBaseUrl: '',
            apiBaseUrl: '',
            serverApiBaseUrl: ''
        }
    },
    devServerHandlers: [],
    image: {
        imgix: {
            baseURL: 'https://assets.imgix.net'
        }
    },
    css: ['~/assets/css/main.css', 'boxicons/css/boxicons.min.css'],
    postcss: {
        plugins: {
            tailwindcss: {},
            autoprefixer: {},
        },
    },
    build: {
        transpile: ['@headlessui/vue']
    },
    routeRules: {
        // // Static page generated on-demand, revalidates in background
        // '/blog/**': { swr: true },
        // // Static page generated on-demand once
        // '/articles/**': { static: true },
        // // Set custom headers matching paths
        // '/_nuxt/**': { headers: { 'cache-control': 's-maxage=0' } },
        // // Render these routes with SPA
        // '/admin/**': { ssr: false },
        // // Add cors headers
        // '/api/v1/**': { cors: true },
        // // Add redirect headers
        // '/old-page': { redirect: '/new-page' },

        // set ssr false on enter url page
        '/': { ssr: false }
    }
})
