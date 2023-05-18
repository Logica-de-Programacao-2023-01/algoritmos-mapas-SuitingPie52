package main

import (
	"fmt"
	"strings"
)

// Escreva uma função que receba uma string contendo uma frase
// e retorne um mapa onde as chaves são as palavras encontradas na frase
// e os valores são mapas contendo a contagem de cada letra em cada palavra.

func contarLetrasEPalavras(s string) map[string]map[string]int {

	mF := make(map[string]map[string]int)

	palavras := strings.Fields(s)

	chars := []string{}

	for _, ranP := range palavras {

		chars = strings.Split(ranP, "")

		mF[ranP] = make(map[string]int)

		for _, ranC := range chars {

			mF[ranP][ranC]++

		}

	}

	return mF

}

func main() {

	fmt.Println(contarLetrasEPalavras("santos time da virada"))

}