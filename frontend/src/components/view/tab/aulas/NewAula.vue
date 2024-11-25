<script setup lang="ts">
import { ref } from 'vue';
import { vMaska } from "maska/vue"
import { useToast } from 'vue-toastification';
import { findProfessor } from '../../../../api/professores';
import { allFilled } from '../../../../util/util';
import { createAula } from '../../../../api/aula';

const toast = useToast();

const ok = ref(false);
const cpf = ref("");
const nomeProfessor = ref("");
const titulo = ref("");
const data = ref("");
const hora = ref("");

function fetchProfessor() {
  findProfessor(cpf.value, (resp) => {
    if(resp.status != 200) {
      toast.error(resp.message);
      return
    }
    if(resp.body!.length < 1) {
      toast.error("Nenhum professor encontrado.");
      return
    }
    resp.body!.forEach(x => nomeProfessor.value = x.nome)
    ok.value = true;
  })
}

function submit() {
  if(!allFilled(cpf, titulo, data, hora)) {
    toast.error("Preencha todos os campos.");
    return
  }
  const dataHoraStr = `${data.value}T${hora.value}:00`;
  const dataHora = new Date(dataHoraStr);
  console.log(Math.floor(dataHora.getTime() / 1000))
  createAula(cpf.value, titulo.value, Math.floor(dataHora.getTime() / 1000), (resp) => {
    if(resp.status != 200) {
      toast.error(resp.message);
      return
    }
    toast.success("Aula criada com sucesso!");
    cancel();
  })
}

function cancel() {
  cpf.value = "";
  nomeProfessor.value = "";
  titulo.value = "";
  ok.value = false;
}

</script>

<template>
    <div class="flex flex-col gap-y-4">
      <div>
        <label>CPF (Professor)</label>
        <div class="flex gap-x-2">
          <input v-model="cpf" :disabled="ok" v-maska="'###.###.###-##'" class="w-64"/>
          <button @mousedown="fetchProfessor" :class="[ok || cpf.length != 14 ? 'bg-gray-300 hover:bg-gray-300' : '']">Buscar</button>
        </div>
      </div>
      <div>
        <label>Professor</label>
        <input disabled class="w-96" v-model="nomeProfessor"/>
      </div>
      <div>
        <label>Título</label>
        <input v-model="titulo" :disabled="!ok" class="w-96"/>
      </div>
      <div class="flex gap-x-2">
        <div>
          <label>Data</label>
          <input :disabled="!ok" v-model="data" type="date"/>
        </div>
        <div>
          <label>Horário</label>
          <input :disabled="!ok" v-model="hora" type="time"/>
        </div>
      </div>
      <div v-if="ok" class="gap-x-2 flex">
        <button @mousedown="cancel" class="w-32 bg-red-400 hover:bg-red-400">Cancelar</button>
        <button @mousedown="submit" class="w-32">Criar</button>
      </div>
    </div>
</template>