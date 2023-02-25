<script setup lang="ts">
import {ref, reactive, computed} from 'vue'
import 'boxicons/css/boxicons.min.css'
import {isValidHttpUrl} from '~/utils/verify'

const urlObj = reactive({
  origin:'',
  isRandom:true
})
const response: any = reactive({data: {}})
const isSuccess = ref(false)

async function submit() {
  if(isValidHttpUrl(urlObj.origin)){
    response.data = await $fetch('http://54.249.0.5:8000/urlShorten', {
      method: 'POST',
      body: urlObj,
    })
    if(response.data.short !== ''){
      isSuccess.value = true
    }
  }
  else {
    alert('請填入正確的URL')
  }
}

const completeUrl = computed(() => {
  return 'http://localhost:3000/' + response.data.short
})
const isCopied = ref(false)
const clickToCopy = () => {
  // alert('被按了 ><')
  navigator.clipboard.writeText('http://localhost:3000/' + response.data.short)
      .then(()=>{
        isCopied.value = true
      })
}
const testImg = 'https://images.pexels.com/photos/226746/pexels-photo-226746.jpeg'

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
  </div>

</template>


<style scoped>
:disabled{
  background-color: white;
}
</style>