package main

import "fmt"

// Escreva uma função que receba um mapa onde as chaves são nomes de pessoas
// e os valores são as quantias de dinheiro
// que cada pessoa gastou em uma conta compartilhada.
// A função deve calcular o valor que cada pessoa deve receber ou
// pagar para igualar as despesas.

func igualarDespesas(m1 map[string]float64) {

	maior := 0.0

	for _, ranM1 := range m1 {

		if ranM1 > maior {

			maior = ranM1

		}

	}

	for nome, quantia := range m1 {

		quantia = maior - quantia

		fmt.Printf("%s deve pagar %.2f reais para  igualar.\n", nome, quantia)

	}

}

func main() {

	m1 := make(map[string]float64)
	m1["Gabriel"] = 10000
	m1["Isabelle"] = 100
	m1["Fábio"] = 500

	igualarDespesas(m1)

}
