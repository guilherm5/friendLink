<script setup lang="ts">
import { CloseCircleOutline, CheckmarkCircleOutline, Close } from '@vicons/ionicons5';
import { useNotificationStore } from '../stores/global';

const notificationStore = useNotificationStore()
</script>

<template>
    <div class="fixed pb-4 top-2 left-2 flex flex-col gap-2 w-full max-w-[350px] max-h-screen overflow-x-hidden overflow-y-auto no-scrollbar z-40" v-auto-animate>
        <div 
            v-for="notification in notificationStore.notifications" 
            :key="notification.id"
            class="bg-white p-2 mr-4 rounded flex flex-col overflow-hidden shadow min-h-[88px]"
        >
            <!-- header -->
            <div class="flex justify-between mb-2">
                <CheckmarkCircleOutline v-if="notification.type === 'success'" class="w-6 bg-green-500 text-white rounded-full"/>
                <CloseCircleOutline v-if="notification.type === 'error'" class="w-6 bg-red-500 text-white rounded-full"/>
                <button 
                    class="text-black" 
                    @click="notificationStore.removeNotification(notification.id)"
                ><Close class="w-6 text-gray-300"/></button>
            </div>
            <!-- body -->
            <div class="text-sm text-gray-600">{{ notification.body }}</div>
            <!-- footer -->
            <div class="text-sm text-gray-500">{{ notification.footer }}</div>
        </div>
    </div>
</template>