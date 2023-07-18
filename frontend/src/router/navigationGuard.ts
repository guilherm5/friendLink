import type { RouteLocationNormalized } from "vue-router";
import { useAuthStore } from '@/stores/global';
import { ApiService } from "@/services/ApiService";
import type { LoginResponse } from '@/types/AuthService';
import type { AxiosError } from "axios";

type RouteLocationNormalizedExtended = RouteLocationNormalized & {
    meta?: {
      requireAuth?: boolean
    }
}

const navigationGuard = async (to: RouteLocationNormalizedExtended) => {
    to.meta?.title && (document.title = to.meta.title as string + ' - ' + document.title)
    checkAuth(to)
}

const checkAuth = async (to: RouteLocationNormalizedExtended) => {
    const authStore = useAuthStore()
    if(to.meta?.requireAuth) {
        if(!authStore.auth?.refreshToken){return { name: 'signin' }}
        const newToken = await refreshToken(authStore.auth?.refreshToken)
        if(newToken){ 
            authStore.setAuth(newToken as string)
        }else{
            authStore.setAuth(undefined)
            return { name: 'signin' }
        }
    }else{
        if(authStore.auth?.token) {
            return { name: 'home' }
        }
    }
}

const refreshToken = async (refreshToken: string): Promise<string | false> => {
    return await ApiService.post('/auth/refresh', null, { headers: { Authorization: refreshToken } })
    .then(res => {
        const response = res.data as LoginResponse
        return response.data?.logged
    })
    .catch((error: AxiosError) => {
        console.log(error)
        return false
    })
}

export default navigationGuard