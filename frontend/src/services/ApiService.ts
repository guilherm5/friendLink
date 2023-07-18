
import axios from 'axios'
import { useAuthStore, useRedirectStore } from '../stores/global';
import router from '@/router';

const APP_API_ENDPOINT = import.meta.env.VITE_API_ENDPOINT
const ApiService = axios.create({
    baseURL: APP_API_ENDPOINT,
    headers: {
        'Content-Type': 'multipart/form-data',
        'Accept': 'application/json',
    }
})
ApiService.interceptors.request.use( async (config) => {
    const authStore = useAuthStore()
    const redirectStore = useRedirectStore()
    const route = router.currentRoute.value

    const tokenExpired = (exp: number) => {
        const now = new Date()
        const expDate = new Date(exp * 1000)
        return now > expDate
    }

    const refreshToken = async (refreshToken: string): Promise<string | false> => {
        return await ApiService.post('/refresh', null, { headers: { Authorization: refreshToken } })
        .then(res => {
            const response = res.data as { new_token: string }
            return response.new_token
        })
        .catch((error) => {
            console.log(error)
            return false
        })
    }

    if(authStore.auth?.token) {
        if(tokenExpired(authStore.auth.exp)) {
            if(!authStore.auth.refreshToken){throw new axios.Cancel('Operation canceled by the user.')}
            const newToken = await refreshToken(authStore.auth.refreshToken)
            if(newToken){
                authStore.setAuth(newToken as string, authStore.auth.refreshToken)
            }else{
                authStore.removeAuth()
                redirectStore.setRedirectTo(route.name as string, 'Sua sessão expirou, faça login novamente')
                router.push({ name: 'signin' })
                throw new axios.Cancel('Operation canceled by the user.')
            }
        }
        config.headers.Authorization = authStore.auth.token
        return config
    }else{
        return config
    }
}, error => {
   return Promise.reject(error) 
})


ApiService.interceptors.response.use( response => {
    return response
}, error => {
    if(error.response.status === 401) {
        const authStore = useAuthStore()
        const redirectStore = useRedirectStore()
        const route = router.currentRoute.value
        authStore.removeAuth()
        redirectStore.setRedirectTo(route.name as string, 'Sua sessão expirou, faça login novamente')
        router.push({ name: 'signin' })
    }
    return Promise.reject(error)
})

export {
    ApiService,
}