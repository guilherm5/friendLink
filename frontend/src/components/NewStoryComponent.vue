<script setup lang="ts">
import ButtonComponent from './ButtonComponent.vue';
import { Close, Image as ImageIonIcon } from '@vicons/ionicons5';
import { onMounted, ref } from 'vue';
import { imageUpload, type ImageUploadReturn } from '@/utils/imageUpload';
import * as fabric from 'fabric'
import { TextAddT24Filled } from '@vicons/fluent';
import { FormatColorFillFilled, FormatColorTextRound } from '@vicons/material';
import type { DefaultResponse } from '@/types/ApiService';
import { useNotificationStore, useAuthStore } from '../stores/global';
import { createStory } from '@/services/StoryService';
import type { StoryResponse } from '@/types/StoryService';
const emit = defineEmits(['setStage'])
const authStore = useAuthStore()
const notificationStore = useNotificationStore()
const formLoading = ref(false)
var ImageInstance = null as fabric.Image | null
var canvas = null as fabric.Canvas | null
var canvasW = 0
var canvasH = 0
var canvasBgColor = '#8c8c8c'
const textColorPicker = ref<HTMLLabelElement | null>(null)
const downloadLink = ref<HTMLAnchorElement | null>(null)
onMounted(() => {
    setWidthAndHeight()
    canvas = new fabric.Canvas('imgCanvas', {
        preserveObjectStacking: true,
        backgroundColor: canvasBgColor,
        width: canvasW,
        height: canvasH,
        enableRetinaScaling: true,
    })
    addTextToCanvas('Novo texto')
})
const setWidthAndHeight = () => {
    const w = window.innerWidth
    const h = window.innerHeight

    if(w > 400){
        canvasW = 360
    }else if(w > 300){
        canvasW = w - 40
    }else{
        canvasW = w -40
    }
    if(h > 700){
        canvasH = 640
    }else if(h > 580){
        canvasH = 500
    }else if(h > 500){
        canvasH = 480
    }else{
        canvasH = h - 70
    }
}
const handleCanvasBgChange = (e: Event) => {
    console.log('changed')
    const target = e.target as HTMLInputElement
    canvasBgColor = target.value
    canvas?.set({backgroundColor: canvasBgColor})
    canvas?.requestRenderAll()
}
const handleTextColorChange = (e: Event) => {
    const target = e.target as HTMLInputElement
    const selectedObjects = canvas?.getActiveObjects()
    if(!selectedObjects){return}
    selectedObjects.forEach(obj => {
        if(obj instanceof fabric.Textbox){
            obj.set({fill: target.value})
        }
    })
    canvas?.requestRenderAll()
}
const addTextToCanvas = (text: string) => {
    const textbox = new fabric.Textbox(text, {
        left: 50,
        top: 300,
        width: 200,
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
    textbox.on('editing:exited', function() {
        if(!textbox.text){canvas?.remove(textbox)}
    });
    textbox.on('selected', () => {
        textColorPicker.value!.classList.toggle('hidden')
    });
    textbox.on('deselected', () => {
        textColorPicker.value!.classList.toggle('hidden')
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
                scaleX: .30,
                scaleY: .30,
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
const handleSave = async () =>{
    formLoading.value = true
    if(!canvas){return}
    const selectedObjects = canvas.getActiveObjects()
    selectedObjects.forEach(obj => {
        obj.set('active', false)
        obj.setCoords();
    });
    canvas.discardActiveObject()
    canvas.overlayImage = null
    canvas.set({ clipPath: null })
    canvas.renderAll()
    const imgData = canvas.toDataURL({
        format:'jpeg', 
        quality: 1,
        multiplier: 3,
    })
    const blob = dataURLtoBlob(imgData)
    // const objurl = URL.createObjectURL(blob)
    // downloadLink.value!.href = objurl
    // downloadLink.value!.download = 'story.png'
    // downloadLink.value!.click()

    const res = await createStory(authStore.auth?.token, blob) as DefaultResponse
    if (res.status) {
        notificationStore.addNotification({
            type: 'success',
            body: 'Você fez um novo story e ele já está disponível para todos.',
        })
        const story = res.data as StoryResponse
        console.log('story adicionado', story)
        // storyStore.addStories([story], true)
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
    <div class="fixed w-full h-screen top-0 left-0 right-0 bg-neutral-900 z-50">
        <div class="h-full flex flex-col">
            <div class="flex justify-center">
                <div class="flex-1 flex justify-center max-w-md max-auto">
                    <button class="md:hidden text-yellow-400 mr-8" @click="handleSave">Salvar</button>
                    <input type="color" id="colorPicker" class="opacity-0 pointer-events-none" @input="handleCanvasBgChange" width="0" height="0" />
                    <input type="color" id="textColorPicker" class="opacity-0 pointer-events-none" @input="handleTextColorChange" width="0" height="0" />
                    <label for="imageUpload" role="button" class="mr-2 py-2 rounded-full bg-neutral-900 text-yellow-400 relative">
                        <ImageIonIcon width="24"/>
                    </label>
                    <button @click="addTextToCanvas('Novo texto')" class="mr-2 py-2 rounded-full bg-neutral-900 text-yellow-400 relative no-outline">
                        <TextAddT24Filled width="24"/>
                    </button>
                    <label for="colorPicker" role="button" class="mr-2 py-2 rounded-full text-yellow-400 relative group">
                        <FormatColorFillFilled width="24" />
                    </label>
                    <label ref="textColorPicker" for="textColorPicker" role="button" class="hidden mr-2 py-2 rounded-full text-yellow-400 relative">
                        <FormatColorTextRound width="24" />
                    </label>
                </div>
                <button class="group flex p-2" @click="$emit('setStage', 'button')">
                    <Close class="w-6 text-neutral-700 group-focus:text-yellow-400 group-hover:text-yellow-400"/>
                </button>
            </div>

            <div class="p-2 flex flex-col items-center w-full h-full">
                <canvas class="rounded" id="imgCanvas" :width="canvasW" :height="canvasH"></canvas>
                <a ref="downloadLink" class="hidden" href=""></a>
                
                <div class="hidden p-2 mt-4 md:block w-full max-w-md">
                    <ButtonComponent title="Pronto" @click="handleSave" class="ml-auto" :loading="formLoading" />
                </div>
            </div>

            <input type="file" id="imageUpload"  class="hidden" accept="image/png, image/jpeg" @change="imageUpload($event, handleImageUpload)" />
        </div>
    </div>
</template>