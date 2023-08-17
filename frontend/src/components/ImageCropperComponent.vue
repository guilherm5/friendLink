<script setup lang="ts">
import { ref, type DefineComponent } from 'vue';
import { Cropper, type Coordinates } from 'vue-advanced-cropper';
import 'vue-advanced-cropper/dist/style.css';
import ButtonComponent from '@/components/ButtonComponent.vue';
import { CutSharp } from '@vicons/ionicons5';

defineProps<{
    stencilWidth: number,
    stencilHeight: number,
    imageUrl: string | null,
    croppedFile: File | null,
}>()

const emit = defineEmits(['updateImageUrl', 'updateCroppedFile'])
const cropper = ref<DefineComponent | null>(null)
const cropImage = () => {
    if(cropper.value){
        const data: { coordinates: Coordinates, canvas: HTMLCanvasElement } = cropper.value.getResult()
        data.canvas.toBlob((blob) => {
            const file = new File([blob as BlobPart], 'image.jpeg', { type: 'image/jpeg' })
            console.log(file)

            emit('updateCroppedFile', file, data.canvas.toDataURL('image/jpeg', 0.95))
            emit('updateImageUrl', null)
        }, 'image/jpeg', 0.95)
    };
}
</script>
<template>
    <div :class="'flex overflow-y-scroll justify-center items-center fixed top-0 left-0 right-0 bottom-0 bg-neutral-900 z-50 h-screen transition-all ' + (imageUrl ? 'opacity-100' : 'opacity-0 pointer-events-none')">
        <div class="p-4 w-full md:max-w-3xl">
            <p class="text-neutral-300 text-center pt-4 pb-2">Utilizamos um aspecto de imagem <a href="https://www.adobe.com/br/express/discover/sizes/photo-aspect-ratio" target="_blank" class="underline">4:3</a>.</p>
            <div class="bg-black rounded-t-lg pt-4 px-4  max-h-[70vh] overflow-y-hidden">
                <Cropper
                    class="cursor-grab active:cursor-grabbing"
                    ref="cropper"
                    :src="imageUrl" 
                    image-restriction="stencil" 
                    :auto-zoom="true" 
                    :stencil-size="{
                        width: 640,
                        height: 480
                    }" 
                />
            </div>
            <div class="bg-black rounded-b-lg flex justify-between w-full mx-auto max-w-3xl p-4">
                <ButtonComponent title="Cancelar" appearence="border" @click="$emit('updateImageUrl', null); $emit('updateCroppedFile', null, null)" />
                <ButtonComponent @click="cropImage" class="" icon-side="right" title="Cortar">
                    <template #icon>
                        <CutSharp width="24"/>
                    </template>
                </ButtonComponent>
            </div>
        </div>
    </div>
</template>