<script setup lang="ts">
import ButtonComponent from './ButtonComponent.vue';
import { Close, Image, TrashBin } from '@vicons/ionicons5';
import { reactive, ref } from 'vue';
import { imageUpload, type ImageUploadReturn } from '@/utils/imageUpload';
import ImageCropperComponent from '@/components/ImageCropperComponent.vue';
import TooltipContainerComponent from '@/components/TooltipContainerComponent.vue';
import { useNotificationStore, useAuthStore } from '../stores/global';
import { createPost } from '@/services/PostService';
import type { DefaultResponse } from '@/types/ApiService';

const form = reactive({
    post_texto: '',
    file: null as File | null,
})
const emit = defineEmits(['setStage'])
const formLoading = ref(false)
const authStore = useAuthStore()
const notificationStore = useNotificationStore()
const editableContent = ref<HTMLElement | null>(null)
const imageToBeCropped = ref<string | null>(null)
const croppedPreview = ref<string | null>(null)
const handleImageUpload = (data: ImageUploadReturn) => {
    if(data.error){return}
    imageToBeCropped.value = data.urlPreview
}
const submit = async () => {
    formLoading.value = true
    form.post_texto = editableContent.value?.textContent ?? ''
    if(!form.post_texto && !form.file){
        notificationStore.addNotification({
            type: 'error',
            body: 'Para ppublicar algo escreva um texto ou selecione uma imagem.',
        })
    }

    const res = await createPost(authStore.auth?.token, form) as DefaultResponse
    if (res.status) {
        notificationStore.addNotification({
            type: 'success',
            body: 'Você fez uma nova postagem e ela já está disponível para todos.',
        })

        editableContent.value!.textContent = ''
        form.post_texto = ''
        form.file = null
        imageToBeCropped.value = null
        croppedPreview.value = null
        emit('setStage', 'button')
    }else{
        notificationStore.addNotification({
            type: 'error',
            body: res.error?.response?.data as string ?? 'Ocorreu um erro inesperado.',
        })
    }
    formLoading.value = false
}
</script>

<template>
    <div class="post-container mx-auto">
        <button class="block ml-auto p-2 pb-0 group" @click="$emit('setStage', 'button')">
            <Close class="w-6 text-neutral-700 group-focus:text-yellow-400 group-hover:text-yellow-400"/>
        </button>
        
        <div class="p-2">
            <div 
            ref="editableContent" 
            role="textbox" 
            contenteditable 
            class="text-area-scrollbar overscroll-contain contenteditable overflow-y-scroll min-h-[100px] max-h-80 outline-yellow-400 bg-neutral-700 placeholder:text-neutral-500 text-white rounded-lg p-2 w-full" 
            placeholder="Compartilhe suas ideias"></div>
            
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
                    <TooltipContainerComponent text="Adicionar foto" width="fit" side="right"/>
                </label>
                <ButtonComponent :loading="formLoading" @click="submit" title="Publicar" />
            </div>
        </div>

        <input type="file" id="imageUpload"  class="hidden" accept="image/png, image/jpeg" @change="imageUpload($event, handleImageUpload)" />
        <ImageCropperComponent 
            :stencilWidth="640"
            :stencilHeight="480"
            :imageUrl="imageToBeCropped"
            :croppedFile="form.file"
            @updateImageUrl="(newValue: null) => imageToBeCropped = newValue"
            @updateCroppedFile="(file: File, urlPreview) => { form.file = file; croppedPreview = urlPreview }"
        />
    </div>
</template>