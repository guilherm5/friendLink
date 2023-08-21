<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue';
import { imageUpload } from '@/utils/imageUpload';
import { useAuthStore, useNotificationStore, usePostStore } from '../stores/global';
import type { ImageUploadReturn } from '@/utils/imageUpload';
import type { FirstStepInfoUpdate } from '@/types/UserService';
import { getUser } from '@/services/UserService';
import type { DefaultResponse } from '@/types/ApiService';

const fotoCapaUrl = ref<string | null>(null)
const authStore = useAuthStore()
const form = reactive<FirstStepInfoUpdate>({
  foto_perfil: null,
  foto_capa: null,
  bio: '',
  arroba: authStore.auth?.Nome.split(' ').join('') + Date.now().toString().slice(0, 5),
})
const fotoUrl = ref<string | null>(null)
const notificationStore = useNotificationStore()
let bioUser = ref<string | null>(null)
let bioArroba = ref<string | null>(null)
let bioFotoDeCapa = ref<string | null>(null)
let bioFotoDePerfil = ref<string | null>(null)
const postStore = usePostStore()

onMounted(async () => {
  const infosUser = await getUser(authStore.auth?.token) as DefaultResponse
  if (!infosUser?.status) {
    notificationStore.addNotification({
      type: 'error',
      body: infosUser?.error?.response?.data as string ?? 'Ocorreu um erro ao carregar as informações.',
    })
  }     

  bioUser.value = infosUser.data?.bio
  bioArroba.value = infosUser.data?.arroba
  bioFotoDeCapa.value = infosUser.data?.link_capa
  bioFotoDePerfil.value = infosUser.data?.link_perfil

})

</script>

<template>
 <main class="bg-neutral-900 container min-h-screen flex flex-col items-center p-2 py-4">
    <div class="bg-neutral-800 pb-6 rounded-lg mx-4 max-w-5xl w-full relative">
        <div class="bg-neutral-700 rounded-t-lg w-full h-80 overflow-hidden relative group">
            <img v-if="bioFotoDeCapa" :src="bioFotoDeCapa" class="w-full h-full object-cover" alt="Foto de capa">
            <label tabindex="0" role="button" for="fotoCapa" class="flex pt-2 md:pt-0 items-start md:items-center justify-center h-full text-yellow-500 absolute top-1/2 -translate-y-1/2 left-1/2 -translate-x-1/2 w-full md:bg-neutral-800
              opacity-100 md:opacity-0 md:group-hover:opacity-90 transition-opacity cursor-pointer">
                <Brush class="w-8 bg-neutral-800 p-1 rounded-l-lg" />
                <p class="bg-neutral-800 p-1 py-[6px] md:p-1 rounded-r-lg text-sm">Escolher foto de capa</p>
            </label>
            <input type="file" id="fotoCapa" class="hidden" name="fotoCapa" accept="image/png, image/jpeg" @change="imageUpload($event, (data: ImageUploadReturn) => {if(data.error){return}; form.foto_capa = data.image; fotoCapaUrl = data.urlPreview})">
        </div>
        <div class="relative">
            <div class="bg-neutral-900 rounded-full w-52 h-52 p-2 absolute -top-16 left-1/2 -translate-x-1/2">
                <div class="relative overflow-hidden rounded-full w-full h-full group">
                    <img v-if="bioFotoDePerfil" :src="bioFotoDePerfil" alt="Foto de perfil" class="w-full h-full object-cover">
                    <label for="fotoPerfil" class=" flex flex-col pt-1 md:pb-0 items-center justify-start md:items-center md:justify-center h-full text-yellow-500 absolute top-1/2 -translate-y-1/2 left-1/2 -translate-x-1/2 w-full md:bg-neutral-800
                    opacity-100 md:opacity-0 md:group-hover:opacity-90 transition-opacity cursor-pointer">
                        <Brush class="w-8 bg-neutral-800 p-1 rounded-lg" />
                        <p class="text-center bg-neutral-800 p-1 rounded-r-lg hidden md:block text-sm">Escolher foto de perfil</p>
                    </label>
                    <input type="file" id="fotoPerfil" class="hidden" name="fotoPerfil" accept="image/png, image/jpeg" @change="imageUpload($event, (data: ImageUploadReturn) => {if(data.error){return}; form.foto_perfil = data.image; fotoUrl = data.urlPreview})">
                </div>
            </div>
        </div>
        <div class="flex w-full pt-40 pl-56">
          <div class="flex m-10 gap-2">
            <p class="text-white font-sans font-medium inline-block">0</p>
            <p class="text-neutral-400 text-lg font-thin md:mb-1">publicações</p>
          </div>
          <div class="flex m-10 gap-2">
            <p class="text-white font-sans font-medium inline-block">795</p>
            <p class="text-neutral-400 text-lg font-thin md:mb-1">seguidores</p>
          </div>
          <div class="flex m-10 gap-2">
            <p class="text-white font-sans font-medium inline-block">4876</p>
            <p class="text-neutral-400 text-lg font-thin md:mb-1">seguindo</p>
          </div>
        </div>
        <div class="pl-64">
          <p class="mb-1 text-left text-white font-bold text-lg">{{ authStore.auth?.Nome }}</p>
          <p class="text-neutral-400 text-left">@{{ form?.arroba }}</p>
          <p class="text-left text-white w-full">{{ bioUser }}</p>
        </div>
       <!--  <hr class="h-px my-14 bg-gray-200 border-0 dark:bg-gray-300"> -->

        <div class="grid grid-cols-3 w-full gap-2 p-2 mt-10">
          <div v-for="post in postStore.posts" class="bg-neutral-800 rounded-lg" :key="post.id_post">
            <!-- post image -->
            <div v-if="post.post_imagem && post.arroba === bioArroba" class="rounded-lg overflow-hidden w-full  h-full max-h-[247px] mx-auto bg-neutral-900">
                <img 
                    :src="post.post_imagem"
                    class="w-full h-full object-cover"
                    loading="lazy"
                />
            </div>
          </div>
        </div>

    </div>
  </main>
</template>  