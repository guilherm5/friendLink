<script setup lang="ts">
import { reactive, ref } from 'vue';
import { handleSignup, handleSignin } from '@/services/AuthService';
import { useNotificationStore, useAuthStore } from '../stores/global';
import type { LoginResponse } from '@/types/AuthService';
import router from '@/router';
import ButtonComponent from '@/components/ButtonComponent.vue';
import { EnterOutline } from '@vicons/ionicons5';
import type { DefaultResponse } from '@/types/ApiService';

const notificationStore = useNotificationStore()
const authStore = useAuthStore()
const formLoading = ref(false)
const form = reactive({
  nome: '',
  email: '',
  senha: '',
})

const handleSubmit = async () => {
  formLoading.value = true
  const res = await handleSignup(form) as DefaultResponse
  if (res.status) {
    const res = await handleSignin({email: form.email, senha: form.senha}) as LoginResponse
    authStore.setAuth(res.data.logged)
    router.push({ name: 'first-steps' })
  }else{
    notificationStore.addNotification({
      type: 'error',
      body: res.error?.response?.data as string ?? 'Ocorreu um erro interno, entre em contato com o suporte.',
    })
  }
  formLoading.value = false
}
</script>

<template>
  <main class="bg-neutral-900 container min-h-screen">

    <div class="w-full px-2 max-w-[400px] fixed top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2">
      <div class="bg-neutral-800 p-8 rounded-lg">
        <img src="@/assets/logo.svg" width="42" class="mx-auto mb-8" />
        <form @submit.prevent="handleSubmit">
          <div class="flex flex-col">
            <label for="name" class="mt-3 text-white text-left">Nome</label>
            <input v-model="form.nome" type="text" id="name" class="bg-neutral-700 text-gray-400 rounded-lg p-2">

            <label for="email" class="mt-3 text-white text-left">Email</label>
            <input v-model="form.email" type="email" id="email" class="bg-neutral-700 text-gray-400 rounded-lg p-2">
            
            <label for="password" class="mt-3 text-white text-left">Senha</label>
            <input v-model="form.senha" type="password" id="password" class="bg-neutral-700 text-gray-400 rounded-lg p-2">

            <ButtonComponent :loading="formLoading" title="Continuar" class="mt-8" iconSide="right">
              <template #icon>
                <EnterOutline height="24"/>
              </template>
            </ButtonComponent>
          </div>
        </form>

        <p class="text-center mt-10 text-gray-400">Já possui uma conta? <RouterLink :to="{name: 'signin'}">entrar</RouterLink></p>
      </div>
    </div>
  </main>
</template>  