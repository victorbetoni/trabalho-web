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


function fetchProfessores() {
  fetching.value = true;
  findProfessores(nome.value, cpf.value, (resp: APIResponse<Professor[]>) => {
    fetching.value = false;
    if(resp.status != 200) {
      failed.value = true;
      errorMessage.value = resp.message!;
      return
    }
    errorMessage.value = "";
    failed.value = false;
    professores.value = resp.body!
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
        <table>
          <tr>
            <th class="w-96">Nome</th>
            <th class="w-32">CPF</th>
            <th class="w-32">Telefone</th>
            <th class="w-24">Aulas dadas</th>
            <th>Ações</th>
          </tr>
          <tr v-for="(professor, index) in professores" :class="[index == professores.length - 1 ? 'border-b-0' : '']">
            <td>
              {{ professor.nome }}
            </td>
            <td>
              {{ professor.cpf }}
            </td>
            <td>
              {{ professor.telefone }} 
            </td>
            <td>
              {{ professor.aulasDadas }}
            </td>
          </tr>
        </table>
      </div>
    </div>
  </ProtectedContent>
</template>