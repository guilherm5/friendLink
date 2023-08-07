<script setup lang="ts">
import { Add12Filled, BorderRight24Filled } from '@vicons/fluent';
import { AccountTreeRound, AdsClickFilled } from '@vicons/material';
import { Close } from '@vicons/ionicons5';
import { onMounted, ref } from 'vue';
import NewPostComponent from '@/components/NewPostComponent.vue';
import NewStoryComponent from './NewStoryComponent.vue';

type Stage = 'button' | 'selectPostOrStory' | 'post' | 'story' | 'advertising'
const props = defineProps<{
    stage?: Stage,
}>()
const mounted = ref(false);
const currentStage = ref(props.stage ?? 'button')
onMounted(() => {
    mounted.value = true;
});

const setStage = (newStage: Stage) => {
    currentStage.value = newStage
    if(newStage === 'story'){
        document.getElementsByTagName('body')[0].style.overflowY = 'hidden'
    }else{
        document.getElementsByTagName('body')[0].style.overflowY = 'auto'
    }
}
</script>

<template>
    <div class="w-full px-2 md:px-0 flex-col justify-center" v-auto-animate>
        <div v-if="mounted" class="max-w-3xl mx-auto mt-4 bg-neutral-800 rounded-lg overflow-hidden h-fit" v-auto-animate>
            <button v-if="currentStage === 'button'" @click="setStage('selectPostOrStory')" class="transition-all w-full max-w-3xl bg-neutral-800 p-1 rounded-lg no-outline group">
                <p 
                    class="text-yellow-400 flex gap-[8px] items-center justify-center font-bold p-3 rounded-lg
                    border border-dashed border-neutral-800 group-focus:border-yellow-400 group-hover:border-yellow-400 transition-all"
                >
                    <Add12Filled height="20"/> Criar <span class="text-neutral-700  group-focus:text-yellow-400 group-hover:text-yellow-400">conteúdo</span>
                </p>
            </button>

            <div v-if="currentStage === 'selectPostOrStory'" class="transition-all bg-neutral-800 w-full max-w-3xl rounded-lg post-container mx-auto pb-2">
                <button class="block ml-auto p-2 pb-0 group" @click="setStage('button')">
                    <Close class="w-6 text-neutral-700 group-focus:text-yellow-400 group-hover:text-yellow-400"/>
                </button>
                <div class="flex h-40 md:h-60 p-1">
                    <button @click="setStage('post')" class="flex-1 bg-neutral-800 p-1 rounded-lg no-outline group">
                        <p 
                            class="h-full text-yellow-400 flex gap-[8px] items-center justify-center font-bold p-3 rounded-lg 
                            group-focus:text-yellow-400 group-hover:text-yellow-400 border border-dashed border-neutral-700 
                            group-focus:border-yellow-400 group-hover:border-yellow-400 transition-all"
                        >
                            <AccountTreeRound height="20"/><span class="text-neutral-700 group-focus:text-yellow-400 group-hover:text-yellow-400">Post</span>
                        </p>
                    </button>
                    <button @click="setStage('story')" class="flex-1 bg-neutral-800 p-1 rounded-lg no-outline group">
                        <p 
                            class="h-full text-yellow-400 flex gap-[8px] items-center justify-center font-bold p-3 rounded-lg 
                            group-focus:text-yellow-400 group-hover:text-yellow-400 border border-dashed border-neutral-700 
                            group-focus:border-yellow-400 group-hover:border-yellow-400 transition-all"
                        >
                            <BorderRight24Filled height="20"/><span class="text-neutral-700 group-focus:text-yellow-400 group-hover:text-yellow-400">Story</span>
                        </p>
                    </button>
                </div>
                <button @click="setStage('advertising')" class="w-full bg-neutral-800 p-2 pt-1 rounded-lg no-outline group">
                    <p 
                        class="text-yellow-400 flex gap-[8px] items-center justify-center font-bold p-3 rounded-lg 
                        group-focus:text-yellow-400 group-hover:text-yellow-400 border border-dashed border-neutral-700 
                        group-focus:border-yellow-400 group-hover:border-yellow-400 transition-all"
                    >
                        <AdsClickFilled height="20"/><span class="text-neutral-700 group-focus:text-yellow-400 group-hover:text-yellow-400">Anúncio</span>
                    </p>
                </button>
            </div>

            <div v-if="currentStage === 'post'" class="transition-all bg-neutral-800 w-full max-w-3xl rounded-lg">
                <NewPostComponent @setStage="setStage"/>
            </div>
            <div v-if="currentStage === 'story'" class="transition-all bg-neutral-800 w-full max-w-3xl rounded-lg">
                <NewStoryComponent @setStage="setStage"/>
            </div>
        </div>
    </div>
</template>
