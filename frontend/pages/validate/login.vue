<script setup lang="ts">
import LoadingAnimation from "~/components/common/loadingAnimation.vue";
import Api from "~/utils/api";
import {useUserStore} from '~/store/user'

const route = useRoute()
const router = useRouter()

const user = useUserStore()

// await router.isReady()
await user.login(<string>route.query.code)
await router.push('/')

</script>

<template>
    <div class="w-screen h-screen flex justify-center items-center">
        <loading-animation/>
    </div>
</template>

<style scoped>
</style>