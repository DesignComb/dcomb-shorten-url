<script setup lang="ts">
import {ref, reactive, computed} from 'vue'
import 'boxicons/css/boxicons.min.css'
import {isValidHttpUrl} from '~/utils/verify'
import { Disclosure, DisclosureButton, DisclosurePanel } from '@headlessui/vue'
import QrController from "~/components/qrCode/qrController.vue";
import LoadingAnimation from "~/components/common/loadingAnimation.vue";

import { getLinkPreview, getPreviewFromContent } from "link-preview-js";

const { appBaseUrl,apiBaseUrl } = useRuntimeConfig().public
import { Response } from 'node-fetch';


const urlObj = reactive({
  origin:'',
  isRandom:true,
  meta:{}
})
const response: any = reactive({data: { short: '' }})
const isSuccess = ref(false)

async function submit() {
  if(isValidHttpUrl(urlObj.origin)){
    response.data = await $fetch(`${apiBaseUrl}/urlShorten`, {
      method: 'POST',
      body: urlObj,
    })
    if(response.data.short !== ''){
      isSuccess.value = true
      // 獲取原始網址的 meta 資料
      const metaResponse = await useFetch(urlObj.origin)
      const metaText = await metaResponse?.text()
      const parser = new DOMParser()
      const htmlDocument = parser.parseFromString(metaText, 'text/html')
      const metaTags = htmlDocument.getElementsByTagName('meta')
      // 從 meta 資料中獲取需要顯示的屬性
      const title = htmlDocument.querySelector('title')?.textContent
      const description = Array.from(metaTags).find(tag => tag.getAttribute('name') === 'description')?.getAttribute('content')
      const imageUrl = Array.from(metaTags).find(tag => tag.getAttribute('property') === 'og:image')?.getAttribute('content')
      urlObj.meta = {
        title,
        description,
        imageUrl,
      }
    }
  }
  else {
    alert('請填入正確的URL')
  }
}

const completeUrl = computed(() => `${appBaseUrl}/${response.data.short}`)

const isCopied = ref(false)

const clickToCopy = () => {
  navigator.clipboard.writeText(completeUrl.value)
      .then(()=>{
        isCopied.value = true
      })
}
</script>

<template>
  <div class="mx-auto max-w-sm flex flex-wrap h-screen content-center items-center">
    <form class="w-full max-w-sm">
      <div class="flex items-center border-b border-teal-500 py-2">
        <input
            v-model="urlObj.origin"
            class="appearance-none bg-transparent border-none w-full text-white mr-3 py-1 px-2 leading-tight focus:outline-none"
            type="text" placeholder="Url" aria-label="Full name">
        <button
            @click="submit"
            class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded"
            type="button">
          Shorten
        </button>
      </div>
    </form>
    <div v-if="isSuccess" class="w-full flex flex-wrap relative mt-4">
      Congratulations! Your Short Url：<br>
      <input class="flex w-full mt-1.5 px-4 py-1.5 text-black rounded"
             ref="shortUrl"
             v-model="completeUrl" disabled/>
      <span @click="clickToCopy" class="absolute w-8 h-8 z-10 -top-1 right-0 flex justify-center items-center cursor-pointer">
        <i v-if="!isCopied" class='bx bx-copy'></i>
        <i v-else class='bx bx-check'></i>
      </span>
    </div>
    <div class="w-full pt-8">
      <Disclosure v-slot="{ open }">
        <DisclosureButton
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
      <Disclosure as="div" class="mt-8" v-slot="{ open }">
        <DisclosureButton
            class="flex w-full justify-between item-center px-2 py-2 rounded-t-md border-b border-teal-500 text-left text-sm font-medium text-teal-500 hover:bg-gray-700 focus:outline-none focus-visible:ring focus-visible:ring-purple-500 focus-visible:ring-opacity-75"
        >
          <span>SEO Renderer</span>
          <i :class="open ? 'rotate-180 transform' : ''"
             class='bx bx-chevron-down'></i>
        </DisclosureButton>
        <DisclosurePanel class="px-4 pt-4 pb-2 text-sm text-gray-500 transform duration-500">
          If you're unhappy with your purchase for any reason, email us within
          90 days and we'll refund you in full, no questions asked.
        </DisclosurePanel>
      </Disclosure>
    </div>
  </div>
  <!-- 顯示原始網址和 meta 資料 -->
  <div v-if="urlObj.meta" class="mt-4">
    <div>{{ urlObj.origin }}</div>
    <div v-if="urlObj.meta.title">{{ urlObj.meta.title }}</div>
    <div v-if="urlObj.meta.description">{{ urlObj.meta.description }}</div>
    <img v-if="urlObj.meta.imageUrl" :src="urlObj.meta.imageUrl" alt="">
  </div>
</template>


<style scoped>
:disabled{
  background-color: white;
}
</style>