<script setup lang="ts">
</script>

<template>
  <div class="p-8 flex justify-around font-grotesk text-indigo-600 border-indigo-600">
    <div class="flex gap-x-2">
      <img src="/logo.svg" class="h-6 my-auto"/>
      <h1 class="text-2xl font-bold">Classee</h1>
    </div>
    <div class="flex gap-x-8 my-auto">
      <RouterLink to="/professores">Professores</RouterLink>
      <RouterLink to="/alunos">Alunos</RouterLink>
      <RouterLink to="/materias">Materias</RouterLink>
      <RouterLink to="/aulas">Aulas</RouterLink>
    </div>
  </div>
</template>