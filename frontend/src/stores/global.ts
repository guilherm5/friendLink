import { ref } from 'vue'
import { defineStore } from 'pinia'
import jwt_decode from 'jwt-decode'
import type { Auth } from '@/types/AuthService'
import type { Post } from '@/types/PostService'
import type { AlertNotification, NotificationProps } from '@/types/Notification'
import { getUser } from '@/services/UserService'
import type { DefaultResponse } from '@/types/ApiService'
import type { UserResponse } from '@/types/UserService'
import type { Story } from '@/types/StoryService'

export const useNotificationStore = defineStore('notification', () => {
  const notifications = ref<AlertNotification[]>([])
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
    try{
      const ls = localStorage.getItem('auth')
      if(ls){
        return JSON.parse(ls)
      }
      return undefined
    }catch(e){
      console.log('Error at get localStorage jwt\n', e)
      localStorage.removeItem('auth')
      return undefined
    }
  }

  function setAuth(_token: string | undefined, _refreshToken: string | undefined) {
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
    const currentUserStore = useCurrentUserStore()
    currentUserStore.clearCurrentUser()
    localStorage.removeItem('auth')
    setAuth(undefined, undefined)
  }

  return { auth, setAuth, removeAuth }
})

export const useCurrentUserStore = defineStore('currentUser', () => {
  const currentUser = ref<UserResponse | undefined>(undefined)
  function setCurrentUser(user: UserResponse) {
    currentUser.value = user
  }
  function clearCurrentUser() {
    currentUser.value = undefined
  }
  async function getCurrentUser() {
    const authStore = useAuthStore()
    const notificationStore = useNotificationStore()
    const res = await getUser(authStore.auth?.token) as DefaultResponse
    if (res.status && res.data) {
      setCurrentUser(res.data as UserResponse)
    }else{
      notificationStore.addNotification({
        type: 'error',
        body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado ao carregar as suas informações.',
      })

      clearCurrentUser()
    }
  }
  return { currentUser, setCurrentUser, clearCurrentUser, getCurrentUser }
})

export const usePostStore = defineStore('post', () => {
  const posts = ref<Post[]>([])
  const updatePostsFunction = ref<Function | null>(null)
  const updatePosts = (): void => updatePostsFunction.value ? updatePostsFunction.value() : null

  function addPosts(newPosts: Post[], start: boolean = false) {
    start ?  posts.value.unshift(...newPosts) : posts.value.push(...newPosts)
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
  function removePost(id_post: number) {
    const index = posts.value.findIndex(p => p.id_post === id_post)
    posts.value.splice(index, 1)
  }
  return { posts, addPosts, onUpdatePosts, updatePosts, setPosts, updatePostValue, removePost }
})
export const useStoryStore = defineStore('story', () => {
  const stories = ref<Story[]>([])
  const updateStoriesFunction = ref<Function | null>(null)
  const updateStories = (): void => updateStoriesFunction.value ? updateStoriesFunction.value() : null

  function onUpdateStories(updateFunction: Function) {
    updateStoriesFunction.value = updateFunction
  }
  function addStories(newStories: Story[], start: boolean = false) {
    start ?  stories.value.unshift(...newStories) : stories.value.push(...newStories)
  }
  function setStories(newStories: Story[]) {
    stories.value = newStories
  }
  return { stories, addStories, setStories, updateStories, onUpdateStories }
})

export const useRedirectStore = defineStore('redirect', () => {
  const redirectTo = ref<string | undefined>()
  const message = ref<string | undefined>()

  function setRedirectTo(path: string, redirectMessage: string) {
    redirectTo.value = path
    message.value = redirectMessage
  }
  function clearMessage() {
    message.value = undefined
  }
  function clearRedirectTo() {
    redirectTo.value = undefined
  }
  return { redirectTo, setRedirectTo, message, clearMessage, clearRedirectTo }
})