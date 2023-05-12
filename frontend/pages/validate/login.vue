<script setup lang="ts">
import LoadingAnimation from "~/components/common/loadingAnimation.vue";
import Api from "~/utils/api";

const route = useRoute()
const router = useRouter()

await router.isReady()

const response = await Api.login(<string>route.query.code)
if (response.data.value) {
    await router.push('/')
}

</script>

<template>
    <div class="w-screen h-screen flex justify-center items-center">
        <loading-animation/>
    </div>
</template>

<style scoped>
</style>