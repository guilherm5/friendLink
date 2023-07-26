<script setup lang="ts">
import { Chat32Regular, Heart32Filled, Heart32Regular } from '@vicons/fluent';
import { DotsVertical, Repeat } from '@vicons/tabler';
import TooltipContainerComponent from './TooltipContainerComponent.vue';
import { usePostStore, useAuthStore, useNotificationStore, useCurrentUserStore } from '../stores/global';
import SkeletonComponent from './SkeletonComponent.vue';
import { getComments, likePost, unlikePost, deletePost } from '@/services/PostService';
import type { Post, CommentResponse } from '@/types/PostService';
import type { DefaultResponse } from '@/types/ApiService';
import { ref } from 'vue';
import CommentComponent from './CommentComponent.vue';

defineProps<{
    postsLoading: boolean,
}>()

const postStore = usePostStore()
const currentUserStore = useCurrentUserStore()
const authStore = useAuthStore()
const notificationStore = useNotificationStore()
const comentsLimit = ref(5)
const lastCommentId = ref(0)

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
const listComments = async (post: Post, loadMore: boolean = false) => {
    if(!loadMore){
        if(post.carregadoUmaVez){
            post.comentarios = []
            lastCommentId.value = 0
            post.carregadoUmaVez = false
            return
        }
    }
    if(post.qtde_comentario === 0 || post.loadingComentarios){return}
    post.loadingComentarios = true
    const res = await getComments(authStore.auth?.token, post.id_post, lastCommentId.value, comentsLimit.value) as DefaultResponse
    if (res.status && res.data) {
        const comments = res.data as CommentResponse[]
        lastCommentId.value = comments[comments.length - 1].id_comentario
        const userComments = comments.filter(comment => comment.id_usuario_cmt === authStore.auth?.ID)
        const otherComments = comments.filter(comment => comment.id_usuario_cmt !== authStore.auth?.ID)
        post.comentarios ?  post.comentarios.push(...userComments, ...otherComments) : post.comentarios = [...userComments, ...otherComments]
        post.carregadoUmaVez = true
    }else{
        notificationStore.addNotification({
            type: 'error',
            body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado.',
        })
    }
    post.loadingComentarios = false
}

const timeAgo = (date: string) => {
    const now = new Date()
    const postDate = new Date(date)
    const diff = now.getTime() - postDate.getTime()
    const diffYears = Math.floor(diff / (1000 * 60 * 60 * 24 * 7 * 4 * 12))
    const diffWeeks = Math.floor(diff / (1000 * 60 * 60 * 24 * 7))
    const diffDays = Math.floor(diff / (1000 * 60 * 60 * 24))
    const diffHours = Math.floor(diff / (1000 * 60 * 60))
    const diffMinutes = Math.floor(diff / (1000 * 60))
    const diffSeconds = Math.floor(diff / (1000))
    if(diffYears > 0){
        return diffYears === 1 ? '1ano atrás' : diffYears + 'anos atrás'
    }else if(diffWeeks > 0){
        return diffWeeks === 1 ? '1s atrás' : diffWeeks + 's atrás'
    }else if(diffDays > 0){
        return diffDays === 1 ? '1d atrás' : diffDays + 'd atrás'
    }else if(diffHours > 0){
        return diffHours === 1 ? '1h atrás' : diffHours + 'h atrás'
    }else if(diffMinutes > 0){
        return diffMinutes === 1 ? '1m atrás' : diffMinutes + 'm atrás'
    }else{
        return Math.abs(diffSeconds) + 's atrás'
    }
}
const handleDeletePost = async (post: Post) => {
    postStore.removePost(post.id_post)
    const res = await deletePost(authStore.auth?.token, post.id_post) as DefaultResponse
    if (!res.status) {
        notificationStore.addNotification({
            type: 'error',
            body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado.',
        })

        postStore.addPosts([post])
    }
}
</script>

<template>
    <div class="mt-4 w-full px-2 md:px-0">
        <div v-for="post in postStore.posts" class="w-full max-w-3xl mx-auto my-4 bg-neutral-800 p-4 rounded-lg" :key="post.id_post">
            <div class="flex flex-col max-w-[640px] mx-auto">
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
                            <p class="text-yellow-400 inline-block text-[12px] font-thin">{{ timeAgo(post.dt_criacao) }}</p>
                        </div>
                    </router-link>

                    <div v-if="currentUserStore.currentUser?.id_usuario === post.id_user" class="relative ml-auto group">
                        <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100  min-w-[50px] flex justify-end no-outline" :id="'options-button'+post.id_post" aria-expanded="true" aria-haspopup="true">
                            <DotsVertical class="h-4 md:h-5"/>
                        </button>

                        <div class="opacity-0 pointer-events-none group-hover:opacity-100 transition-opacity group-hover:pointer-events-auto absolute right-0 z-10 mt-0 w-56 origin-top-right rounded-md bg-black border border-neutral-700 shadow-lg focus:outline-none" role="menu" aria-orientation="vertical" :aria-labelledby="'options-button'+post.id_post" tabindex="-1">
                            <div class="py-1 text-yellow-400 flex flex-col" role="none">
                                <button @click="handleDeletePost(post)" class="px-4 py-2 text-sm hover:bg-neutral-800" role="menuitem" tabindex="-1">Apagar</button>
                            </div>
                        </div>
                    </div>
                </div>
                <!-- post text -->
                <div class="post-text text-white my-2 text-sm whitespace-pre-wrap">
                    <div v-if="post.post_texto" class="my-2">{{ post.post_texto.replace(/\n\s*\n/g, '\n\n') }}</div>
                </div>
                <!-- post image -->
                <div v-if="post.post_imagem" class="rounded-lg overflow-hidden w-full max-w-[640px] h-full max-h-[480px] mx-auto bg-neutral-900">
                    <img 
                        :src="post.post_imagem"
                        class="w-full h-full object-cover"
                        loading="lazy"
                    />
                </div>
                <!-- post iteractions -->
                <div class="flex mt-2 mb-2 gap-2">
                    <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group relative no-outline" @click="like(post)">
                        <Heart32Filled v-if="post.curtiu" class="text-yellow-400" height="24"/>
                        <Heart32Regular v-else height="24" class="hover:text-yellow-400"/>
                        <TooltipContainerComponent text="Curtir" width="fit"/>
                        <p class="text-[10px]">{{ post.qtde_curtida > 0 ? interaction(post.qtde_curtida) : '&nbsp;' }}</p>
                    </button>
                    <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group relative no-outline" @click="listComments(post)">
                        <Chat32Regular height="24" class="hover:text-yellow-400"/>
                        <TooltipContainerComponent text="Ver comentários" width="fit"/>
                        <p class="text-[10px]">{{ post.qtde_comentario > 0 ? interaction(post.qtde_comentario) : '&nbsp;' }}</p>
                    </button>
                    <button class="text-neutral-300 md:text-neutral-400 hover:text-neutral-100 group relative no-outline">
                        <Repeat height="24" class="hover:text-yellow-400"/>
                        <TooltipContainerComponent text="Repostar" width="fit"/>
                        <p class="text-[10px]">&nbsp;</p>
                    </button>
                </div>
                <CommentComponent :post="post" :listComments="listComments" :interaction="interaction" :timeAgo="timeAgo" />
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