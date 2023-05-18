package main

import "fmt"

// Escreva uma função que receba um slice de inteiros
// e retorne um mapa onde as chaves são os pares de números
// encontrados no slice e os valores são as frequências de cada par.

func pares(s []int) map[int]int {

	frequenciaS := make(map[int]int)

	for _, ranS := range s {

		frequenciaS[ranS] += 1

	}

	for numero, frequencia := range frequenciaS {

		if frequencia%2 != 0 {

			delete(frequenciaS, numero)

		} else {

			frequenciaS[numero] /= 2

		}

	}

	return frequenciaS

}

func main() {

	s := []int{0, 2, 3, 1, 2, 1, 2, 5, 5, 3, 4, 4, 4, 4}

	fmt.Println(pares(s))

}
