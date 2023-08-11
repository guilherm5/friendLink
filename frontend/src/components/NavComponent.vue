<script setup lang="ts">
import { Search } from '@vicons/ionicons5';
import { 
  Home32Regular, Home32Filled, Chat32Filled, Chat32Regular, 
  Heart32Filled, Heart32Regular, Settings32Regular, Settings32Filled, 
  Search32Filled, Search32Regular, AddSquare24Regular, CaretDown12Filled 
} from '@vicons/fluent';
import router from '@/router';
import { computed } from 'vue';
import TooltipContainerComponent from '@/components/TooltipContainerComponent.vue';
import { usePostStore, useAuthStore, useCurrentUserStore, useStoryStore } from '../stores/global';

const siteUrl = window.location.href
const currentRoute = computed(() => router.currentRoute.value.name as string)
const postStore = usePostStore()
const storyStore = useStoryStore()
const authStore = useAuthStore()
const currentUserStore = useCurrentUserStore()
const handleLogout = () => {
  authStore.removeAuth()
  router.push({name: 'signin'})
}
const updateFeed = () => {
  postStore.updatePosts()
  storyStore.updateStories()
}
</script>

<template>
  <nav class="z-50 w-full flex sticky left-0 top-0 md:hidden px-4 py-2 bg-neutral-900">
    <ul class="flex w-full justify-between items-center">
      <li>
        <a :href="siteUrl"><img src="@/assets/logo_horizontal.svg" width="90" /></a>
      </li>
      <li class="group relative" id="nav-config2">
        <router-link :to="{name:'configurations'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Settings32Filled v-if="(currentRoute === 'configurations')" class="text-yellow-400" height="28"/>
          <Settings32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Configurações" width="fit" side="left"/>
        </router-link>

        <div class="opacity-0 pointer-events-none group-hover:opacity-100 transition-opacity group-hover:pointer-events-auto absolute right-0 z-10 mt-0 w-56 origin-top-right rounded-md bg-black border border-neutral-700 shadow-lg focus:outline-none" role="menu" aria-orientation="vertical" aria-labelledby="nav-config2" tabindex="-1">
          <div class="py-1 text-yellow-400 flex flex-col items-center text-center" role="none">
            <router-link v-if="currentUserStore.currentUser" :to="{name: 'profile', params:{arroba: currentUserStore.currentUser.arroba}}" class="w-full px-4 py-2 text-sm hover:bg-neutral-800" role="menuitem" tabindex="-1">
              Meu perfil
            </router-link>
            <button @click="handleLogout" class="w-full px-4 py-2 text-sm hover:bg-neutral-800" role="menuitem" tabindex="-1">Sair</button>
          </div>
        </div>
      </li>
    </ul>
  </nav>
  <nav class="z-50 w-full flex fixed left-0 bottom-0 md:top-0 md:bottom-auto px-4 pb-1 md:py-2 bg-black md:bg-neutral-900">
    <ul class="hidden md:flex w-full max-w-xs justify-start">
      <li class="mr-8">
        <a :href="siteUrl"><img src="@/assets/logo.svg" width="42" class="mx-auto" /></a>
      </li>
      <li class="w-full block">
          <div class="relative">
            <label for="search" class="text-neutral-400 absolute top-1/2 -translate-y-1/2 left-2">
              <Search height="20"/>
            </label>
            <input type="text" id="search" placeholder="explorar..." class="bg-neutral-700 text-neutral-400 placeholder:text-neutral-500 rounded-lg p-2 pl-8 w-full">
          </div>
      </li>
    </ul>

    <ul class="flex w-full justify-between md:justify-center py-2 md:gap-12 mx-auto">
      <li class="relative">
        <router-link @click="currentRoute === 'home' ? updateFeed() : false" :to="{name:'home'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Home32Filled v-if="(currentRoute === 'home')" class="text-yellow-400" height="28"/>
          <Home32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Feed" width="fit"/>
        </router-link>
        <div v-if="(currentRoute === 'home')" class="hidden md:block w-2 h-2 rounded-full bg-yellow-400 absolute top-30 left-1/2 -translate-x-1/2"></div>
      </li>
      <li class="relative md:hidden">
        <router-link :to="{name:'search'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Search32Filled v-if="(currentRoute === 'search')" class="text-yellow-400" height="28"/>
          <Search32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Explorar" width="fit"/>
        </router-link>
        <div v-if="(currentRoute === 'search')" class="hidden md:block w-2 h-2 rounded-full bg-yellow-400 absolute top-30 left-1/2 -translate-x-1/2"></div>
      </li>
      <li class="relative md:hidden">
        <AddSquare24Regular height="28" class="text-neutral-300 md:text-neutral-400 active:text-yellow-400"/>
      </li>
      <li class="relative">
        <router-link :to="{name:'inbox'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Chat32Filled v-if="(currentRoute === 'inbox')" class="text-yellow-400" height="28"/>
          <Chat32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Inbox" width="fit"/>
        </router-link>
        <div v-if="(currentRoute === 'inbox')" class="hidden md:block w-2 h-2 rounded-full bg-yellow-400 absolute top-30 left-1/2 -translate-x-1/2"></div>
      </li>
      <li class="relative">
        <router-link :to="{name:'notifications'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Heart32Filled v-if="(currentRoute  === 'notifications')" class="text-yellow-400" height="28"/>
          <Heart32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Notificações" width="fit"/>
        </router-link>
        <div v-if="(currentRoute === 'notifications')" class="hidden md:block w-2 h-2 rounded-full bg-yellow-400 absolute top-30 left-1/2 -translate-x-1/2"></div>
      </li>
    </ul>

    <ul class="hidden md:flex lg:w-full lg:max-w-xs justify-end">
      <li class="group relative" id="nav-config">
        <div class="text-neutral-100 group flex gap-2 items-center pl-1 pr-2 py-1 bg-neutral-800 rounded-lg cursor-pointer">
          <img :src="currentUserStore.currentUser?.link_perfil" alt="Sua foto de perfil" class="w-6 h-6 rounded-lg p-[3px] bg-neutral-900 object-cover">
          {{ currentUserStore.currentUser?.nome }}
          <CaretDown12Filled height="20"/>
        </div>
        <div class="opacity-0 pointer-events-none group-hover:opacity-100 transition-opacity group-hover:pointer-events-auto absolute right-0 z-10 mt-0 w-56 origin-top-right rounded-md bg-black border border-neutral-700 shadow-lg focus:outline-none" role="menu" aria-orientation="vertical" aria-labelledby="nav-config" tabindex="-1">
            <div class="py-1 text-yellow-400 flex flex-col items-center text-center" role="none">
              <router-link v-if="currentUserStore.currentUser" :to="{name: 'profile', params:{arroba: currentUserStore.currentUser.arroba}}" class="w-full px-4 py-2 text-sm hover:bg-neutral-800" role="menuitem" tabindex="-1">
                Meu perfil
              </router-link>
              <button @click="handleLogout" class="w-full px-4 py-2 text-sm hover:bg-neutral-800" role="menuitem" tabindex="-1">Sair</button>
            </div>
        </div>
      </li>
    </ul>
  </nav>
</template>