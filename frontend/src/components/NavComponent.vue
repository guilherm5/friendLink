<script setup lang="ts">
import { Search } from '@vicons/ionicons5';
import { 
  Home32Regular, Home32Filled, Chat32Filled, Chat32Regular, 
  Heart32Filled, Heart32Regular, Settings32Regular, Settings32Filled, 
  Search32Filled, Search32Regular, AddSquare24Regular 
} from '@vicons/fluent';
import router from '@/router';
import { computed } from 'vue';
import TooltipContainerComponent from '@/components/TooltipContainerComponent.vue';

const siteUrl = window.location.href
const currentRoute = computed(() => router.currentRoute.value.name)
</script>

<template>
  <nav class="z-10 w-full flex sticky left-0 top-0 md:hidden px-4 py-2 bg-black">
    <ul class="flex w-full justify-between items-center">
      <li>
        <a :href="siteUrl"><img src="@/assets/logo_horizontal.svg" width="90" /></a>
      </li>
      <li>
        <router-link :to="{name:'configurations'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Settings32Filled v-if="(currentRoute as string === 'configurations')" class="text-yellow-400" height="28"/>
          <Settings32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Configurações" width="fit" side="left"/>
        </router-link>
      </li>
    </ul>
  </nav>
  <nav class="z-10 w-full flex fixed left-0 bottom-0 md:top-0 md:bottom-auto px-4 pb-1 md:py-2 bg-black md:bg-neutral-900">
    <ul class="hidden md:flex w-full max-w-xs justify-start">
      <li class="mr-8">
        <a :href="siteUrl"><img src="@/assets/logo.svg" width="42" class="mx-auto" /></a>
      </li>
      <li class="w-full block">
        <div class="relative">
            <label for="search" class="text-neutral-400 absolute top-1/2 -translate-y-1/2 left-2">
              <Search height="20"/>
            </label>
            <input type="text" id="search" placeholder="explorar..." class="bg-neutral-700 text-neutral-400 placeholder:text-neutral-500 rounded-2xl p-2 pl-8 w-full">
          </div>
      </li>
    </ul>

    <ul class="flex w-full justify-between md:justify-center py-2 md:gap-12 mx-auto">
      <li class="relative">
        <router-link :to="{name:'home'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Home32Filled v-if="(currentRoute as string === 'home')" class="text-yellow-400" height="28"/>
          <Home32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Feed" width="fit"/>
        </router-link>
        <div v-if="(currentRoute as string === 'home')" class="hidden md:block w-2 h-2 rounded-full bg-yellow-400 absolute top-30 left-1/2 -translate-x-1/2"></div>
      </li>
      <li class="relative md:hidden">
        <router-link :to="{name:'search'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Search32Filled v-if="(currentRoute as string === 'search')" class="text-yellow-400" height="28"/>
          <Search32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Explorar" width="fit"/>
        </router-link>
        <div v-if="(currentRoute as string === 'search')" class="hidden md:block w-2 h-2 rounded-full bg-yellow-400 absolute top-30 left-1/2 -translate-x-1/2"></div>
      </li>
      <li class="relative md:hidden">
        <AddSquare24Regular height="28" class="text-neutral-300 md:text-neutral-400 active:text-yellow-400"/>
      </li>
      <li class="relative">
        <router-link :to="{name:'inbox'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Chat32Filled v-if="(currentRoute as string === 'inbox')" class="text-yellow-400" height="28"/>
          <Chat32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Inbox" width="fit"/>
        </router-link>
        <div v-if="(currentRoute as string === 'inbox')" class="hidden md:block w-2 h-2 rounded-full bg-yellow-400 absolute top-30 left-1/2 -translate-x-1/2"></div>
      </li>
      <li class="relative">
        <router-link :to="{name:'notifications'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Heart32Filled v-if="(currentRoute as string === 'notifications')" class="text-yellow-400" height="28"/>
          <Heart32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Notificações" width="fit"/>
        </router-link>
        <div v-if="(currentRoute as string === 'notifications')" class="hidden md:block w-2 h-2 rounded-full bg-yellow-400 absolute top-30 left-1/2 -translate-x-1/2"></div>
      </li>
    </ul>

    <ul class="hidden md:flex lg:w-full lg:max-w-xs justify-end">
      <li class="relative">
        <router-link :to="{name:'configurations'}" class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group">
          <Settings32Filled v-if="(currentRoute as string === 'configurations')" class="text-yellow-400" height="28"/>
          <Settings32Regular v-else height="28" class="hover:text-yellow-400"/>
          <TooltipContainerComponent text="Configurações" width="fit" side="left"/>
        </router-link>
        <div v-if="(currentRoute as string === 'configurations')" class="hidden md:block w-2 h-2 rounded-full bg-yellow-400 absolute top-30 left-1/2 -translate-x-1/2"></div>
      </li>
    </ul>
  </nav>
</template>
