<script setup lang="ts">
import FeedComponent from '@/components/FeedComponent.vue';
import NavComponent from '@/components/NavComponent.vue';
import CreateNewContentComponent from '@/components/CreateNewContentComponent.vue';
import StoryComponent from '@/components/StoryComponent.vue';
import { useNotificationStore } from '@/stores/global';
import { getPosts } from '@/services/PostService';
import { getStories } from '@/services/StoryService';
import { useAuthStore } from '@/stores/global';
import type { DefaultResponse } from '@/types/ApiService';
import { ref, onMounted } from 'vue';
import { usePostStore, useStoryStore } from '../stores/global';
import type { Post } from '@/types/PostService';
import type { StoryResponse, StoryItem, Story } from '@/types/StoryService';

const notificationStore = useNotificationStore()
const postsLoading = ref(false)
const storiesLoading = ref(false)
const authStore = useAuthStore()
const postStore = usePostStore()
const storyStore = useStoryStore()
const postsPerPage = ref(10)
const lastPostId = ref(0)
const endOfPosts = ref(false)

const loadMorePosts = async (scrollSensor: HTMLElement | null) => {
  if (scrollSensor && window.scrollY >= scrollSensor.offsetTop - window.innerHeight && !postsLoading.value && !endOfPosts.value) {
    await fetchNewPosts()
  }
}
onMounted(async () => {
  await fetchNewPosts()
  await fetchStories()
  const scrollSensor = document.getElementById('loadMorePosts')
  window.addEventListener('scroll', () => loadMorePosts(scrollSensor))

  postStore.onUpdatePosts(async () => {
    window.removeEventListener('scroll', async () => await loadMorePosts(scrollSensor))
    window.scrollTo(0, 0)
    endOfPosts.value = false
    const posts = await fetchPosts()
    if(posts){
      postStore.setPosts(posts)
    }
  })
  storyStore.onUpdateStories(async () => {
    await fetchStories()
  })
})

async function fetchNewPosts() {
  const newPosts = await fetchPosts()
  if(newPosts){
    postStore.addPosts(newPosts)
  }
}
async function fetchPosts(): Promise<false | Post[]>{
  postsLoading.value = true
  const res = await getPosts(authStore.auth?.token, lastPostId.value, postsPerPage.value) as DefaultResponse
  if (res.status) {
    if (res.data === null) {
      endOfPosts.value = true
      postsLoading.value = false
      lastPostId.value = 0
      return false
    }else{
      endOfPosts.value = false
      lastPostId.value = res.data[res.data.length - 1].id_post
      postsLoading.value = false
      return res.data
    }
  }else{
    notificationStore.addNotification({
      type: 'error',
      body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado ao carregar os posts.',
    })
    postsLoading.value = false
    return false
  }
}
async function fetchStories(){
  storiesLoading.value = true
  const res = await getStories(authStore.auth?.token) as DefaultResponse
  if (res.status && res.data) {
    const storiesResponse = res.data as StoryResponse[]
    const stories = storiesResponse.map(story => {
      return {...story, stories: JSON.parse(story.stories as string) as StoryItem[]} as Story
    })
    storyStore.setStories(stories)
    storiesLoading.value = false
  }else{
    if(res.data === null){
      storiesLoading.value = false
      return
    }
    notificationStore.addNotification({
      type: 'error',
      body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado ao carregar os stories.',
    })
    storiesLoading.value = false
  }
}
</script>

<template>
  <main class="bg-neutral-900 container min-h-screen relative pb-4">
    <NavComponent />
    <div class="flex md:pt-[6rem] gap-4">
      <section class="h-16 hidden md:flex w-full max-w-xs bg-neutral-800 ml-4">
        
      </section>
      <section class="h-16 flex flex-col w-full max-w-[680px] items-center mx-auto" v-auto-animate>
        <StoryComponent :storiesLoading="storiesLoading" />
        <CreateNewContentComponent stage="button" />
        <FeedComponent :postsLoading="postsLoading" />
        <div @click="postStore.updatePosts()" :class="(endOfPosts ? 'opacity-70' : 'opacity-0 pointer-events-none') + ' text-neutral-500 font-medium pb-[70px] pt-4 md:pb-6 cursor-pointer'" id="loadMorePosts">🌠 você chegou ao fim do feed. Atualizar</div>
      </section>
      <section class="hidden md:flex w-full max-w-xs bg-neutral-800 mr-4">
        
      </section>
    </div>
  </main>
</template>
