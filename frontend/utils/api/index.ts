const {appBaseUrl, apiBaseUrl} = useRuntimeConfig().public

export default {
     getGoogleAuthUrl: () :any => {
         return useFetch(`${apiBaseUrl}/api/ouath/google/url`,{
             method: 'GET',
         })
    },
    getUserInfo: () :any => {
        return useFetch(`${apiBaseUrl}/api/user/info`,{
            method: 'GET',
            credentials: 'include',
        })
    }
}