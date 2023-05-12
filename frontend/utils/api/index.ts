const {appBaseUrl, apiBaseUrl} = useRuntimeConfig()

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
    },
    login: (code:string) :any => {
        return useFetch(`${apiBaseUrl}/api/ouath/google/login?code=${code}`, {
            method: 'GET',
            credentials: 'include',
        })
    }
}