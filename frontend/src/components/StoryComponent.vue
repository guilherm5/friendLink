<script setup lang="ts">
import { ChevronCircleLeft32Filled, ChevronCircleRight32Filled } from '@vicons/fluent';
import { onMounted, ref, reactive, watch } from 'vue';
import SkeletonComponent from './SkeletonComponent.vue';
import { useStoryStore } from '../stores/global';
import type { Story } from '@/types/StoryService';
import { Close } from '@vicons/ionicons5';
import type { StoryItem } from '@/types/StoryService';
import { timeAgo } from '@/utils/calculateTimeAgo';

defineProps<{
    storiesLoading: boolean,
}>()

type ShowingStory = {
    story: Story | null,
    index: number
    storyItem: StoryItem | null
}

const storyStore = useStoryStore()
const showScrollLeft = ref(false)
const showScrollRight = ref(false)
const showStoryModal = ref(false)
const showingStory = reactive<ShowingStory>({
    story: null,
    index: 0,
    storyItem: null
})
onMounted(() => {
    const scrollContainer = document.getElementById('storyContainer')
    scrollContainer && setScrollArrows(scrollContainer)
    scrollContainer?.addEventListener('scroll', () => {
        setScrollArrows(scrollContainer)
    });
});
const setScrollArrows = (scrollContainer: HTMLElement) => {
    const scrollLeft = scrollContainer.scrollLeft
    const scrollWidth = scrollContainer.scrollWidth
    if (scrollLeft > 0) {
        showScrollLeft.value = true
    }else {
        showScrollLeft.value = false
    }
    if (scrollLeft + scrollContainer.clientWidth < scrollWidth) {
        showScrollRight.value = true
    }else {
        showScrollRight.value = false
    }
}
const scrollHalf = (side: 'right' | 'left') => {
    const scrollContainer = document.getElementById('storyContainer')
    scrollContainer?.scrollBy({
        left: side === 'right' ? scrollContainer.clientWidth / 2 : -scrollContainer.clientWidth / 2,
        behavior: 'smooth'
    });
};


const storyTimerProgressWidth = ref<string>('0%')
const storyTimerProgress = ref<number>(0)
const storyTimer = ref<number>(0)
const storyTimerStart = (story: Story, index = 0) => {
    showingStory.story = story
    showingStory.index = index
    showingStory.storyItem = story.stories[index]
    storyTimerProgressWidth.value = '0%'
    storyTimerProgress.value = 0
    storyTimer.value = window.setInterval(() => {
        storyTimerProgress.value += 100
        storyTimerProgressWidth.value = `${(storyTimerProgress.value / 5000) * 100}%`
        if (storyTimerProgress.value >= 5000) {
            storyTimerStop()
            story.stories[index].visualizou = true
            if(index < story.stories.length - 1){
                // call request to set story as visualized

                storyTimerStart(story, index + 1)
            }else{
                showStoryModal.value = false
            }
        }
    }, 100);
}
const storyTimerStop = () => {
    window.clearInterval(storyTimer.value)
}
const handleShowStory = (story: Story) => {
    showStoryModal.value = true
    storyTimerStart(story)
}
const closeStoryModal = () => {
    showStoryModal.value = false
    storyTimerStop()
}
watch(showStoryModal, (value) => {
    if(value){
        document.body.style.overflow = 'hidden'
    }else{
        document.body.style.overflow = 'auto'
    }
})
</script>

<template>
    <div class="relative w-full md:w-auto" v-if="storyStore.stories"  v-auto-animate>
        <div class="max-w-[700px] mx-auto">
            <div v-if="showScrollLeft" class="pointer-events-none absolute top-0 left-0 w-1/3 h-full bg-gradient-to-r from-neutral-900 z-10"></div>
            <div v-if="showScrollRight" class="pointer-events-none absolute top-0 right-0 w-1/3 h-full bg-gradient-to-l from-neutral-900 z-10"></div>
            <div class="hidden absolute top-0 bottom-0 -left-3 md:flex items-center z-10">
                <ChevronCircleLeft32Filled v-if="showScrollLeft" class="mb-4 transition text-neutral-500 hover:text-neutral-300 cursor-pointer" @click="scrollHalf('left')" height="32"/>
            </div>
            <div class="hidden absolute top-0 bottom-0 -right-3 md:flex items-center z-10">
                <ChevronCircleRight32Filled v-if="showScrollRight" class="mb-4 transition text-neutral-500 hover:text-neutral-300 cursor-pointer" @click="scrollHalf('right')" height="32"/>
            </div>
            
            <div class="overflow-x-scroll w-full max-w-[768px] snap-x no-scrollbar" id="storyContainer">
                <div class="flex w-full gap-4 py-4">
                    <div class="w-8 opacity-0 md:hidden">.</div>
                    <div v-for="story in storyStore.stories" @click="handleShowStory(story)" class="max-w-[80px] snap-center cursor-pointer" :key="story.id_usuario" :title="story.nome">
                        <img 
                            :class="(story.stories.find(storyItem => !storyItem.visualizou) ? 'border-yellow-400' : 'border-neutral-700') + ' w-[80px] max-w-[80px] h-[80px] object-cover rounded-3xl p-1 bg-neutral-800 border-2'"
                            :src="story.link_perfil" alt="Foto de perfil do autor do story" />
                        <p :class="(story.stories.find(storyItem => !storyItem.visualizou) ? 'text-white' : 'text-neutral-400') + ' mt-2 text-xs font-bold text-center truncate'">
                            {{ story.arroba }}
                        </p>
                    </div>
                    <div class="w-8 opacity-0 md:hidden">.</div>
                </div>
            </div>
            <div v-if="storiesLoading" class="flex gap-4">
                <div class="max-w-[80px] bg-neutral-800 rounded-3xl">
                    <SkeletonComponent w="80px" h="80px" class="rounded-3xl"/>
                </div>
            </div>
        </div>
    </div>

    <div class="fixed w-full h-screen overflow-y-scroll top-0 left-0 right-0 bg-neutral-900 z-50" v-auto-animate v-if="showStoryModal && showingStory.story && showingStory.storyItem">
        <div class="w-full h-full flex items-center flex-col relative p-4 pt-0">
            <div class="pt-4 pb-2 md:m-2 rounded-sm">
                <div class="w-full flex gap-1 items-center">
                    <div v-for="storyItem, index in showingStory.story.stories" class="h-1 w-full rounded bg-neutral-600" :key="storyItem.id_story">
                        <div class="h-1 rounded bg-white transition-all" :style="{width: index < showingStory.index ? '100%' : index == showingStory.index ? storyTimerProgressWidth : '0%'}"></div>
                    </div>
                    <button class="p-2 pr-0" @click="closeStoryModal">
                        <Close class="w-6 text-white group-focus:text-yellow-400 group-hover:text-yellow-400"/>
                    </button>
                </div>
                <div class="flex mb-2 gap-2">
                    <img :src="showingStory.story.link_perfil" alt="Foto de perfil" class="shadow-neutral-800 shadow-md w-8 h-8 md:w-10 md:h-10 rounded-full bg-neutral-900 object-cover">
                    <div class="flex items-center">
                        <p class="text-white font-sans font-medium inline-block">{{showingStory.story.nome}}</p>
                        <div class="mx-2 inline-block w-1 h-1 bg-yellow-400 rounded-full"></div>
                        <p class="text-yellow-400 inline-block text-[12px] font-thin">{{ timeAgo(showingStory.storyItem.dt_criacao) }}</p>
                    </div>
                </div>
                <img :src="showingStory.storyItem.link_story" :title="`Story número ${showingStory.index}`" class="md:max-w-[360px] rounded-lg shadow-neutral-800 shadow-md" />
            </div>
        </div>
    </div>
</template>
