import {defineStore} from 'pinia'
import Api from '~/utils/api';

import {isValidHttpUrl} from '~/utils/verify'
import {reactive} from "vue";

export const useMainStore = defineStore('main', {
    state: () => ({
        googleAuthUrl: '',
        searchResult: []
    }),
    actions: {
        async getGoogleAuthUrl() {
            const response = await Api.getGoogleAuthUrl()
            this.googleAuthUrl = response.url
        },
        async searchUrl(keyword: string) {
            const response = await Api.searchUrl(keyword)
            if (response) {
                this.searchResult = response
            }
        },

    },
})
