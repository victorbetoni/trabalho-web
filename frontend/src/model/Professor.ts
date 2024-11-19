import type Endereco from "./Endereco";

export default interface Professor {
  nome: string;
  aulasDadas: number;
  cpf: string;
  formacao: string;
  telefone: string;
  endereco: Endereco;
}