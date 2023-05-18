package main

import (
	"fmt"
	"sort"
	"strings"
)

// Escreva uma função que receba um slice de palavras
// e retorne um mapa onde as chaves são palavras que são anagramas
// entre si e os valores são os grupos de palavras anagramas.

func encontrarAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {

		sorted := sortString(palavra)

		anagramas[sorted] = append(anagramas[sorted], palavra)
	}

	return anagramas
}

func sortString(s string) string {

	chars := strings.Split(s, "")

	sort.Strings(chars)

	return strings.Join(chars, "")
}

func main() {

	palavras := []string{"ovo", "amar", "eba", "oba", "voo", "raam", "bae", "eab"}

	fmt.Println(encontrarAnagramas(palavras))

}
