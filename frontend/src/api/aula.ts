import { APIRequest, POST, type Handler } from "./api";

export function createAula(prof: string, title: string, time: number, handler: Handler<any>) {
  POST(new APIRequest("aula", null, ({
    professor: prof,
    titulo: title,
    time: time,
  }), null), handler)
}