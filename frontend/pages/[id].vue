<script setup lang="ts">
import {isValidHttpUrl} from '~/utils/verify'

const route = useRoute()
const {id} = route.params

const { data, pending, error, refresh } = await useAsyncData(
    'id',
    () => $fetch(`http://54.249.0.5:8000/r/${id}`)
)

const redirect = onMounted(() =>{
  console.log('aaa')
  // console.log(data?.value?.origin)
  // if(isValidHttpUrl(data?.value?.origin)){
  //   window.location = data?.value?.origin
  // }
})
</script>

<template>
  <Head>
    <Title>{{ id }}</Title>
    <Meta name="description" :content="data?.origin"/>
  </Head>
  <h1>這裡是redirect 頁面 id1111 = {{ id }}</h1>
  Error:{{error}}
  <p class="text-2xl text-gray-600">
    請求狀態:
    {{ pending ? '請求中' : '完成' }}
  </p>
  <span class="mt-4 text-2xl text-gray-600">回傳資料:</span>
  <p class="mt-4 text-3xl font-semibold text-blue-500">{{ data?.origin }}</p>
  <button
      class="mt-6 rounded-sm bg-blue-500 py-3 px-8 text-xl font-medium text-white hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:ring-offset-2"
      @click="refresh"
  >
    重新獲取資料
  </button>
</template>

<style scoped>
</style>