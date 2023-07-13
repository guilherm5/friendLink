<script setup lang="ts">
defineProps<{
    w: string,
    h: string,
}>()
</script>

<template>
    <div class="animate-pulse bg-neutral-900 rounded-lg" :style="`width:${w};height:${h};`"></div>
</template>