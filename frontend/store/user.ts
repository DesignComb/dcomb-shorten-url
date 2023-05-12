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
            if (response.data.value) {
                this.userInfo = response.data.value.data
            }
        },
        async login(code: string) {
            await Api.login(code)
        },
    },
})