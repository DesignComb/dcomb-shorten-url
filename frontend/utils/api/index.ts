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
    },
    searchUrl: (keyword:string) :any => {
        return useFetch(`${apiBaseUrl}/api/urlShorten/search?keyword=${keyword}`, {
            method: 'GET',
        })
    },
    userSearchUrl: (userID:number,keyword:string) :any => {
        return useFetch(`${apiBaseUrl}/api/user/${{userID}}/urlShorten/search?keyword=${keyword}`, {
            method: 'GET',
            credentials: 'include',
        })
    },
    postUrl: (urlObj:object) :any => {
        return useFetch(`${apiBaseUrl}/api/urlShorten`, {
            method: 'POST',
            body:urlObj
        })
    },
    userPostUrl: (userID:number,userUrlObj:object) :any => {
        return useFetch(`${apiBaseUrl}/api/user/${{userID}}/urlShorten`, {
            method: 'POST',
            body:userUrlObj,
            credentials: 'include',
        })
    },
}