import type Professor from "../model/Professor";
import { APIRequest, GET, POST, PUT, type Handler } from "./api";

export function findProfessores(nome: string, cpf: string, page: number, handler: Handler<Professor[]>) {
  GET(new APIRequest("professor", null, null, {
    nome: nome,
    cpf: cpf,
    page: page,
    limit: 5,
  }), handler)
}

export function findProfessor(cpf: string, handler: Handler<Professor[]>) {
  GET(new APIRequest("professor", null, null, {
    nome: "",
    cpf: cpf,
    page: 1,
    limit: 1,
  }), handler)
}

export function updateProfessor(prof: Professor, handler: Handler<any>) {
  PUT(new APIRequest("professor", null, prof, null), handler)
}

export function createProfessor(prof: Professor, handler: Handler<any>) {
  POST(new APIRequest("professor", null, prof, null), handler)
}