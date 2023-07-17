<script setup lang="ts">
defineProps<{
    text: string,
    width?: 'fit' | 'full' | string,
    side?: 'left' | 'right',
}>()
</script>


<template>
    <div :class="'relative mx-auto ' + (width ? 'w-'+width : 'w-full')">
        <div role="tooltip" :class="`pointer-events-none hidden md:group-hover:block group-hover:opacity-100 z-50 transition-opacity bg-black border border-neutral-800 text-yellow-500 px-2 py-1 text-sm rounded 
            absolute -top-0 group-hover:-top-4 -translate-y-full opacity-0 m-4 mx-auto w-max ` + (side === 'left' ? ' -left-1 -translate-x-full' : ' -right-1 translate-x-full')"
        >
            {{ text }}
        </div>
    </div>
</template>