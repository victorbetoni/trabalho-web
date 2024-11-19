<script setup lang="ts">
import { ref } from 'vue';
import { login as apiLogin } from '../../api/auth'; 
import { allFilled } from '../../util/util';
import { useRoute, useRouter } from 'vue-router';
import Loading from '../Loading.vue';

const username = ref("");
const password = ref("");
const errorMsg = ref("");
const fetching = ref(false);

const router = useRouter();

function login() {
  fetching.value = true;
  if(!allFilled(username, password)) {
    errorMsg.value = "Preencha todos os campos.";
    return
  }
  apiLogin(username.value, password.value, async (resp) => {
    fetching.value = false;
    if(resp.status != 200) {
      errorMsg.value = resp.message!;
      return;
    }
    sessionStorage.setItem("current_user", resp.body!)
    await router.push("/professores");
    router.go(0);
  })
  }

</script>

<template>
  <div class="h-[80vh] flex flex-col justify-center">
    <div class="shadow-lg shadow-indigo-300 border-b-[6px] mx-auto gap-y-4 flex flex-col my-auto w-1/4 rounded-lg p-6 border-indigo-500 border-[1px]">
      <div class="mb-8">
        <div class="flex gap-x-4 justify-center">
          <img src="/logo.svg" class="w-8"/>
          <h1 class="text-4xl font-grotesk text-indigo-600 font-bold">Classee</h1>
        </div>
        <p class="text-center font-thin font-grotesk mt-2">Bem vindo de volta!</p>
      </div>
      <div>
        <label>Username</label>
        <input v-model="username" class="w-full" type="text"/>
      </div>
      <div>
        <label>Senha</label>
        <input v-model="password" class="w-full" type="password" />
      </div>
      <p v-if="errorMsg != ''" class="text-xs text-red-400 font-grotesk">{{ errorMsg }}</p>
      <div class="flex gap-x-4">
        <button :disabled="fetching" @mousedown="login" class="w-32">Entrar</button>
        <Loading v-if="fetching"/>
      </div>
    </div>
  </div>

</template>