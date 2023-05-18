package main

import (
	"fmt"
	"strings"
)

// Escreva uma função que conte a ocorrência de cada palavra em uma string
// e retorne um mapa onde as chaves são as palavras encontradas
// e os valores são o número de ocorrências de cada palavra.

func contarPalavras(s string) map[string]int {

	frequenciaPalavras := make(map[string]int)

	strings.Fields(s)

	for _, ranS := range strings.Fields(s) {

		frequenciaPalavras[ranS] += 1

	}

	return frequenciaPalavras

}

func main() {

	s := "oba oba he he kkkkkkkkkkkkkk oba"

	fmt.Println(contarPalavras(s))

}
