import type Professor from "./Professor";

export interface Aula {
  id: string;
  professor: Professor;
  titulo: string;
  carga_horaria: number;
  data: string;
}