package main

import "fmt"

// Escreva uma função que gere a sequência de Fibonacci até um determinado número n
// e retorne um mapa onde as chaves são
// os índices da sequência e os valores são os números correspondentes.

func fibonacci(n int) map[int]int {

	fibonas := make(map[int]int)

	for i := 0; len(fibonas) != n; i++ {

		if i < 2 {

			fibonas[i] = i

		} else {

			fibonas[i] = fibonas[i-2] + fibonas[i-1]

		}

	}

	return fibonas

}

func main() {

	fmt.Println(fibonacci(20))

}
