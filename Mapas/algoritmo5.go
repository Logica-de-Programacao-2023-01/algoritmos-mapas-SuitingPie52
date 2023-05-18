package main

import (
	"fmt"
	"strings"
)

// Escreva uma função que receba uma string
// e retorne um mapa onde as chaves são os caracteres presentes na string e os
// valores são a frequência de cada caractere.

func frequenciaLetras(s string) map[string]int {

	m1 := make(map[string]int)
	chars := strings.Split(s, "")

	for _, ranC := range chars {

		m1[ranC] += 1

	}

	return m1

}

func main() {

	fmt.Println(frequenciaLetras("õ ovo ou a galinha"))

}
