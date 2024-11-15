import type Endereco from "./Endereco";

export default interface Professor {
  nome: string;
  aulasDadas: number;
  cpf: string;
  telefone: string;
  endereco: Endereco;
}