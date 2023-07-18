import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import jwt_decode from 'jwt-decode'
import type { Auth } from '@/types/AuthService'
import type { Post } from '@/types/PostService'

type Notification = {
  id: string,
  type: 'success' | 'error',
  body: string,
  footer?: string
}
type NotificationProps = Omit<Notification, 'id'>

export const useCounterStore = defineStore('counter', () => {
  const count = ref(0)
  const doubleCount = computed(() => count.value * 2)
  function increment() {
    count.value++
  }

  return { count, doubleCount, increment }
})

export const useNotificationStore = defineStore('notification', () => {
  const notifications = ref<Notification[]>([])
  function addNotification(notification: NotificationProps) {
    notifications.value.unshift({ 
      id: Date.now() + Math.random().toString().slice(2),
      ...notification 
    })
  }
  function removeNotification(id: string) {
    notifications.value = notifications.value.filter(n => n.id !== id)
  }
  return { notifications, addNotification, removeNotification }
})

export const useAuthStore = defineStore('auth', () => {
  const auth = ref<Auth | undefined>(getAuth())
  
  function getAuth() {
    const ls = localStorage.getItem('auth')
    if(ls){
      return JSON.parse(ls)
    }
    return undefined
  }

  function setAuth(_token: string | undefined, _refreshToken: string | undefined = 'refreshToken should not have a default value') {
    if(_token && _refreshToken) {
      const decoded = jwt_decode<{exp: number, Nome: string, Issuer: number}>(_token)
      auth.value = {
        token: _token,
        exp: decoded.exp,
        Nome: decoded.Nome,
        ID: decoded.Issuer,
        refreshToken: _refreshToken
      }
    }else{
      auth.value = undefined
    }
    localStorage.setItem('auth', JSON.stringify(auth.value))
  }

  function removeAuth() {
    localStorage.removeItem('auth')
    setAuth(undefined)
  }

  return { auth, setAuth, removeAuth }
})

export const usePostStore = defineStore('post', () => {
  const posts = ref<Post[]>([])
  const updatePostsFunction = ref<Function | null>(null)
  const updatePosts = (): void => updatePostsFunction.value ? updatePostsFunction.value() : null

  function addPosts(newPosts: Post[]) {
    posts.value.push(...newPosts)
  }
  function setPosts(newPosts: Post[]) {
    posts.value = newPosts
  }
  function onUpdatePosts(updateFunction: Function) {
    updatePostsFunction.value = updateFunction
  }
  function updatePostValue(post: Post) {
    const index = posts.value.findIndex(p => p.id_post === post.id_post)
    posts.value[index] = post
  }
  return { posts, addPosts, onUpdatePosts, updatePosts, setPosts, updatePostValue }
})