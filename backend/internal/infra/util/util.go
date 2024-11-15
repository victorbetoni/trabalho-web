package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ValidCPF(cpf string) bool {
	numeros := regexp.MustCompile(`\D`).ReplaceAllString(cpf, "")
	if len(cpf) != 11 || strings.Repeat(string(cpf[0]), 11) == cpf {
		return false
	}

	// ChatGPT, confesso.
	pesos1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	digito1 := calcularDigitoVerificador(numeros, pesos1)

	pesos2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	digito2 := calcularDigitoVerificador(numeros, pesos2)

	digitoVerificador := cpf[9:]
	return fmt.Sprintf("%d%d", digito1, digito2) == digitoVerificador

}

// ChatGPT, confesso.
func calcularDigitoVerificador(cpf string, pesos []int) int {
	soma := 0
	for i, peso := range pesos {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * peso
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}
