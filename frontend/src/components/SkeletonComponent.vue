<script setup lang="ts">
defineProps<{
    w: string,
    h: string,
    rounded?: 'lg' | 'full' | string,
}>()
</script>

<template>
    <div :class="'animate-pulse bg-neutral-900 rounded-' + (rounded ?? 'lg')" :style="`width:${w};height:${h};`"></div>
</template>