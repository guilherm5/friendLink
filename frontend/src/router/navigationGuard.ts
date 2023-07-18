import type { RouteLocationNormalized } from "vue-router";
import { useAuthStore, useRedirectStore } from '@/stores/global';
import { ApiService } from "@/services/ApiService";
import type { AxiosError } from "axios";

type RouteLocationNormalizedExtended = RouteLocationNormalized & {
    meta?: {
      requireAuth?: boolean
    }
}

const navigationGuard = async (to: RouteLocationNormalizedExtended) => {
    to.meta?.title && (document.title = to.meta.title as string + ' - ' + document.title)
    return checkAuth(to)
}

const refreshToken = async (refreshToken: string): Promise<string | false> => {
    return await ApiService.post('/refresh', null, { headers: { Authorization: refreshToken } })
    .then(res => {
        const response = res.data as { new_token: string }
        return response.new_token
    })
    .catch((error: AxiosError) => {
        console.log(error)
        return false
    })
}

const checkAuth = async (to: RouteLocationNormalizedExtended) => {
    const authStore = useAuthStore()
    const redirectStore = useRedirectStore()
    if(to.meta?.requireAuth) {
        if(!authStore.auth?.refreshToken){return { name: 'signin' }}
        const newToken = await refreshToken(authStore.auth?.refreshToken)
        if(newToken){ 
            authStore.setAuth(newToken as string, authStore.auth?.refreshToken)
        }else{
            authStore.removeAuth()
            redirectStore.setRedirectTo(to.name as string, 'Sua sessão expirou, faça login novamente')
            return { name: 'signin' }
        }
    }else{
        if(authStore.auth?.token) {
            return { name: 'home' }
        }
    }
}

export default navigationGuard