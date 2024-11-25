import type Professor from "./Professor";

export interface Aula {
  professor: Professor;
  titulo: string;
  time: number;
}