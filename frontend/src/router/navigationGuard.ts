import type { RouteLocationNormalized } from "vue-router";
import { useAuthStore } from '@/stores/global';

type RouteLocationNormalizedExtended = RouteLocationNormalized & {
    meta?: {
      requireAuth?: boolean
    }
}

const navigationGuard = async (to: RouteLocationNormalizedExtended) => {
    to.meta?.title && (document.title = to.meta.title as string + ' - ' + document.title)
    return checkAuth(to)
}

const checkAuth = async (to: RouteLocationNormalizedExtended) => {
    const authStore = useAuthStore()
    if(to.meta?.requireAuth) {
        if(!authStore.auth?.token || !authStore.auth?.token){
            return { name: 'signin' }
        }
    }else{
        if(authStore.auth?.token) {
            return { name: 'home' }
        }
    }
}

export default navigationGuard