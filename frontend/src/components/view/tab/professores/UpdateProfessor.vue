<script setup lang="ts">
import { vMaska } from "maska/vue"
import { findProfessor, updateProfessor } from "../../../../api/professores";
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
const locked = ref(false);

const toast = useToast();

function clear() {
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

function update() {
  if(!allFilled(nome,cpf,telefone,formacao,cep,rua,bairro,cidade,numero)) {
    toast.error("Preencha todos os campos.", { timeout: 5000, })
    return
  }
  fetching.value = true
  updateProfessor({
    aulasDadas: 0,
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
      toast.error(resp.message, { timeout: 5000, })
      return;
    }
    toast.success("Professor atualizado com sucesso.", { timeout: 5000 });
    clear(); 
  })
}

function fetchCPF() {
  errorMessage.value = "";
  if (cpf.value.length < 14) {
    toast.error("Digite um CPF.", { timeout: 5000, })
    return
  }
  clear();
  findProfessor(cpf.value, (resp) => {
    if(resp.status != 200) {
      toast.error(resp.message, { timeout: 5000, })
      return
    }
    if(resp.body!.length < 1) {
      toast.error("Nenhum professor encontrado.", { timeout: 5000, })
      return;
    }
    resp.body!.forEach(x => {
      cep.value = x.endereco.cep;
      rua.value = x.endereco.rua;
      bairro.value = x.endereco.bairro;
      cidade.value = x.endereco.cidade;
      numero.value = x.endereco.numero;
      formacao.value = x.formacao;
      telefone.value = x.telefone;
      nome.value = x.nome;
      onChange();
    })
    locked.value = true;
  })
}

function fetchCep() {
  fetchCEP(cep.value, (resp) => {
    rua.value = resp.logradouro;
    bairro.value = resp.bairro;
    cidade.value = `${resp.localidade} - ${resp.uf}`
    onChange()
  }, () => {
    toast.error("CEP inválido.", { timeout: 5000, })
  })
}

const handleNameChanged = () => nome.value = nome.value.replace(/\d+/g, '');
const handleCPFChanged = () => cpf.value = cpf.value.replace(/[^0-9Xx.\-]/g, '');
const handleTelefoneChanged = () => telefone.value = telefone.value.replace(/[^0-9\(\)\s-]/g, '');

</script>

<template>
  <ProtectedContent>
    <SuccessModal message="teste"/>
    <div class="flex flex-col gap-y-4">
      <label class="text-xs font-bold">CPF</label>
      <div class="flex gap-x-2">
        <input @input="(_:any) => {onChange(); handleCPFChanged()}" :disabled="fetching || locked" v-model="cpf" v-maska="'###.###.###-##'" placeholder="123.456.789-10" id="cpf" type="text" class="w-48">
        <button @mousedown="fetchCPF" class="mt-auto">Buscar</button>
      </div>
      <label class="text-xs font-bold">* Informações pessoais</label>
      <div class="flex flex-col md:grid md:grid-cols-3 gap-y-4 gap-x-2 w-fit">
        <div class="col-span-2">
          <label>Nome</label>
          <input :disabled="fetching" @input="(_:any) => {onChange(); handleNameChanged()}" v-model="nome" type="text" class="w-72 md:w-96">
        </div>
        <div class="col-span-2">
          <label>Formação</label>
          <input :disabled="fetching" v-model="formacao" @input="onChange" itype="text" class="w-72 md:w-96">
        </div>
        <div class="col-span-1">
          <label>Telefone</label>
          <input class="w-72 md:w-96" :disabled="fetching" v-model="telefone" @input="(_:any) => {onChange(); handleTelefoneChanged()}" v-maska="'(##) #####-####'" placeholder="(12) 34567-8910"  type="text">
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
        <div class="w-96">
          <label>Rua</label>
          <input disabled v-model="rua" @input="onChange" type="text" class="w-72 md:w-96">
        </div>
        <div class="col-span-1">
          <label>Número</label>
          <input v-maska="'#####'" :disabled="fetching" v-model="numero" @input="onChange" type="number" class="w-48" max="999999">
        </div>
        <div class="col-span-1">
          <label>Bairro</label>
          <input disabled v-model="bairro" @input="onChange" type="text" class="w-72 md:w-96">
        </div>
        <div class="col-span-1">
          <label>Cidade</label>
          <input disabled @input="onChange" v-model="cidade"  type="text">
        </div>
      </div>
      <p class="text-xs font-grotesk text-red-400 font-bold">{{ errorMessage }}</p>
      <div class="flex gap-x-4">
        <button v-if="locked" @mousedown="locked = false; clear()" class="w-32 bg-red-500">Cancelar</button>
        <button :disabled="fetching || !locked" @mousedown="update" class="w-32">Atualizar</button>
        <Loading v-if="fetching"/>
      </div>
    </div>
  </ProtectedContent>
</template>