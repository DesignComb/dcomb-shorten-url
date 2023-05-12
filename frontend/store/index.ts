import { defineStore } from 'pinia'
import Api from '~/utils/api';

export const useMainStore = defineStore('main', {
    state: () => ({
        googleAuthUrl: '',
    }),
    actions: {
        async getGoogleAuthUrl() {
            const response = await Api.getGoogleAuthUrl()
            this.googleAuthUrl = response.url
        },
    },
})//33