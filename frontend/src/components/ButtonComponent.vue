<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
    appearence?: 'border' | 'fill',
    loading?: Boolean,
    title?: String,
    iconSide?: 'left' | 'right',
}>()

const getAppearence = computed(() => {
    const fill = 'bg-yellow-500 hover:bg-yellow-600 active:bg-yellow-700 text-black'
    const border = 'border-2 border-yellow-500 hover:border-yellow-600 active:border-yellow-700 bg-transparent text-yellow-500'
    if(props.appearence === 'border'){
        return border
    }else if(props.appearence === 'fill'){
        return fill
    }else{
        return fill
    }
})
</script>

<template>
    <button v-if="loading" :class="`flex justify-center items-center text-black rounded-lg p-2 px-4 font-medium min-h-[40px] ` + getAppearence">
        <svg class="animate-spin h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
    </button>
    <button v-else :class="`transition-all text-black rounded-lg p-2 px-4 font-medium flex flex-row gap-4 items-center min-h-[40px] disabled:opacity-50 ` + getAppearence">
        <div v-if="$slots.icon && iconSide === 'left'" :class="!title ? 'mx-auto' : ''">
            <slot name="icon"></slot>
        </div>
        <div v-if="title" class="text-center flex-1 font-bold">
            {{ title }}
        </div>
        <slot v-else></slot>
        <div v-if="$slots.icon && iconSide === 'right'" :class="!title ? 'mx-auto' : ''">
            <slot name="icon"></slot>
        </div>
    </button>
</template>