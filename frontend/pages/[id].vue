<script setup lang="ts">
import {isValidHttpUrl} from '~/utils/verify'
import LoadingAnimation from "~/components/common/loadingAnimation.vue";
import Api from "~/utils/api";

const route = useRoute()
const {id} = route.params
const {appBaseUrl, apiBaseUrl, serverApiBaseUrl} = useRuntimeConfig()

interface resData {
    code: string;
    data: object;
}

// useAsyncData 不能被封裝到其他地方
const {data, pending, error, refresh} = await useAsyncData('id',
    () => $fetch(`${serverApiBaseUrl}/api/r/${id}`))
const resData = data.value as resData

onMounted(async () => {
    await redirect(resData.data.origin)
})
const redirect = (redirectUrl: string) => {
    if (isValidHttpUrl(redirectUrl)) {
        const url = new URL(redirectUrl)
        window.location.href = url.href
    }
}
const handleRefresh = async () => {
    await refresh()
    await redirect(resData.data.origin)
}

</script>
<template>
    <Head>
        <Title>{{ resData.data.origin }}</Title>
        <Meta name="description" :content="resData.data.origin"/>
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