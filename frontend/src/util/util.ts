import type { Ref } from "vue";

export function allFilled(...refs: Ref<any>[]): boolean {
  return !refs.some(x => x.value == null || x.value == undefined || x.value == 0 || x.value == "");
}

export function filterLetters(event: any, ref: Ref<any>) {
  ref.value = event.target.value.replace(/[^a-zA-Z]/g, '');
}