import {defineStore} from 'pinia'
import Api from '~/utils/api';
import {dataURItoBlob} from  '~/utils/dataURI2Blob'

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
        },
        imageObj:{
            uploadImgData:'',
            croppedImg:'',
            uploadedUrl:''
        },
    }),
    actions: {
        async getUserInfo() {
            const response = await Api.getUserInfo()
            if (response) {
                this.userInfo = response
            }
        },
        async login(code: string) {
            await Api.login(code)
        },
        async userSearchUrl(keyword:string){
            const response = await Api.userSearchUrl(this.userInfo.user_id,keyword)
            if (response) {
                this.searchResult = response
            }
        },
        async userPostImage(){
            const formData = new FormData()
            formData.append('file',dataURItoBlob(this.imageObj.croppedImg))
            const response = await Api.userPostImage(this.userInfo.user_id,formData)
            if (response) {
                this.imageObj.uploadedUrl = response.url
            }
        }
    },
})