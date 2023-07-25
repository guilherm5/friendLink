<script setup lang="ts">
import { ChevronCircleLeft32Filled, ChevronCircleRight32Filled } from '@vicons/fluent';
import { onMounted, ref } from 'vue';

const stories = ref<boolean[]>([false, false, false, false, false, false, false, false, false, false, false, false])
const showScrollLeft = ref<boolean>(false)
const showScrollRight = ref<boolean>(false)

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
const handleShowStory = (index: number) => {
    stories.value[index] = true
}
</script>

<template>
    <div class="relative w-full md:w-auto"  v-auto-animate>
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
                    <div v-for="(story, index) in stories" @click="handleShowStory(index)" class="max-w-[80px] snap-center cursor-pointer" :key="index">
                        <img 
                            :class="(story ? 'border-neutral-700' : 'border-yellow-400') + ' w-[80px] max-w-[80px] h-[80px] object-cover rounded-3xl p-1 bg-neutral-800 border-2'"
                            src="https://media.tenor.com/ZC0QOFGb9AgAAAAC/robin.gif" />
                        <p :class="(story ? 'text-neutral-400' : 'text-white') + ' mt-2 text-xs font-bold text-center truncate'">rangel_pci</p>
                    </div>
                    <div class="w-8 opacity-0 md:hidden">.</div>
                </div>
            </div>
        </div>
    </div>
</template>
