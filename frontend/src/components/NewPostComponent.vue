<script setup lang="ts">
import ButtonComponent from '@/components/ButtonComponent.vue';
import { Add12Filled } from '@vicons/fluent';
import { onMounted, ref } from 'vue';

const mounted = ref(false);
onMounted(() => {
    mounted.value = true;
});
</script>

<template>
    <div class="w-full px-2 md:px-0" v-auto-animate>
        <div class="mt-4 w-full max-w-3xl mx-auto bg-neutral-800 p-4 rounded-lg" v-auto-animate>
            <p class="text-center text-neutral-700">--botão/componente de postagem aqui--</p>
            <!-- <ButtonComponent v-if="mounted" class="mx-auto rounded-sm gap-1 px-6" title="Criar" iconSide="left">
                <template #icon>
                    <Add12Filled height="20"/>
                </template>
            </ButtonComponent> -->
        </div>
    </div>
</template>
