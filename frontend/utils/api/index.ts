const { appBaseUrl, apiBaseUrl } = useRuntimeConfig();

interface resData {
    code: string;
    data: object;
}

export default {
    getGoogleAuthUrl: async (): Promise<any> => {
        const { data, pending, error, refresh } = await useFetch(`${apiBaseUrl}/api/ouath/google/url`, {
            method: 'GET',
        });
        const resData = data.value as resData;
        return {...resData.data}
    },
    getUserInfo: async (): Promise<any> => {
        const { data, pending, error, refresh } = await useFetch(`${apiBaseUrl}/api/user/info`, {
            method: 'GET',
            credentials: 'include',
        });
        const resData = data.value as resData;
        return {...resData.data}
    },
    login: async (code: string): Promise<any> => {
        const { data, pending, error, refresh } = await useFetch(`${apiBaseUrl}/api/ouath/google/login?code=${code}`, {
            method: 'GET',
            credentials: 'include',
        });
        const resData = data.value as resData;
        return {...resData.data}
    },
    searchUrl: async (keyword: string): Promise<any> => {
        const { data, pending, error, refresh } = await useFetch(`${apiBaseUrl}/api/urlShorten/search?keyword=${keyword}`, {
            method: 'GET',
        });
        const resData = data.value as resData;
        return {...resData.data}
    },
    userSearchUrl: async (userID: number, keyword: string): Promise<any> => {
        const { data, pending, error, refresh } = await useFetch(`${apiBaseUrl}/api/user/${userID}/urlShorten/search?keyword=${keyword}`, {
            method: 'GET',
            credentials: 'include',
        });
        const resData = data.value as resData;
        return {...resData.data}
    },
    postUrl: async (urlObj: object): Promise<any> => {
        const { data, pending, error, refresh } = await useFetch(`${apiBaseUrl}/api/urlShorten`, {
            method: 'POST',
            body: urlObj,
        });
        const resData = data.value as resData;
        return {...resData.data}
    },
    userPostUrl: async (userID: number, userUrlObj: object): Promise<any> => {
        const { data, pending, error, refresh } = await useFetch(`${apiBaseUrl}/api/user/${userID}/urlShorten`, {
            method: 'POST',
            body: userUrlObj,
            credentials: 'include',
        });
        const resData = data.value as resData;
        return {...resData.data}
    },
    userPostImage: async (userID: number, imgObj: FormData): Promise<any> => {
        const { data, pending, error, refresh } = await useFetch(`${apiBaseUrl}/api/user/${userID}/image`, {
            method: 'POST',
            body: imgObj,
            credentials: 'include',
        });
        const resData = data.value as resData;
        return {...resData.data}
    },
};
