<script setup lang="ts">
import { Chat32Regular, Heart32Regular } from '@vicons/fluent';
import { DotsVertical, Repeat } from '@vicons/tabler';
import TooltipContainerComponent from './TooltipContainerComponent.vue';
import { usePostStore } from '../stores/global';
// import SkeletonComponent from './SkeletonComponent.vue';

defineProps<{
    postsLoading: boolean,
}>()

const postStore = usePostStore()
</script>

<template>
    <div class="mt-4 w-full px-2 md:px-0">
        <div v-for="post in postStore.posts" class="w-full max-w-3xl mx-auto my-4 bg-neutral-800 p-4 rounded-lg" :key="post.id_post">
            <div class="flex flex-col max-w-[720px]">
                <!-- post header -->
                <div class="flex gap-2 items-center">
                    <router-link :to="{name: 'profile', params:{arroba: post.arroba}}">
                        <img :src="post.link_perfil" :alt="'Foto de perfil de ' + post.nome" class="w-10 h-10 md:w-16 md:h-16 p-1 rounded-xl md:rounded-3xl bg-neutral-900 object-cover">
                    </router-link>

                    <router-link :to="{name: 'profile', params:{arroba: post.arroba}}">
                        <p class="text-neutral-400 text-sm font-thin md:mb-1">{{ post.arroba }}</p>
                        <div class="flex items-center">
                            <p class="text-white font-sans font-medium inline-block">{{ post.nome }}</p>
                            <div class="mx-2 inline-block w-1 h-1 bg-yellow-400 rounded-full"></div>
                            <p class="text-yellow-400 inline-block text-[12px] font-thin">2 dias atrás</p>
                        </div>
                    </router-link>

                    <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group ml-auto">
                        <DotsVertical class="h-4 md:h-5"/>
                    </button>
                </div>
                <!-- post text -->
                <div class="post-text text-white my-4 text-sm">
                    <p v-if="post.post_texto">{{ post.post_texto }}</p>
                </div>
                <!-- post image -->
                <div v-if="post.post_imagem" class="rounded-lg overflow-hidden w-full max-w-[720px] max-h-[400px] mx-auto">
                    <img 
                        :src="post.post_imagem"
                        class="w-full h-full object-cover"
                    />
                </div>
                <!-- post iteractions -->
                <div class="flex my-4 gap-2">
                    <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group relative focus:outline-none">
                        <!-- <Chat32Filled class="text-yellow-400" height="28"/> -->
                        <Heart32Regular height="24" class="hover:text-yellow-400"/>
                        <TooltipContainerComponent text="Curtir" width="fit"/>
                    </button>
                    <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group relative focus:outline-none">
                        <!-- <Chat32Filled class="text-yellow-400" height="28"/> -->
                        <Chat32Regular height="24" class="hover:text-yellow-400"/>
                        <TooltipContainerComponent text="Ver comentários" width="fit"/>
                    </button>
                    <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group relative focus:outline-none">
                        <!-- <Chat32Filled class="text-yellow-400" height="28"/> -->
                        <Repeat height="24" class="hover:text-yellow-400"/>
                        <TooltipContainerComponent text="Repostar" width="fit"/>
                    </button>
                </div>
                <!-- write your comment -->
                <div class="flex gap-4 items-center md:mt-2">
                    <img src="@/assets/user_default_foto.jpeg" alt="Foto de perfil" class="w-8 h-8 md:w-10 md:h-10 rounded-full bg-neutral-900 object-cover">
                    <input type="text" class="bg-neutral-700 placeholder:text-neutral-500 text-white rounded-lg p-2 w-full" placeholder="Adicione um comentário...">
                </div>
            </div>
        </div>

        <div v-if="postsLoading" class="w-full p-4 rounded-lg text-neutral-500 text-center font-medium pb-[70px] pt-4 md:pb-6">
            <p>Carregando...</p>
        </div>
        <!-- <div v-if="postsLoading" class="w-full h-[300px] max-w-3xl mx-auto bg-neutral-800 p-4 rounded-lg flex flex-col">
            <div class="flex">
                <SkeletonComponent w="50px" h="50px" />
                <div class="flex flex-col ml-4 justify-center">
                    <SkeletonComponent class="mb-2" w="80px" h="15px" />
                    <SkeletonComponent class="" w="150px" h="15px" />
                </div>
            </div>
            <SkeletonComponent class="mt-4 mb-2" w="90%" h="20px" />
            <SkeletonComponent w="70%" h="20px" />
            <SkeletonComponent class="mt-4" w="100%" h="100%" />
        </div> -->
    </div>
</template>
