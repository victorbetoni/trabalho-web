import type Endereco from "./Endereco";

export default interface Professor {
  nome: string;
  aulas_dadas: number;
  cpf: string;
  formacao: string;
  telefone: string;
  endereco: Endereco;
}