<script setup lang="ts">
defineProps<{
    text: string,
    width?: 'fit' | 'full' | string,
}>()
</script>


<template>
    <div :class="'relative group mx-auto ' + (width ? 'w-'+width : 'w-full')">
        <slot name="content"></slot>
        <div role="tooltip" class="group-hover:opacity-100 z-50 transition-opacity bg-black border-2 border-neutral-800 text-gray-400 px-2 py-1 text-sm rounded 
            absolute right-2 translate-x-full top-2 -translate-y-full opacity-0 m-4 mx-auto w-max"
        >
            {{ text }}
        </div>
    </div>
</template>