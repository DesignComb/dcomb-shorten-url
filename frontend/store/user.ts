import {defineStore} from 'pinia'
import Api from '~/utils/api';

export const useUserStore = defineStore('user', {
    state: () => ({
        userInfo: {
            user_id: 0,
            user_name: '',
            user_picture: ''
        },
        searchResult:{},
        urlPreview:{
            title:'',
            description:'',
            picture:''
        }
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
        async userSearchUrl(userID:number,keyword:string){
            const response = await Api.userSearchUrl(userID,keyword)
            if (response.data.value) {
                this.searchResult = response.data.value.data
            }
        }
    },
})