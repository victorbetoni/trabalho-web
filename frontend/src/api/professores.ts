import type Professor from "../model/Professor";
import { APIRequest, GET, POST, type Handler } from "./api";

export function findProfessores(nome: string, cpf: string, page: number, handler: Handler<Professor[]>) {
  GET(new APIRequest("professor", null, null, {
    nome: nome,
    cpf: cpf,
    page: page,
    limit: 5,
  }), handler)
}

export function createProfessor(prof: Professor, handler: Handler<any>) {
  POST(new APIRequest("professor", null, prof, null), handler)
}