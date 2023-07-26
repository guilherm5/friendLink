import type { RouteLocationNormalized } from "vue-router";
import { useAuthStore, useCurrentUserStore } from '@/stores/global';

type RouteLocationNormalizedExtended = RouteLocationNormalized & {
    meta?: {
      requireAuth?: boolean
    }
}

const publicRoutes = ['signin', 'signup']

const navigationGuard = async (to: RouteLocationNormalizedExtended) => {
    if(to.meta?.title){
        document.title = to.meta.title as string + ' - ' + document.title
    }else{
        document.title = import.meta.env.VITE_APP_NAME
    }
    return checkAuth(to)
}

const checkAuth = async (to: RouteLocationNormalizedExtended) => {
    const authStore = useAuthStore()
    const currentUserStore = useCurrentUserStore()
    const token = authStore.auth?.token
    const refreshToken = authStore.auth?.refreshToken
    const requireAuth = to.meta?.requireAuth
    const routeName = to.name

    if(requireAuth && !token){
        return { name: 'signin' }
    }else if(!requireAuth && token){
        return { name: 'home' }
    }else if(token && refreshToken){
        if(routeName && publicRoutes.includes(routeName as string)){
            return { name: 'home' }
        }else{
            if(!currentUserStore.currentUser && token && refreshToken){
                await currentUserStore.getCurrentUser()
            }

            return true
        }
    }else{
        return true
    }
}

export default navigationGuard