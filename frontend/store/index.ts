import { defineStore } from 'pinia'
import Api from '~/utils/api';

import {isValidHttpUrl} from '~/utils/verify'
import {reactive} from "vue";

export const useMainStore = defineStore('main', {
    state: () => ({
        googleAuthUrl: '',
        searchResult:[],
        urlObj:{
            origin: '',
            isRandom: true,
            meta: {}
        },
        shortenUrl: ''
    }),
    actions: {
        async getGoogleAuthUrl() {
            const response = await Api.getGoogleAuthUrl()
            this.googleAuthUrl = response.data.value.url
        },
        async searchUrl(keyword:string){
            const response = await Api.searchUrl(keyword)
            if (response) {
                this.searchResult = response
            }
        },
        async postUrl(){
            if(isValidHttpUrl(this.urlObj.origin)){
                const response = await Api.postUrl(this.urlObj)
                if (response.data.value) {
                    this.shortenUrl = response.data.value.short
                }
            }
            else {
                alert('')
            }

        },
    },
})
