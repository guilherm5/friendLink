<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useNotificationStore, useAuthStore } from '../stores/global';
import type { DefaultResponse } from '@/types/ApiService';
import router from '@/router';
import ButtonComponent from '@/components/ButtonComponent.vue';
import { ArrowForward, Brush } from '@vicons/ionicons5';
import { handleInfoUpdate } from '@/services/UserService';
import type { FirstStepInfoUpdate } from '@/types/UserService';
import { imageUpload } from '@/utils/imageUpload';
import type { ImageUploadReturn } from '@/utils/imageUpload';

const notificationStore = useNotificationStore()
const authStore = useAuthStore()
const formLoading = ref(false)
const form = reactive<FirstStepInfoUpdate>({
  foto_perfil: null,
  foto_capa: null,
  bio: '',
  arroba: authStore.auth?.Nome.split(' ').join('') + Date.now().toString().slice(0, 5),
})
const fotoUrl = ref<string | null>(null)
const fotoCapaUrl = ref<string | null>(null)
const setFotoAndCoverImageFiles = async () => {
    await fetch('/src/assets/user_default_cover.jpg')
    .then((res) => res.blob())
    .then((blob) => {
        form.foto_capa = new File([blob], 'user_default_cover.jpg', { type: 'image/jpeg' })
    })
    .catch((err) => console.error(err))
    await fetch('/src/assets/user_default_foto.jpeg')
    .then((res) => res.blob())
    .then((blob) => {
        form.foto_perfil = new File([blob], 'user_default_foto.jpeg', { type: 'image/jpeg' })
    })
    .catch((err) => console.error(err))
}

await setFotoAndCoverImageFiles()
const handleSubmit = async () => {
    formLoading.value = true
    form.arroba = form.arroba.toLowerCase().split(' ').join('').replace(/[^a-z0-9]/g, '')
    const res = await handleInfoUpdate(authStore.auth?.token, form) as DefaultResponse
    if (res.status) {
        router.push({ name: 'home' })
    }else{
        notificationStore.addNotification({
        type: 'error',
        body: res.error?.response?.data as string ?? 'Ocorreu um erro ao atualizar as informações.',
        })
    }
    formLoading.value = false
}
</script>

<template>
  <main class="bg-neutral-900 container min-h-screen flex flex-col items-center p-2 py-4">
    <p class="my-8 md:my-24 text-white">Vamos começar personalizando o seu perfil</p>
    <div class="bg-neutral-800 pb-6 rounded-lg mx-4 w-full max-w-xl relative">
        <div class="bg-neutral-700 rounded-t-lg w-full h-56 overflow-hidden relative group">
            <img v-if="fotoCapaUrl" :src="fotoCapaUrl" class="w-full h-full object-cover" alt="Foto de capa">
            <img v-else src="@/assets/user_default_cover.jpg" class="w-full h-full object-cover" alt="Foto de capa">
            <label tabindex="0" role="button" for="fotoCapa" class="flex pt-2 md:pt-0 items-start md:items-center justify-center h-full text-yellow-500 absolute top-1/2 -translate-y-1/2 left-1/2 -translate-x-1/2 w-full md:bg-neutral-800
            opacity-100 md:opacity-0 md:group-hover:opacity-90 transition-opacity cursor-pointer">
                <Brush class="w-8 bg-neutral-800 p-1 rounded-l-lg" />
                <p class="bg-neutral-800 p-1 py-[6px] md:p-1 rounded-r-lg text-sm">Escolher foto de capa</p>
            </label>
            <input type="file" id="fotoCapa" class="hidden" name="fotoCapa" accept="image/png, image/jpeg" @change="imageUpload($event, (data: ImageUploadReturn) => {if(data.error){return}; form.foto_capa = data.image; fotoCapaUrl = data.urlPreview})">
        </div>
        <div class="relative">
            <div class="bg-neutral-900 rounded-full p-2 w-36 h-36 absolute -top-16 left-1/2 -translate-x-1/2">
                <div class="relative overflow-hidden rounded-full w-full h-full group">
                    <img v-if="fotoUrl" :src="fotoUrl" alt="Foto de perfil" class="w-full h-full object-cover">
                    <img v-else src="@/assets/user_default_foto.jpeg" alt="Foto de perfil" class="w-full h-full object-cover">

                    <label for="fotoPerfil" class=" flex flex-col pt-1 md:pb-0 items-center justify-start md:items-center md:justify-center h-full text-yellow-500 absolute top-1/2 -translate-y-1/2 left-1/2 -translate-x-1/2 w-full md:bg-neutral-800
                    opacity-100 md:opacity-0 md:group-hover:opacity-90 transition-opacity cursor-pointer">
                        <Brush class="w-8 bg-neutral-800 p-1 rounded-lg" />
                        <p class="text-center bg-neutral-800 p-1 rounded-r-lg hidden md:block text-sm">Escolher foto de perfil</p>
                    </label>
                    <input type="file" id="fotoPerfil" class="hidden" name="fotoPerfil" accept="image/png, image/jpeg" @change="imageUpload($event, (data: ImageUploadReturn) => {if(data.error){return}; form.foto_perfil = data.image; fotoUrl = data.urlPreview})">
                </div>
            </div>
        </div>
        <div class="flex items-center flex-col mx-4">
            <p class="mt-24 mb-1 text-center text-white font-bold text-lg">{{ authStore.auth?.Nome }}</p>
            <div class="relative">
                <label for="arroba" class="text-gray-400 absolute top-1/2 -translate-y-1/2 left-2">@</label>
                <input id="arroba" v-model="form.arroba" type="text" class="bg-neutral-700 text-neutral-400 placeholder:text-neutral-400 rounded-lg text-sm p-2 pl-6" placeholder="seu_arroba">
            </div>
            <textarea class="mt-6 bg-neutral-700 text-neutral-400 placeholder:text-neutral-400 rounded-lg text-sm p-2 w-full max-w-md no-scrollbar" v-model="form.bio" placeholder="Sobre você..." rows="4" maxlength="200">
            </textarea>
        </div>
    </div>
    <div class="mt-4  w-full max-w-xl flex justify-end">
        <ButtonComponent :loading="formLoading" @click="handleSubmit" title="Pronto">
            <template #icon>
            <ArrowForward height="24"/>
            </template>
        </ButtonComponent>
    </div>
  </main>
</template>  