<script setup lang="ts">
import ButtonComponent from './ButtonComponent.vue';
import { Close, Image } from '@vicons/ionicons5';
import TooltipContainerComponent from './TooltipContainerComponent.vue';
import { reactive, ref } from 'vue';

const form = reactive({
    text: '',
})
const editableContent = ref<HTMLElement | null>(null)
const submit = () => {
    form.text = editableContent.value?.textContent ?? ''
}
</script>

<template>
    <div>
        <button class="block ml-auto p-2 pb-0 group" @click="$emit('setStage', 'button')">
            <Close class="w-6 text-neutral-700 group-focus:text-yellow-400 group-hover:text-yellow-400"/>
        </button>
        
        <div class="p-2">
            <div 
            ref="editableContent" 
            role="textbox" 
            contenteditable 
            class="text-area-scrollbar overscroll-contain contenteditable overflow-y-scroll max-h-80 outline-yellow-400 bg-neutral-700 placeholder:text-neutral-500 text-white rounded-lg p-2 w-full" 
            placeholder="Compartilhe suas ideias"></div>
            
            <div class="mt-2 flex justify-between">
                <button class="p-2 rounded-full bg-neutral-900 text-yellow-400 relative group">
                    <Image width="24"/> 
                    <TooltipContainerComponent text="Adicionar foto" width="fit" side="right"/>
                </button>
                <ButtonComponent @click="submit" title="Publicar" />
            </div>
        </div>
    </div>
</template>