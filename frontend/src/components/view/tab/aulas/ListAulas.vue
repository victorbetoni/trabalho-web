<script setup lang="ts">
import { Ref, ref } from 'vue';
import { APIResponse } from '../../../../api/api';
import ProtectedContent from '../../ProtectedContent.vue';
import { useToast } from 'vue-toastification';
import { Aula } from '../../../../model/Aula';
import { deleteAula, fetchAulas } from '../../../../api/aula';

const aulas: Ref<Aula[]> = ref([])
const nextPage: Ref<Aula[]> = ref([])
const cpf = ref("")
const failed = ref(false)
const fetching = ref(false)

const deleting = ref("");

const page = ref(1)

const toast = useToast()

function previous() {
  page.value--;
  nextPage.value = [];
  buscar(true);
}

function next() { 
  page.value++;
  aulas.value = [...nextPage.value];
  nextPage.value = [];
  buscar(false);
}

function deletar() {
  if(deleting.value.length < 1) {
    return;
  }
  deleteAula(deleting.value, (resp) => {
    deleting.value = "";
    if(resp.status != 200) {
      toast.error(resp.message, {timeout: 5000})
      return;
    }
    toast.success("Aula excluída com sucesso!")
    page.value = 1
    buscar(true)
  });
}

function handleResp(resp: APIResponse<Aula[]>, next: boolean) {
  fetching.value = false;
    if(resp.status != 200) {
      failed.value = true;
      toast.error(resp.message);
      return
    }
    failed.value = false;
    if(next) {
      nextPage.value = resp.body! as Aula[];
      return
    }
    aulas.value = resp.body! as Aula[];
}

function buscar(fetchCurrent: boolean) {
  fetching.value = true;
  if(fetchCurrent) {
    fetchAulas(page.value, cpf.value, (resp: APIResponse<Aula[]>) => {
      handleResp(resp, false)
    })
  }

  // Busca a próxima pagina e deixa em cache.
  fetchAulas(page.value+1, cpf.value, (resp: APIResponse<Aula[]>) => {
    handleResp(resp, true)
  })
}

function formatarData(data: string) {
  const d = new Date(data)
  const dia = String(d.getDate()).padStart(2, '0');
  const mes = String(d.getMonth() + 1).padStart(2, '0');
  const ano = d.getFullYear();
  const horas = String(d.getHours()).padStart(2, '0');
  const minutos = String(d.getMinutes()).padStart(2, '0');

  return `${dia}/${mes}/${ano} ${horas}:${minutos}`;
}

</script>

<template>
  <div v-if="deleting != ''" class="fixed left-[5%] md:left-[40%] w-80 z-10 text-gray-600 top-[40%] flex flex-col gap-y-8 p-8 text-center font-grotesk bg-white border-2 shadow-lg border-b-8 rounded-lg border-yellow-200">
    <img src="/warn.svg" class="w-12 mx-auto"/>
    <p>Excluindo aula de ID: {{ deleting }}<br>Essa ação é irreversível. Tem certeza?</p>
    <div class="flex gap-x-2 mx-auto">
      <button class="bg-red-400 hover:bg-red-500" @mousedown="deleting = ''">Cancelar</button>
      <button class="bg-yellow-400 hover:bg-yellow-500" @mousedown="deletar">Confirmar</button>
    </div>
  </div>
  <ProtectedContent>
    <div>
      <header class="flex md:flex-row flex-col gap-y-2 gap-x-2">
        <div>
          <label for="CPF" >Professor</label>
          <input id="cpf" v-model="cpf" class="w-full md:w-48"/>
        </div>
        <button @mousedown="page = 1; buscar(true)" class="block">Buscar</button>
      </header>
      <div class="mt-8 w-full md:block hidden">
        <table class=" flex flex-col text-gray-600">
          <thead>
            <tr class="border-b-0">
              <th class="w-1/4">Professor</th>
              <th class="w-1/4">Título</th>
              <th class="w-1/6">Data</th>
              <th class="w-1/6">Carga Horária</th>
              <th>Ações</th>
            </tr>
          </thead>
          <tbody>
            <tr class="font-grotesk text-sm" v-for="(aula, index) in aulas" :class="[index == aulas.length - 1 ? 'border-b-0' : '']">
              <td class="w-1/4">
                {{ aula.professor.nome }}
              </td>
              <td class="w-1/4">
                {{ aula.titulo }}
              </td>
              <td class="w-1/6">
                {{ formatarData(aula.data) }} 
              </td>
              <td class="w-1/6">
                {{ aula.carga_horaria }} 
              </td>
              <td>
                <button class="bg-red-400 hover:bg-red-500" @mousedown="deleting = aula.id">Excluir</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="block md:hidden mt-8 font-grotesk text-gray-600 text-xs">
        <div v-for="aula in aulas" class="w-full border-[1px] border-indigo-500 p-4 rounded-md mt-2 flex gap-x-4 justify-between">
          <div class="flex flex-col gap-y-1">
            <p>{{ aula.professor.nome }}</p>
            <p>{{ aula.titulo }}</p>
            <p>{{ aula.data }}</p>
            <p>Carga Horária: {{ aula.carga_horaria }}</p>
          </div>
          <div @mousedown="deleting = aula.id" class="h-fit mt-auto w-fit rounded-md p-2 bg-red-400">
            <img src="/trash.svg" class="w-6">
          </div>
        </div>
      </div>
      <div class="flex justify-end gap-x-2 mt-4">
        <button v-if="page > 1" class="px-1 py-1" @mousedown="previous"><img src="/left.svg" class="w-4"/></button>
        <span class="my-auto  font-grotesk font-bold">{{ page }}</span>
        <button v-if="nextPage.length > 0" class="px-1 py-1"><img src="/right.svg" @mousedown="next" class="w-4"/></button>
      </div>
    </div>
  </ProtectedContent>
</template>