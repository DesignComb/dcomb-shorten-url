<script setup lang="ts">
import {isValidHttpUrl} from '~/utils/verify'
import LoadingAnimation from "~/components/common/loadingAnimation.vue";

const route = useRoute()
const {id} = route.params
const { appBaseUrl,apiBaseUrl } = useRuntimeConfig().public

const {data, pending, error, refresh} = await useAsyncData(
    'id',
    () => $fetch(`http://172.17.0.1:8000/api/r/${id}`)
)

onMounted(() => {
  redirect()
})

const handleRefresh = async () => {
  console.log(data)
  await refresh()
  await redirect()
}

const redirect = () => {
  const value = (data as { value?: { origin: string } }).value
  console.log(value)
  if (value && isValidHttpUrl(value.origin)) {
    const url = new URL(value.origin)
    window.location.href = url.href
  }
}


</script>
<template>
  <Head>
    <Title>{{ data?.origin }}</Title>
    <Meta name="description" :content="data?.origin"/>
  </Head>
  <div class="flex justify-center items-center h-screen">
    <div class="text-center mb-4">
      <loading-animation/>
      <div v-if="error"> 錯誤: {{ error }}<br>
        <button
            class="rounded-sm bg-blue-500 py-3 px-8 text-xl font-medium text-white hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:ring-offset-2"
            @click="handleRefresh"
        >
          點此以重新導向
        </button>
      </div>
    </div>
  </div>
</template>
<style scoped>
</style>