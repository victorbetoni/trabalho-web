import type Professor from "../model/Professor";
import { APIRequest, GET, POST, type Handler } from "./api";

export function findProfessores(nome: string, cpf: string, handler: Handler<Professor[]>) {
  GET(new APIRequest("professores", null, null, {
    nome: nome,
    cpf: cpf,
  }), handler)
}

export function createProfessor(prof: Professor, handler: Handler<any>) {
  POST(new APIRequest("professores", null, prof, null), handler)
}