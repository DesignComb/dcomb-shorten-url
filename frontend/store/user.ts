import {defineStore} from 'pinia'
import Api from '~/utils/api';

export const useUserStore = defineStore('user', {
    state: () => ({
        userInfo: {
            user_id: 0,
            user_name: '',
            user_picture: ''
        },
    }),
    actions: {
        async getUserInfo() {
            const response = await Api.getUserInfo()
            this.userInfo = response.data
        },
    },
})//33