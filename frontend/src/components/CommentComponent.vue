<script setup lang="ts">
import { Heart32Filled, Heart32Regular } from '@vicons/fluent'
import TooltipContainerComponent from '@/components/TooltipContainerComponent.vue'
import SkeletonComponent from '@/components/SkeletonComponent.vue'
import { likeComment, unlikeComment, createReply, getReplies, unlikeReply, likeReply, createComment } from '@/services/PostService';
import type { Post, Comment, ReplyResponse } from '@/types/PostService';
import { useAuthStore, useNotificationStore } from '@/stores/global';
import { ref } from 'vue';
import type { DefaultResponse } from '@/types/ApiService';

defineProps<{
    post: Post,
    listComments: (post: Post, loadMore?: boolean) => void,
    interaction: (qtde: number) => string | false,
    timeAgo: (date: string) => string,
}>()

const repliesLimit = ref<number>(5)
const replyingCommentId = ref<number | false>(false)
const authStore = useAuthStore()
const notificationStore = useNotificationStore()
const like = (comment: Comment) => {
    if(comment.curtiu_comentario){
        unlikeComment(authStore.auth?.token, comment.id_comentario)
        comment.curtiu_comentario = false
        comment.qtde_curtida_comentario--
    }else{
        likeComment(authStore.auth?.token, comment.id_comentario)
        comment.curtiu_comentario = true
        comment.qtde_curtida_comentario++
    }
}
const replyLike = (reply: ReplyResponse) => {
    if(reply.curtiu_resp_comentario){
        unlikeReply(authStore.auth?.token, reply.id_resp_comentario)
        reply.curtiu_resp_comentario = false
        reply.qtde_curtida_resp_comentario--
    }else{
        likeReply(authStore.auth?.token, reply.id_resp_comentario)
        reply.curtiu_resp_comentario = true
        reply.qtde_curtida_resp_comentario++
    }
}
const listReplies = async (comment: Comment, loadMore: boolean = false) => {
    if(!loadMore){
        if(comment.carregadoUmaVez){
            comment.respostas = []
            comment.lastReplyId = 0
            comment.carregadoUmaVez = false
            return
        }
    }
    if(comment.qtde_resposta_comentario === 0 || comment.loadingRespostas){return}
    comment.loadingRespostas = true
    const res = await getReplies(authStore.auth?.token, comment.id_comentario, comment.lastReplyId ?? 0, repliesLimit.value) as DefaultResponse
    if (res.status && res.data) {
        const replies = res.data as ReplyResponse[]
        comment.lastReplyId = replies[replies.length - 1].id_resp_comentario
        comment.respostas ? comment.respostas.push(...res.data) : comment.respostas = res.data
        comment.carregadoUmaVez = true
    }else{
        notificationStore.addNotification({
            type: 'error',
            body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado.',
        })
    }
    comment.loadingRespostas = false
}
const handleNewReply = async (e: Event, comment: Comment) => {
    const input = (e.target as HTMLFormElement).elements[0] as HTMLInputElement
    const newReply = input.value.trim()
    if(newReply.length === 0){return false}
    const commentId = replyingCommentId.value
    input.value = ''
    replyingCommentId.value = false
    const res = await createReply(authStore.auth?.token, comment.id_comentario, newReply) as DefaultResponse
    if (res.status) {
        comment.respostas ? comment.respostas.unshift(res.data) : comment.respostas = [res.data]
        comment.qtde_resposta_comentario++
    }else{
        notificationStore.addNotification({
            type: 'error',
            body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado.',
        })
        input.value = newReply
        replyingCommentId.value = commentId
    }   
}
const handleNewComment = async (e: Event, post: Post) => {
    const input = (e.target as HTMLFormElement).elements[0] as HTMLInputElement
    const newComment = input.value.trim()
    if(newComment.length === 0){return false}
    input.value = ''
    const res = await createComment(authStore.auth?.token, post.id_post, newComment) as DefaultResponse
    if (res.status) {
        post.comentarios ? post.comentarios.unshift(res.data) : post.comentarios = [res.data]
        post.qtde_comentario++
        input.value = ''
    }else{
        notificationStore.addNotification({
            type: 'error',
            body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado.',
        })
        input.value = newComment
    }
}
</script>

<template>
    <div 
        :class="'pt-2 overflow-hidden transition-opacity duration-1000 border-dashed border-neutral-700 rounded'
        +(post.comentarios && post.comentarios?.length > 0 ? ' border' : '')" 
        v-auto-animate
    >
        <div v-for="comment in post.comentarios" class="flex gap-2 mx-2 first-of-type:mt-2" :key="comment.id_comentario">
            <img :src="comment.link_perfil" alt="Foto de perfil" class="mt-2 w-9 h-9 rounded-full bg-neutral-900 object-cover">
            <div class="flex flex-col w-full">
                <router-link :to="{name: 'profile', params:{arroba: comment.arroba}}" class="text-white text-sm mb-1 w-fit">{{ comment.nome }}
                    <span class="text-neutral-400 inline-block text-[12px] font-thin">{{ comment.arroba }}</span>
                    <span class="text-neutral-400 inline-block text-[12px] font-thin">
                    <div class="ml-1 mb-[2px] inline-block w-[3px] h-[3px] bg-neutral-400 rounded-full"></div> {{ timeAgo(comment.dt_comentario) }}</span>
                </router-link>
                <p class="text-white text-sm flex-1">{{ comment.comentario }}</p>
                <div class="flex mt-2 gap-1 items-end">
                    <button 
                        class="text-[12px] text-neutral-400 hover:text-yellow-400 bg-neutral-700 p-1 px-2 rounded-sm" 
                        @click="replyingCommentId = replyingCommentId ? false : comment.id_comentario">
                        Responder
                    </button>
                    <button v-if="comment.qtde_resposta_comentario > 0"
                        class="text-[12px] text-neutral-400 hover:text-yellow-400 bg-neutral-700 p-1 px-2 rounded-sm"
                        @click="listReplies(comment)">
                        Ver {{ interaction(comment.qtde_resposta_comentario) }} respostas
                    </button>
                    <button class="bg-neutral-700 p-1 px-2 rounded-sm text-neutral-400 md:text-neutral-400 hover:text-yellow-400 group relative" @click="like(comment)">
                        <div class="flex items-center">
                            <Heart32Filled v-if="comment.curtiu_comentario" class="text-yellow-400" height="18"/>
                            <Heart32Regular v-else height="18" class="hover:text-yellow-400"/>
                            <p class="text-[10px]">{{ comment.qtde_curtida_comentario > 0 ? interaction(comment.qtde_curtida_comentario) : '' }}</p>
                        </div>
                        <TooltipContainerComponent text="Curtir" width="fit"/>
                    </button>
                </div>

                <!-- Replies -->
                <div 
                    :class="'my-2 overflow-hidden transition-opacity duration-1000 border-dashed border-neutral-700 rounded'
                    +((comment.respostas && comment.respostas?.length > 0) || replyingCommentId === comment.id_comentario ? ' border p-2' : 'p-0')" 
                    v-auto-animate
                >
                    <form v-if="replyingCommentId === comment.id_comentario" @submit.prevent="handleNewReply($event, comment)" class="">
                        <!-- close button with x icon -->
                        <input id="aa" type="text" name="reply" class="bg-neutral-700 placeholder:text-neutral-500 text-white rounded p-2 w-full" placeholder="Adicione uma resposta.">
                    </form>

                    <div v-for="reply in comment.respostas" class="flex gap-2 mb-4 mx-2 first-of-type:mt-2" :key="reply.id_resp_comentario">
                        <img :src="reply.link_perfil" alt="Foto de perfil" class="mt-2 w-9 h-9 rounded-full bg-neutral-900 object-cover">
                        <div class="flex flex-col w-full">
                            <router-link :to="{name: 'profile', params:{arroba: reply.arroba}}" class="text-white text-sm mb-1">{{ reply.nome }}
                                <span class="text-neutral-400 inline-block text-[12px] font-thin">{{ reply.arroba }}</span>
                                <span class="text-neutral-400 inline-block text-[12px] font-thin">
                                <div class="ml-1 mb-[2px] inline-block w-[3px] h-[3px] bg-neutral-400 rounded-full"></div> {{ timeAgo(reply.dt_resp_comentario) }}</span>
                            </router-link>
                            <p class="text-white text-sm flex-1">{{ reply.resposta }}</p>
                            <div class="flex mt-2 gap-1 items-end"> 
                                <button class="bg-neutral-700 p-1 px-2 rounded-sm text-neutral-400 md:text-neutral-400 hover:text-yellow-400 group relative" @click="replyLike(reply)">
                                    <div class="flex items-center">
                                        <Heart32Filled v-if="reply.curtiu_resp_comentario" class="text-yellow-400" height="18"/>
                                        <Heart32Regular v-else height="18" class="hover:text-yellow-400"/>
                                        <p class="text-[10px]">{{ reply.qtde_curtida_resp_comentario > 0 ? interaction(reply.qtde_curtida_resp_comentario) : '' }}</p>
                                    </div>
                                    <TooltipContainerComponent text="Curtir" width="fit"/>
                                </button>
                            </div>
                        </div>
                    </div>
                    <template v-if="comment.loadingRespostas">
                        <div class="flex gap-2 mb-4 mt-2 mx-2">
                            <SkeletonComponent w="38px" h="35px" rounded="full" class="mt-2" />
                            <div class="flex flex-col w-full">
                                <SkeletonComponent class="mb-2" w="80px" h="10px" />
                                <SkeletonComponent class="mb-1" w="50%" h="10px" />
                                <SkeletonComponent class="mb-1" w="150px" h="10px" />
                                <div clas="flex">
                                    <SkeletonComponent class="inline-block" w="50px" h="15px" />
                                    <SkeletonComponent class="inline-block ml-1" w="50px" h="15px" />
                                    <SkeletonComponent class="inline-block ml-1" w="30px" h="15px" />
                                </div>
                            </div>
                        </div>
                    </template>
                    <button 
                        v-if="!comment.loadingRespostas && (comment.respostas && comment.respostas.length < comment.qtde_resposta_comentario && comment.respostas.length > 0)" 
                        @click="listReplies(comment, true)"
                        class="text-white text-sm pl-8 pb-2 no-outline"
                    >Ver mais repostas...</button>
                </div>
            </div>
        </div>
        <template  v-if="post.loadingComentarios">
            <div class="flex gap-2 mb-4 mt-2 mx-2">
                <SkeletonComponent w="38px" h="35px" rounded="full" class="mt-2" />
                <div class="flex flex-col w-full">
                    <SkeletonComponent class="mb-2" w="80px" h="10px" />
                    <SkeletonComponent class="mb-1" w="50%" h="10px" />
                    <SkeletonComponent class="mb-1" w="150px" h="10px" />
                    <div clas="flex">
                        <SkeletonComponent class="inline-block" w="50px" h="15px" />
                        <SkeletonComponent class="inline-block ml-1" w="50px" h="15px" />
                        <SkeletonComponent class="inline-block ml-1" w="30px" h="15px" />
                    </div>
                </div>
            </div>
        </template>
        <button 
            v-if="!post.loadingComentarios && (post.comentarios && post.comentarios.length < post.qtde_comentario && post.comentarios.length > 0)" 
            @click="listComments(post, true)"
            class="text-white text-sm pl-8 pb-2 no-outline"
        >Ver mais comentários...</button>
    </div>

    <form @submit.prevent="handleNewComment($event, post)" class="flex gap-4 items-center mt-4">
        <img src="@/assets/user_default_foto.jpeg" alt="Foto de perfil" class="w-8 h-8 md:w-10 md:h-10 rounded-full bg-neutral-900 object-cover">
        <input type="text" name="comment" class="bg-neutral-700 placeholder:text-neutral-500 text-white rounded-lg p-2 w-full" placeholder="Adicione um comentário.">
    </form>
</template>