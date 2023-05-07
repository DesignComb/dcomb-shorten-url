<script setup lang="ts">
import LoadingAnimation from "~/components/common/loadingAnimation.vue";

const route = useRoute()
const router = useRouter()

const {appBaseUrl, apiBaseUrl} = useRuntimeConfig().public

onMounted( async () => {
  await getUrlQueryParams()
  await login(<string>route.query.code)
});

const getUrlQueryParams = async () => {
  //router is async, so we wait for it to be ready
  await router.isReady()
  //once its ready we can access the query params
  console.log(route.query)
}

async function login(code:string) {
  await $fetch(`${apiBaseUrl}/api/ouath/google/login?code=${code}&scope=email+profile+https://www.googleapis.com/auth/userinfo.profile+https://www.googleapis.com/auth/userinfo.email+openid&authuser=0&prompt=consent`, {
    method: 'GET',
    credentials: 'include',
  })
      .then((res) => {
        console.log(res)
        // getUserInfo()
        router.push('/')
      })
      .finally(() => {
        // getUserInfo()
      })
}
//
// async function getUserInfo() {
//   await $fetch(`${apiBaseUrl}/api/user/info`, {
//     method: 'GET',
//     credentials: 'include',
//   })
//       .then((res) => {
//         console.log(res)
//         router.push('/')
//       })
// }
</script>

<template>
  <div class="w-screen h-screen flex justify-center items-center">
    <loading-animation/>
  </div>
</template>

<style scoped>
</style>