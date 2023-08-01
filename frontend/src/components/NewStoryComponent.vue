<script setup lang="ts">
import ButtonComponent from './ButtonComponent.vue';
import { Close, Image, TrashBin } from '@vicons/ionicons5';
import { ref } from 'vue';
import { imageUpload, type ImageUploadReturn } from '@/utils/imageUpload';
import TooltipContainerComponent from '@/components/TooltipContainerComponent.vue';
// import { useNotificationStore, useAuthStore, usePostStore } from '../stores/global';
// import { createPost } from '@/services/PostService';
// import type { DefaultResponse } from '@/types/ApiService';
// import type { Post } from '@/types/PostService';
// import * as fabric from 'fabric';

// const imageFile = ref<File | null>(null)
// const canvas = ref<fabric.Canvas | null>(null)
// const emit = defineEmits(['setStage'])
// const formLoading = ref(false)
// const authStore = useAuthStore()
// const postStore = usePostStore()
// const notificationStore = useNotificationStore()
// const editableContent = ref<HTMLElement | null>(null)
const imageToBeCropped = ref<string | null>(null)
const croppedPreview = ref<string | null>(null)
const handleImageUpload = (data: ImageUploadReturn) => {
    if(data.error){return}
    imageToBeCropped.value = data.urlPreview
}
// const submit = async () => {
//     formLoading.value = true
//     form.post_texto = editableContent.value?.innerText ?? ''
//     if(!form.post_texto && !form.file){
//         notificationStore.addNotification({
//             type: 'error',
//             body: 'Para publicar algo escreva um texto ou selecione uma imagem.',
//         })
//     }

//     const res = await createPost(authStore.auth?.token, form) as DefaultResponse
//     if (res.status) {
//         notificationStore.addNotification({
//             type: 'success',
//             body: 'Você fez uma nova postagem e ela já está disponível para todos.',
//         })
//         const post = res.data as Post
//         postStore.addPosts([post], true)
        
//         editableContent.value!.innerText = ''
//         form.post_texto = ''
//         form.file = null
//         imageToBeCropped.value = null
//         croppedPreview.value = null
//         emit('setStage', 'button')
//     }else{
//         notificationStore.addNotification({
//             type: 'error',
//             body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado.',
//         })
//     }
//     formLoading.value = false
// }
</script>

<template>
    <div class="post-container mx-auto">
        <button class="block ml-auto p-2 pb-0 group" @click="$emit('setStage', 'button')">
            <Close class="w-6 text-neutral-700 group-focus:text-yellow-400 group-hover:text-yellow-400"/>
        </button>

        <div class="p-2">
            <canvas width="360" height="640"></canvas>
        </div>
        
        <div class="p-2">
            <div v-if="croppedPreview" class="rounded-lg overflow-hidden w-full max-w-[640px] max-h-[480px] mx-auto mt-4 bg-neutral-900 relative">
                <img 
                    :src="croppedPreview"
                    title="Imagem a ser publicada"
                    class="w-full h-full object-cover"
                />
                <button @click="croppedPreview = null" class="absolute mt-2 mr-2 top-0 right-0 p-2 rounded-full bg-neutral-800 hover:bg-neutral-900 transition-all">
                    <TrashBin class="w-6 text-red-500 group-focus:text-yellow-400 group-hover:text-yellow-400"/>
                    <div class="relative">
                        <TooltipContainerComponent text="Apagar imagem" width="fit" side="left"/>
                    </div>
                </button>
            </div>

            <div class="mt-2 flex justify-between">
                <label for="imageUpload" role="button" class="p-2 rounded-full bg-neutral-900 text-yellow-400 relative group">
                    <Image width="24"/>
                </label>
                <ButtonComponent title="OK" />
            </div>
        </div>

        <input type="file" id="imageUpload"  class="hidden" accept="image/png, image/jpeg" @change="imageUpload($event, handleImageUpload)" />
    </div>
</template>