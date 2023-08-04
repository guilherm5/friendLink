<script setup lang="ts">
import ButtonComponent from './ButtonComponent.vue';
import { Close, Image as ImageIonIcon } from '@vicons/ionicons5';
import { onMounted, ref } from 'vue';
import { imageUpload, type ImageUploadReturn } from '@/utils/imageUpload';
// import TooltipContainerComponent from '@/components/TooltipContainerComponent.vue';
import * as fabric from 'fabric'
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
// canvas for stories like instagram
var ImageInstance = null as fabric.Image | null
var canvas = null as fabric.Canvas | null
const downloadLink = ref<HTMLAnchorElement | null>(null)
onMounted(() => {
    canvas = new fabric.Canvas('imgCanvas', {
        preserveObjectStacking: true,
        controlsAboveOverlay: true,
        backgroundColor: 'rgb(140,140,140)',
        // clipTo: (ctx: CanvasRenderingContext2D) => {
        //     ctx.rect(220, 80, 360, 640);
        // },
    })
    fabric.Image.fromURL('https://www.acwe.co.uk/tutorials/story-maker-tutorial-part-1/img/overlay-bg.png').then((overlayImg: fabric.Image) => {
        overlayImg.set({
            opacity: 0.5,
            angle: 0,
            left: 0,
            top: 0,
            originX: 'left',
            originY: 'top',
            crossOrigin: 'anonymous',
            selectable: false
        })
        canvas!.overlayImage = overlayImg
        canvas?.renderAll()
    })
    addTextToCanvas('Story text goes here - Lorem Ipsum')
})

const addTextToCanvas = (text: string) => {
    const textbox = new fabric.Textbox(text, {
        left: 20,
        top: 455,
        width: 320,
        fontSize: 24,
        fill: '#fff',
        fontFamily: 'sans-serif',
        fontWeight: 500,
        textAlign: 'center',      
        borderColor: 'green',
        cornerColor: 'green',
        cornerSize: 12,
        transparentCorners: false,
        shadow: 'rgba(0,0,0,0.3) 2px 2px 10px',
    });
    canvas?.add(textbox);
}
const handleImageUpload = (data: ImageUploadReturn) => {
    if(!data.error && data.urlPreview){   
        if(ImageInstance){
            canvas?.remove(ImageInstance)
        }
        const storyImageCover = new Image()
        storyImageCover.src = data.urlPreview
        storyImageCover.onload = () => {
            if(!storyImageCover){return}
            ImageInstance = new fabric.Image(storyImageCover, {
                angle: 0,
                left: 0,
                top: 0,
                selectable: true,      
                borderColor: '#EAB308',
                cornerColor: '#EAB308',
                cornerSize: 10,
                transparentCorners: true,
            });
            canvas?.add(ImageInstance);
            canvas?.sendObjectToBack(ImageInstance)
        }
    }
}
const dataURLtoBlob = (dataUrl: string) => {
    const arr = dataUrl.split(','), mime = arr[0].match(/:(.*?);/)?.[1]
    const bstr = atob(arr[1])
    let n = bstr.length
    const u8arr = new Uint8Array(n)
    while(n--){
        u8arr[n] = bstr.charCodeAt(n)
    }
    return new Blob([u8arr], {type:mime})
}
const handleSave = () =>{
    if(!canvas){return}
    const imgData = canvas.toDataURL({format:'png', multiplier: 3});
    const blob = dataURLtoBlob(imgData);
    const objurl = URL.createObjectURL(blob)
    downloadLink.value!.href = objurl
    downloadLink.value!.download = 'story.png'
    downloadLink.value!.click()
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
    <div class="fixed w-full h-full top-0 left-0 right-0 bg-neutral-900 z-50">
        <div>
            <button class="flex items-center text-white ml-auto p-2 pb-0 group" @click="$emit('setStage', 'button')">
                <span class="text-neutral-700 group-focus:text-yellow-400 group-hover:text-yellow-400">Fechar</span> 
                <Close class="w-6 text-neutral-700 group-focus:text-yellow-400 group-hover:text-yellow-400"/>
            </button>

            <div class="p-2">
                <canvas id="imgCanvas" width="800" height="800"></canvas>
                <a ref="downloadLink" class="hidden" href=""></a>
            </div>
            
            <div class="p-2">
                <div class="mt-2 flex justify-between">
                    <label for="imageUpload" role="button" class="p-2 rounded-full bg-neutral-900 text-yellow-400 relative group">
                        <ImageIonIcon width="24"/>
                    </label>
                    <ButtonComponent title="Pronto" @click="handleSave" />
                </div>
            </div>

            <input type="file" id="imageUpload"  class="hidden" accept="image/png, image/jpeg" @change="imageUpload($event, handleImageUpload)" />
        </div>
    </div>
</template>