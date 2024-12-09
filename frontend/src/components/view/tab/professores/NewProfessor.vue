<script setup lang="ts">
import { vMaska } from "maska/vue"
import { createProfessor } from "../../../../api/professores";
import { ref } from "vue";
import { fetchCEP } from "../../../../api/cep";
import { allFilled } from "../../../../util/util";
import Loading from "../../../Loading.vue";
import ProtectedContent from "../../ProtectedContent.vue";
import { useToast } from "vue-toastification";

const nome = ref("");
const cpf = ref("");
const telefone = ref("");
const formacao = ref("");
const fetching = ref(false);
const failed = ref(false);
const errorMessage = ref("");
const cep = ref("");
const rua = ref("");
const bairro = ref("");
const cidade = ref("");
const numero = ref("");

const toast = useToast();

function clear() {
  cpf.value = "";
  cep.value = "";
  rua.value = "";
  bairro.value = "";
  cidade.value = "";
  numero.value = "";
  formacao.value = "";
  telefone.value = "";
  nome.value = "";
}

function onChange() {
  errorMessage.value = "";
  if(!allFilled(nome,cpf,telefone,formacao,cep,rua,bairro,cidade,numero)) {
    errorMessage.value = "Preencha todos os campos!"
    return
  }
}

function cadastrar() {
  if(!allFilled(nome,cpf,telefone,formacao,cep,rua,bairro,cidade,numero)) {
    toast.error("Preencha todos os campos.", {timeout:5000})
    return
  }
  fetching.value = true
  createProfessor({
    aulas_dadas: 0,
    cpf: cpf.value,
    endereco: {
      cep: cep.value,
      cidade: cidade.value,
      numero: numero.value,
      bairro: bairro.value,
      rua: rua.value
    },
    formacao: formacao.value,
    nome: nome.value,
    telefone: telefone.value
  }, (resp) => {
    fetching.value = false;
    if(resp.status != 200) {
      failed.value = true;
      toast.error(resp.message, {timeout:5000})
      return;
    }
    clear();
    toast.success("Professor cadastrado com sucesso.", {timeout:5000})
  })
}

function fetchCep() {
  fetchCEP(cep.value, (resp) => {
    rua.value = resp.logradouro;
    bairro.value = resp.bairro;
    cidade.value = `${resp.localidade} - ${resp.uf}`
    onChange()
  }, () => {
    toast.error("CEP inválido.", {timeout:5000})
  })
}

const handleNameChanged = () => nome.value = nome.value.replace(/\d+/g, '');
const handleCPFChanged = () => cpf.value = cpf.value.replace(/[^0-9Xx.\-]/g, '');
const handleTelefoneChanged = () => telefone.value = telefone.value.replace(/[^0-9\(\)\s-]/g, '');

</script>

<template>
  <ProtectedContent>
    <div class="flex flex-col gap-y-4">
      <label class="text-xs font-bold">* Informações pessoais</label>
      <div class="flex flex-col md:grid md:grid-cols-3 gap-y-4 gap-x-2 w-full md:w-fit">
        <div class="col-span-2">
          <label>Nome</label>
          <input :disabled="fetching" @input="(_:any) => {onChange(); handleNameChanged()}" v-model="nome" type="text" class="w-full md:w-96">
        </div>
        <div class="col-span-2 md:col-span-1">
          <label>CPF</label>
          <input @input="(_:any) => {onChange(); handleCPFChanged()}" :disabled="fetching" v-model="cpf" v-maska="'###.###.###-##'" placeholder="123.456.789-10" id="cpf" type="text" class="w-full md:w-48">
        </div>
        <div class="col-span-2">
          <label>Formação</label>
          <input :disabled="fetching" v-model="formacao" @input="onChange" itype="text" class="w-full md:w-96">
        </div>
        <div class="col-span-2 md:col-span-1">
          <label>Telefone</label>
          <input :disabled="fetching" v-model="telefone" @input="(_:any) => {onChange(); handleTelefoneChanged()}" v-maska="'(##) #####-####'" placeholder="(12) 34567-8910"  type="text" class="w-full md:w-48">
        </div>
      </div>
      <label class="text-xs font-bold">* Endereço</label>
      <div class="flex flex-col md:grid md:grid-cols-2 grid-rows-3 gap-y-4 gap-x-4 w-fit">
        <div class="flex gap-x-4">
          <div>
            <label>CEP</label>
            <input :disabled="fetching" v-maska="'#####-###'" @input="onChange" v-model="cep" type="text" class="w-32">
          </div>
          <div class="mt-auto">
            <button @mousedown="fetchCep" class="mt-auto">Buscar</button>
          </div>
        </div>
        <div class="md:block hidden"></div>
        <div class="">
          <label>Rua</label>
          <input disabled v-model="rua" @input="onChange" type="text" class="w-72 md:w-96">
        </div>
        <div class="col-span-1">
          <label>Número</label>
          <input v-maska="'#####'" :disabled="fetching" v-model="numero" @input="onChange" type="number" class="w-48" max="999999">
        </div>
        <div class="col-span-1">
          <label>Bairro</label>
          <input disabled v-model="bairro" @input="onChange" type="text" class="w-full md:w-96">
        </div>
        <div class="col-span-1">
          <label>Cidade</label>
          <input disabled @input="onChange" v-model="cidade"  type="text">
        </div>
      </div>
      <p class="text-xs font-grotesk text-red-400 font-bold">{{ errorMessage }}</p>
      <div class="flex gap-x-4">
        <button :disabled="fetching" @mousedown="cadastrar" class="w-32">Cadastrar</button>
        <Loading v-if="fetching"/>
      </div>
    </div>
  </ProtectedContent>
</template>