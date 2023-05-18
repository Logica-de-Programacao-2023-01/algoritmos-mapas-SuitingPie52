package main

import "fmt"

// Escreva uma função que receba um mapa com valores inteiros e retorne a soma de todos os valores.

func somaMapa(m map[string]int) int {

	soma := 0

	for _, ranM := range m {

		soma += ranM

	}

	return soma

}

func main() {

	m := map[string]int{

		"a": 423,
		"b": 1212,
		"c": 212,
	}

	fmt.Println(somaMapa(m))

}