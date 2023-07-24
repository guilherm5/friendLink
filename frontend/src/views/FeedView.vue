<script setup lang="ts">
import FeedComponent from '@/components/FeedComponent.vue';
import NavComponent from '@/components/NavComponent.vue';
import NewPostComponent from '@/components/NewPostComponent.vue';
import StoryComponent from '@/components/StoryComponent.vue';
import { useNotificationStore } from '@/stores/global';
import { getPosts } from '@/services/PostService';
import { useAuthStore } from '@/stores/global';
import type { DefaultResponse } from '@/types/ApiService';
import { ref, onMounted } from 'vue';
import { usePostStore } from '../stores/global';
import type { Post } from '@/types/PostService';

const notificationStore = useNotificationStore()
const postsLoading = ref(false)
const authStore = useAuthStore()
const postStore = usePostStore()
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
</script>

<template>
  <main class="bg-neutral-900 container min-h-screen relative pb-4">
    <NavComponent />
    <div class="flex md:pt-[6rem] gap-4">
      <section class="h-16 hidden md:flex w-full max-w-xs bg-neutral-800 ml-4">
        
      </section>
      <section class="h-16 flex flex-col w-full items-center mx-auto">
        <StoryComponent />
        <NewPostComponent />
        <FeedComponent :postsLoading="postsLoading" />
        <div @click="postStore.updatePosts()" :class="(endOfPosts ? 'opacity-70' : 'opacity-0 pointer-events-none') + ' text-neutral-500 font-medium pb-[70px] pt-4 md:pb-6 cursor-pointer'" id="loadMorePosts">🌠 você chegou ao fim do feed. Atualizar</div>
      </section>
      <section class="hidden md:flex w-full max-w-xs bg-neutral-800 mr-4">
        
      </section>
    </div>
  </main>
</template>
