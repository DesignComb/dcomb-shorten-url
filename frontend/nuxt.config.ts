// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    modules: [
        '@nuxt/image-edge',
    ],
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
    }
})
