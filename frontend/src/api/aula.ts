import { APIRequest, DELETE, GET, POST, type Handler } from "./api";

export function createAula(prof: string, title: string, data: string, horario: string, carga: number, handler: Handler<any>) {
  POST(new APIRequest("aula", null, ({
    cpf: prof,
    titulo: title,
    data: data,
    horario: horario,
    carga_horario: carga,
  }), null), handler)
}

export function fetchAulas(page: number, cpf: string, handler: Handler<any>) {
  GET(new APIRequest("aula", null, null, ({
    page: page,
    cpf: cpf
  })), handler)
}

export function deleteAula(id: string, handler: Handler<any>) {
  DELETE(new APIRequest("aula", null, ({
    id: id
  }), null), handler)
}