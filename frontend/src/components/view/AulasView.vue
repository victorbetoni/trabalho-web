<template>
  <ProtectedContent>
    <div class="w-full p-4 md:p-0 md:w-4/5 md:mx-auto overflow-x-hidden">
      <div class="p-2 border-indigo-50 border-2 rounded-md">
        <div class="p-4 md:pl-8">
          <h1 class="text-indigo-500 font-grotesk font-bold">Aulas</h1>
        </div>
        <div class="flex gap-x-2 md:px-8">
          <div :class="[currentIndex == 0 ? 'header-active' : 'header-inactive']" @mousedown="changeTab(0)">
            <h6>Listar</h6>
          </div>
          <div :class="[currentIndex == 1 ? 'header-active' : 'header-inactive']" @mousedown="changeTab(1)">
            <h6>Cadastrar</h6>
          </div>
          <div :class="[currentIndex == 2 ? 'header-active' : 'header-inactive']" @mousedown="changeTab(2)">
            <h6>Presenças</h6>
          </div>
        </div>
        <div class="p-2 md:p-8">
          <NewAula v-if="currentIndex==1"/>
          <ListAulas v-if="currentIndex==0"/>
        </div>
      </div>
    </div>
  </ProtectedContent>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import ProtectedContent from './ProtectedContent.vue';
import NewAula from './tab/aulas/NewAula.vue';
import ListAulas from './tab/aulas/ListAulas.vue';

const currentIndex = ref(0);

function changeTab(i: number) {
  currentIndex.value = i
}

</script>
