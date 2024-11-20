<script setup lang="ts">
import { Ref, ref } from 'vue';
import Professor from '../../../../model/Professor';
import { findProfessores } from '../../../../api/professores';
import { APIResponse } from '../../../../api/api';
import ProtectedContent from '../../ProtectedContent.vue';

const professores: Ref<Professor[]> = ref([])

const cpf = ref("")
const nome = ref("")
const failed = ref(false)
const fetching = ref(false)
const errorMessage = ref("")

const page = ref(1)


function fetchProfessores() {
  fetching.value = true;
  findProfessores(nome.value, cpf.value, page.value, (resp: APIResponse<Professor[]>) => {
    fetching.value = false;
    if(resp.status != 200) {
      failed.value = true;
      errorMessage.value = resp.message!;
      return
    }
    errorMessage.value = "";
    failed.value = false;
    professores.value = resp.body! as Professor[];
    console.log(professores.value)
  })
}

</script>

<template>
  <ProtectedContent>
    <div>
      <header class="flex gap-x-2">
        <div>
          <label for="nome">Nome</label>
          <input id="nome" class="w-72"/>
        </div>
        <div>
          <label for="CPF" >CPF</label>
          <input id="cpf" class="w-48"/>
        </div>
        <button @mousedown="fetchProfessores" class="block">Buscar</button>
      </header>
      <div class="mt-8 w-full">
        <table class=" flex flex-col text-gray-600">
          <tr class="border-b-0">
            <th class="w-80">Nome</th>
            <th class="w-36">CPF</th>
            <th class="w-48">Telefone</th>
            <th class="w-96">Formação</th>
            <th class="w-24">Aulas dadas</th>
            <th>Ações</th>
          </tr>
          <tr class="font-grotesk text-lg" v-for="(professor, index) in professores" :class="[index == professores.length - 1 ? 'border-b-0' : '']">
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
          </tr>
        </table>
      </div>
    </div>
  </ProtectedContent>
</template>