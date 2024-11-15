<template>
  <div class="w-4/5 mx-auto">
    <div class="border-indigo-50 border-2 rounded-md">
      <div class="p-4 pl-8">
        <h1 class="text-indigo-500 font-grotesk font-bold">Professores</h1>
      </div>
      <div class="flex gap-x-2 px-8">
        <div :class="[currentIndex == 0 ? 'header-active' : 'header-inactive']" @mousedown="changeTab(0)">
          <h6>Listar Professores</h6>
        </div>
        <div :class="[currentIndex == 1 ? 'header-active' : 'header-inactive']" @mousedown="changeTab(1)">
          <h6>Cadastrar Professor</h6>
        </div>
        <div :class="[currentIndex == 2 ? 'header-active' : 'header-inactive']" @mousedown="changeTab(2)">
          <h6>Editar Professor</h6>
        </div>
      </div>
      <div class="p-8">
        <ListProfessores v-if="currentIndex==0"/>
        <NewProfessor v-if="currentIndex==1"/>
        <UpdateProfessor v-if="currentIndex==2"/>
      </div>
   </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import ListProfessores from './tab/professores/ListProfessores.vue';
import NewProfessor from './tab/professores/NewProfessor.vue';
import UpdateProfessor from './tab/professores/UpdateProfessor.vue';

const currentIndex = ref(0);

function changeTab(i: number) {
  currentIndex.value = i
}

</script>
