<script setup lang="ts">
import { Ref, ref } from 'vue';
import Professor from '../../../../model/Professor';
import { deleteProfessor, findProfessores } from '../../../../api/professores';
import { APIResponse } from '../../../../api/api';
import ProtectedContent from '../../ProtectedContent.vue';
import { useToast } from 'vue-toastification';

const professores: Ref<Professor[]> = ref([])
const nextPage: Ref<Professor[]> = ref([])
const cpf = ref("")
const nome = ref("")
const failed = ref(false)
const fetching = ref(false)

const deleting = ref("");

const page = ref(1)

const toast = useToast()

function previous() {
  page.value--;
  nextPage.value = [];
  fetchProfessores(true);
}

function next() { 
  page.value++;
  professores.value = [...nextPage.value];
  nextPage.value = [];
  fetchProfessores(false);
}

function deletar() {
  if(deleting.value.length < 1) {
    return;
  }
  deleteProfessor(deleting.value, (resp) => {
    deleting.value = "";
    if(resp.status != 200) {
      toast.error(resp.message, {timeout: 5000})
      return;
    }
    toast.success("Professor excluído com sucesso!")
    page.value = 1
    fetchProfessores(true)
  });
}

function handleResp(resp: APIResponse<Professor[]>, next: boolean) {
  fetching.value = false;
    if(resp.status != 200) {
      failed.value = true;
      toast.error(resp.message);
      return
    }
    failed.value = false;
    if(next) {
      nextPage.value = resp.body! as Professor[];
      return
    }
    professores.value = resp.body! as Professor[];
}

function fetchProfessores(fetchCurrent: boolean) {
  fetching.value = true;
  if(fetchCurrent) {
    findProfessores(nome.value, cpf.value, page.value, (resp: APIResponse<Professor[]>) => {
      handleResp(resp, false)
    })
  }

  // Busca a próxima pagina e deixa em cache.
  findProfessores(nome.value, cpf.value, page.value+1, (resp: APIResponse<Professor[]>) => {
    handleResp(resp, true)
  })
}

</script>

<template>
  <div v-if="deleting != ''" class="fixed left-[5%] md:left-[40%] w-80 z-10 text-gray-600 top-[40%] flex flex-col gap-y-8 p-8 text-center font-grotesk bg-white border-2 shadow-lg border-b-8 rounded-lg border-yellow-200">
    <img src="/warn.svg" class="w-12 mx-auto"/>
    <p>Excluindo professor de CPF: {{ deleting }}<br>Essa ação é irreversível. Tem certeza?</p>
    <div class="flex gap-x-2 mx-auto">
      <button class="bg-red-400 hover:bg-red-500" @mousedown="deleting = ''">Cancelar</button>
      <button class="bg-yellow-400 hover:bg-yellow-500" @mousedown="deletar">Confirmar</button>
    </div>
  </div>
  <ProtectedContent>
    <div>
      <header class="flex md:flex-row flex-col gap-y-2 gap-x-2">
        <div>
          <label for="nome">Nome</label>
          <input v-model="nome" id="nome" class="w-full md:w-72"/>
        </div>
        <div>
          <label for="CPF" >CPF</label>
          <input id="cpf" v-model="cpf" class="w-full md:w-48"/>
        </div>
        <button @mousedown="page = 1; fetchProfessores(true)" class="block">Buscar</button>
      </header>
      <div class="mt-8 w-full md:block hidden">
        <table class=" flex flex-col text-gray-600">
          <tr class="border-b-0">
            <th class="w-80">Nome</th>
            <th class="w-36">CPF</th>
            <th class="w-48">Telefone</th>
            <th class="w-96">Formação</th>
            <th class="w-24">Aulas dadas</th>
            <th>Ações</th>
          </tr>
          <tr class="font-grotesk text-sm" v-for="(professor, index) in professores" :class="[index == professores.length - 1 ? 'border-b-0' : '']">
            <td class="w-80">
              {{ professor.nome }}
            </td>
            <td class="w-36">
              {{ professor.cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, '$1.$2.$3-$4') }}
            </td>
            <td class="w-48">
              {{ professor.telefone }} 
            </td>
            <td class="w-96">
              {{ professor.formacao }}
            </td>
            <td class="w-24">
              {{ professor.aulasDadas }}
            </td>
            <td>
              <button class="bg-red-400 hover:bg-red-500" @mousedown="deleting = professor.cpf">Excluir</button>
            </td>
          </tr>
        </table>
      </div>
      <div class="block md:hidden mt-8 font-grotesk text-gray-600 text-xs">
        <div v-for="professor in professores" class="w-full border-[1px] border-indigo-500 p-4 rounded-md mt-2 flex gap-x-4 justify-between">
          <div class="flex flex-col gap-y-1">
            <p>{{ professor.nome }}</p>
            <p>{{ professor.cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, '$1.$2.$3-$4') }}</p>
            <p>{{ professor.telefone }}</p>
            <p>Formação: {{ professor.formacao }}</p>
            <p>Aulas dadas: {{ professor.aulasDadas == null || professor.aulasDadas == undefined ? 0 : professor.aulasDadas }}</p>
          </div>
          <div @mousedown="deleting = professor.cpf" class="h-fit mt-auto w-fit rounded-md p-2 bg-red-400">
            <img src="/trash.svg" class="w-6">
          </div>
        </div>
      </div>
      <div class="flex justify-end gap-x-2">
        <button v-if="page > 1" class="px-1 py-1" @mousedown="previous"><img src="/left.svg" class="w-4"/></button>
        <span class="my-auto  font-grotesk font-bold">{{ page }}</span>
        <button v-if="nextPage.length > 0" class="px-1 py-1"><img src="/right.svg" @mousedown="next" class="w-4"/></button>
      </div>
    </div>
  </ProtectedContent>
</template>