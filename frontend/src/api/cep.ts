import type CEPResponse from "../model/cep/CEPResponse";

export function fetchCEP(cep: string, handler: (resp: CEPResponse) => void) {
  fetch(`${import.meta.env.VITE_VIACEP_EP}/${cep}/json`).then(x => x.json()).then(j => handler(j as CEPResponse))
}