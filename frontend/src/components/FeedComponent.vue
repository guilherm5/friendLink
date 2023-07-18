<script setup lang="ts">
import { Chat32Regular, Heart32Filled, Heart32Regular } from '@vicons/fluent';
import { DotsVertical, Repeat } from '@vicons/tabler';
import TooltipContainerComponent from './TooltipContainerComponent.vue';
import { usePostStore, useAuthStore, useNotificationStore } from '../stores/global';
import SkeletonComponent from './SkeletonComponent.vue';
import { getComments, likePost, unlikePost } from '@/services/PostService';
import type { Post, CommentResponse } from '@/types/Post';
import type { DefaultResponse } from '@/types/ApiService';
import { createComment } from '@/services/PostService';
import { ref } from 'vue';

defineProps<{
    postsLoading: boolean,
}>()

const postStore = usePostStore()
const authStore = useAuthStore()
const notificationStore = useNotificationStore()
const newComment = ref('')

const interaction = (number: number) => {
    let interaction: string
    if (number > 0) {
        if(number > 1000) {
            interaction =  (number / 1000).toFixed(1) + 'k'
        }else{
            interaction = number.toString()
        }
        return interaction
    }
    return false
}
const like = (post: Post) => {
    if(post.curtiu){
        unlikePost(authStore.auth?.token, post.id_post)
        post.curtiu = false
        post.qtde_curtida--
    }else{
        likePost(authStore.auth?.token, post.id_post)
        post.curtiu = true
        post.qtde_curtida++
    }
}
const listComments = async (post: Post) => {
    if(post.carregadoUmaVez){
        post.comentarios = []
        post.carregadoUmaVez = false
        return
    }
    if(post.qtde_comentario === 0){return}
    post.loadingComentarios = true
    const res = await getComments(authStore.auth?.token, post.id_post) as DefaultResponse
    if (res.status) {
        const comments = res.data as CommentResponse[]
        const userComments = comments.filter(comment => comment.id_usuario_cmt === authStore.auth?.ID)
        const otherComments = comments.filter(comment => comment.id_usuario_cmt !== authStore.auth?.ID)
        post.comentarios = [...userComments, ...otherComments]
        post.carregadoUmaVez = true
    }else{
        notificationStore.addNotification({
            type: 'error',
            body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado.',
        })
    }
    post.loadingComentarios = false
}
const handleNewComment = async (post: Post) => {
    if(newComment.value.length === 0){return false}
    const res = await createComment(authStore.auth?.token, post.id_post, newComment.value) as DefaultResponse
    if (res.status) {
        post.comentarios ? post.comentarios.unshift(res.data) : post.comentarios = [res.data]
        post.qtde_comentario++
        newComment.value = ''
    }else{
        notificationStore.addNotification({
            type: 'error',
            body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado.',
        })
    }
}
const timeAgo = (date: string) => {
    const now = new Date()
    const postDate = new Date(date)
    const diff = now.getTime() - postDate.getTime()
    const diffDays = Math.floor(diff / (1000 * 60 * 60 * 24))
    const diffHours = Math.floor(diff / (1000 * 60 * 60))
    const diffMinutes = Math.floor(diff / (1000 * 60))
    const diffSeconds = Math.floor(diff / (1000))
    if(diffDays > 0){
        return diffDays === 1 ? '1d atrás' : diffDays + 'd atrás'
    }else if(diffHours > 0){
        return diffHours === 1 ? '1h atrás' : diffHours + 'h atrás'
    }else if(diffMinutes > 0){
        return diffMinutes === 1 ? '1m atrás' : diffMinutes + 'm atrás'
    }else{
        return Math.abs(diffSeconds) + 's atrás'
    }
}
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
                <div class="post-text text-white my-2 text-sm whitespace-pre-wrap">
                    <p v-if="post.post_texto" class="my-2">{{ post.post_texto }}</p>
                </div>
                <!-- post image -->
                <div v-if="post.post_imagem" class="rounded-lg overflow-hidden w-full max-w-[720px] max-h-[400px] mx-auto">
                    <img 
                        :src="post.post_imagem"
                        class="w-full h-full object-cover"
                    />
                </div>
                <!-- post iteractions -->
                <div class="flex mt-2 mb-2 gap-2">
                    <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group relative focus:outline-none" @click="like(post)">
                        <Heart32Filled v-if="post.curtiu" class="text-yellow-400" height="24"/>
                        <Heart32Regular v-else height="24" class="hover:text-yellow-400"/>
                        <TooltipContainerComponent text="Curtir" width="fit"/>
                        <p class="text-[10px]">{{ post.qtde_curtida > 0 ? interaction(post.qtde_curtida) : '&nbsp;' }}</p>
                    </button>
                    <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group relative focus:outline-none" @click="listComments(post)">
                        <!-- <Chat32Filled class="text-yellow-400" height="28"/> -->
                        <Chat32Regular height="24" class="hover:text-yellow-400"/>
                        <TooltipContainerComponent text="Ver comentários" width="fit"/>
                        <p class="text-[10px]">{{ post.qtde_comentario > 0 ? interaction(post.qtde_comentario) : '&nbsp;' }}</p>
                    </button>
                    <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group relative focus:outline-none">
                        <!-- <Chat32Filled class="text-yellow-400" height="28"/> -->
                        <Repeat height="24" class="hover:text-yellow-400"/>
                        <TooltipContainerComponent text="Repostar" width="fit"/>
                        <p class="text-[10px]">&nbsp;</p>
                    </button>
                </div>
                <!-- comments -->
                <div 
                    :class="'transition-opacity duration-1000 border-dashed border-neutral-700 rounded'
                    +(post.comentarios && post.comentarios?.length > 0 ? ' border' : '')" 
                    v-auto-animate
                >
                    <template  v-if="post.loadingComentarios">
                        <div class="flex gap-2 mb-4 mt-2 mx-2">
                            <SkeletonComponent w="30px" h="30px" rounded="full" class="mt-2" />
                            <div class="flex flex-col w-full">
                                <SkeletonComponent class="mb-2" w="80px" h="10px" />
                                <SkeletonComponent class="mb-1" w="50%" h="10px" />
                                <SkeletonComponent class="" w="150px" h="10px" />
                            </div>
                        </div>
                    </template>
                    <div v-for="comment in post.comentarios" class="flex gap-2 mb-4 mx-2 first-of-type:mt-2" :key="comment.id_comentario">
                        <img :src="comment.link_perfil" alt="Foto de perfil" class="mt-2 w-9 h-9 rounded-full bg-neutral-900 object-cover">
                        <div class="flex flex-col">
                            <p class="text-white text-sm mb-1">{{ comment.nome }}<span class="text-neutral-400 inline-block text-[12px] font-thin">{{ comment.arroba }}</span>
                                <span class="text-neutral-400 inline-block text-[12px] font-thin">
                                <div class="ml-1 mb-[2px] inline-block w-[3px] h-[3px] bg-neutral-400 rounded-full"></div> {{ timeAgo(comment.dt_comentario) }}</span>
                            </p>
                            <p class="text-white text-sm">{{ comment.comentario }}</p>
                        </div>
                    </div>
                </div>
                <!-- write your comment -->
                <form @submit.prevent="handleNewComment(post)" class="flex gap-4 items-center mt-4">
                    <img src="@/assets/user_default_foto.jpeg" alt="Foto de perfil" class="w-8 h-8 md:w-10 md:h-10 rounded-full bg-neutral-900 object-cover">
                    <input v-model="newComment" type="text" class="bg-neutral-700 placeholder:text-neutral-500 text-white rounded-lg p-2 w-full" placeholder="Adicione um comentário...">
                </form>
            </div>
        </div>

        <div v-if="postsLoading" class="w-full h-[300px] max-w-3xl mx-auto bg-neutral-800 p-4 rounded-lg flex flex-col">
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
        </div>
    </div>
</template>
