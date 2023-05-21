<script setup lang="ts">
import {ref, reactive, computed} from 'vue'
import {onMounted} from 'vue';

import 'boxicons/css/boxicons.min.css'
import {isValidHttpUrl} from '~/utils/verify'
import {Disclosure, DisclosureButton, DisclosurePanel, Popover, PopoverButton, PopoverPanel} from '@headlessui/vue'
import QrController from "~/components/qrCode/qrController.vue";
import LoadingAnimation from "~/components/common/loadingAnimation.vue";

import {getLinkPreview, getPreviewFromContent} from "link-preview-js";

const {appBaseUrl, apiBaseUrl} = useRuntimeConfig()

import {useMainStore} from '~/store'
import {useUserStore} from '~/store/user'
import ImgDialog from "~/components/common/imgDialog.vue";
import ImgCropper from "~/components/common/imgCropper.vue";

const main = useMainStore()
const user = useUserStore()

// import Api from '~/utils/api'

// const urlObj = reactive({
//     origin: '',
//     isRandom: true,
//     meta: {}
// })
const response: any = reactive({data: {short: ''}})
// const isSuccess = ref(false)

// async function submit() {
//     if (isValidHttpUrl(urlObj.origin)) {
//         response.data = await $fetch(`${apiBaseUrl}/api/urlShorten`, {
//             method: 'POST',
//             body: urlObj,
//         })
//         if (response.data.short !== '') {
//             isSuccess.value = true
//             // // 獲取原始網址的 meta 資料
//             // const metaResponse = await useFetch(urlObj.origin)
//             // const metaText = await metaResponse?.text()
//             // const parser = new DOMParser()
//             // const htmlDocument = parser.parseFromString(metaText, 'text/html')
//             // const metaTags = htmlDocument.getElementsByTagName('meta')
//             // // 從 meta 資料中獲取需要顯示的屬性
//             // const title = htmlDocument.querySelector('title')?.textContent
//             // const description = Array.from(metaTags).find(tag => tag.getAttribute('name') === 'description')?.getAttribute('content')
//             // const imageUrl = Array.from(metaTags).find(tag => tag.getAttribute('property') === 'og:image')?.getAttribute('content')
//             // urlObj.meta = {
//             //   title,
//             //   description,
//             //   imageUrl,
//             // }
//         }
//     } else {
//         alert('請填入正確的URL')
//     }
// }

const protocol = window.location.protocol;
const currentHost = window.location.host;
const completeUrl = computed(() => `${protocol}//${currentHost}/${user.shortenUrl}`)


const isCopied = ref(false)

const clickToCopy = () => {
    navigator.clipboard.writeText(completeUrl.value)
        .then(() => {
            isCopied.value = true
        })
}
main.getGoogleAuthUrl()
user.getUserInfo()

const handleUrlInputChange = () => {
    // main.searchUrl(urlObj.origin)
}
const handleUrlShorten = () => {
    if (user.userInfo.user_id >　0){
        user.userPostUrl()
    }
    else {
        user.postUrl()
    }
}
</script>

<template>

    <div class="relative mx-auto max-w-sm flex flex-wrap h-screen content-center items-center">
        <Popover v-slot="{ open }"
                 class="absolute w-full flex flex-wrap justify-end right-0 top-0">
            <div class="w-full flex justify-end">
                <div v-if="user.userInfo.user_id > 0" class="py-4 pl-2">
                    <div class="flex items-center w-full">
                        <img class="rounded-full w-6 h-6 mr-2" :src="user.userInfo.user_picture" alt="">
                        {{ user.userInfo.user_name }}
                    </div>
                </div>

                <PopoverButton v-else
                               class=" w-10 h-10 p-2 m-2 mr-0 cursor-pointer hover:bg-gray-600 rounded-full ">
                    <i class="bx bx-user"></i>
                </PopoverButton>
            </div>
            <transition
                    enter-active-class="transition duration-200 ease-out"
                    enter-from-class="-translate-y-1 opacity-0"
                    enter-to-class="translate-y-0 opacity-100"
                    leave-active-class="transition duration-150 ease-in"
                    leave-from-class="translate-y-0 opacity-100"
                    leave-to-class="-translate-y-1 opacity-0"
            >
                <PopoverPanel
                        class="relative right-0 z-10 w-1/2 mt-1.5 transform"
                >
                    <div
                            class="overflow-hidden rounded-lg shadow-lg ring-1 ring-black ring-opacity-5"
                    >
                        <a :href="main.googleAuthUrl" class="relative grid gap-8 bg-gray-600 p-1">
                            <img class="cursor-pointer"
                                 src="/btn_google_signin_dark_normal_web.png" alt="btn_google_dark_normal">
                        </a>
                    </div>
                </PopoverPanel>
            </transition>
        </Popover>
        <form class="w-full max-w-sm">
            <div class="flex items-center border-b border-teal-500 py-2">
                <input
                        v-model="user.urlObj.origin"
                        @change="handleUrlInputChange"
                        class="appearance-none bg-transparent border-none w-full text-white mr-3 py-1 px-2 leading-tight focus:outline-none"
                        type="text" placeholder="Url" aria-label="Full name">
                <button
                        @click="handleUrlShorten"
                        class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded"
                        type="button">
                    Shorten
                </button>
            </div>
        </form>
        <div v-if="user.shortenUrl" class="w-full flex flex-wrap relative mt-4">
            Congratulations! Your Short Url：<br>
            <input class="flex w-full mt-1.5 px-4 py-1.5 text-black rounded"
                   ref="shortenUrl" name="shortenUrl"
                   v-model="completeUrl" disabled/>
            <span @click="clickToCopy"
                  class="absolute w-8 h-8 z-10 -top-1 right-0 flex justify-center items-center cursor-pointer">
        <i v-if="!isCopied" class='bx bx-copy'></i>
        <i v-else class='bx bx-check'></i>
      </span>
        </div>
        <div class="w-full pt-8">
            <Disclosure v-slot="{ open }">
                <DisclosureButton :disabled="user.userInfo.user_id == 0"
                                  class="flex w-full justify-between item-center px-2 py-2 rounded-t-md border-b border-teal-500 text-left text-sm font-medium text-teal-500 hover:bg-gray-700 focus:outline-none focus-visible:ring focus-visible:ring-purple-500 focus-visible:ring-opacity-75"
                >
                    <span>Link preview Setting (registered only)</span>
                    <i :class="open ? 'rotate-180 transform' : ''"
                       class='bx bx-chevron-down'></i>
                </DisclosureButton>
                <DisclosurePanel class="flex flex-wrap pt-4 pb-2 text-gray-500 transform duration-500">
                    <label>Picture</label>
                    <div class="flex flex-wrap w-full my-1.5">
                        <img-dialog/>
                    </div>
                    <label>title</label>
                    <input class="flex w-full my-1.5 px-4 py-1.5 text-black rounded"
                           placeholder="preview Title"
                           v-model="user.urlObj.title"/>
                    <label>Description</label>
                    <textarea class="flex w-full my-1.5 px-4 py-1.5 text-black rounded"
                              rows="5" placeholder="preview Description"
                              v-model="user.urlObj.description"/>
                </DisclosurePanel>
            </Disclosure>
            <Disclosure as="div" class="mt-8"  v-slot="{ open }">
                <DisclosureButton :disabled="!user.shortenUrl"
                                  class="flex w-full justify-between item-center px-2 py-2 rounded-t-md border-b border-teal-500 text-left text-sm font-medium text-teal-500 hover:bg-gray-700 focus:outline-none focus-visible:ring focus-visible:ring-purple-500 focus-visible:ring-opacity-75"
                >
                    <span>QR code</span>
                    <i :class="open ? 'rotate-180 transform' : ''"
                       class='bx bx-chevron-down'></i>
                </DisclosureButton>
                <DisclosurePanel class="pt-4 pb-2 text-sm text-gray-500 transform duration-500">
                    <qr-controller :complete-url="completeUrl"/>
                </DisclosurePanel>
            </Disclosure>
        </div>
    </div>
<!--    <div class="flex w-full my-1.5">-->
<!--        <img-dialog/>-->
<!--    </div>-->
  <!-- 顯示原始網址和 meta 資料 -->
  <!--  <div v-if="urlObj.meta" class="mt-4">-->
  <!--    <div>{{ urlObj.origin }}</div>-->
  <!--    <div v-if="urlObj.meta.title">{{ urlObj.meta.title }}</div>-->
  <!--    <div v-if="urlObj.meta.description">{{ urlObj.meta.description }}</div>-->
  <!--    <img v-if="urlObj.meta.imageUrl" :src="urlObj.meta.imageUrl" alt="">-->
  <!--  </div>-->
</template>


<style scoped>
:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

input[name="shortenUrl"]:disabled {
    background-color: white;
    opacity: initial;
    cursor: initial;
}
</style>
