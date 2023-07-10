<script setup lang="ts">
import NotificationComponent from './components/NotificationComponent.vue';
</script>

<template>
  <NotificationComponent />
  <Suspense>
    <RouterView />
  </Suspense>
</template>