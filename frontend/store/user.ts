import {defineStore} from 'pinia'
import Api from '~/utils/api';
import {dataURItoBlob} from  '~/utils/dataURI2Blob'
import {isValidHttpUrl} from "~/utils/verify";

export const useUserStore = defineStore('user', {
    state: () => ({
        userInfo: {
            user_id: 0,
            user_name: '',
            user_picture: ''
        },
        searchResult:{},
        urlObj: {
            origin: '',
            isRandom: true,
            imageId: 0,
            title: '',
            description: ''
        },
        shortenUrl: '',
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
        async postUrl() {
            if (isValidHttpUrl(this.urlObj.origin)) {
                const response = await Api.postUrl(this.urlObj)
                if (response) {
                    this.shortenUrl = response.short
                }
            } else {
                alert('請輸入正確的網址格式')
            }
        },
        async userPostUrl() {
            if (isValidHttpUrl(this.urlObj.origin)) {
                const response = await Api.userPostUrl(this.userInfo.user_id,this.urlObj)
                if (response) {
                    this.shortenUrl = response.short
                }
            } else {
                alert('請輸入正確的網址格式')
            }
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
                this.urlObj.imageId = response.id
            }
        }
    },
})