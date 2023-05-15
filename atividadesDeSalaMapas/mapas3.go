package main

import "fmt"

func tirarRepetidos(conj []int) []int {

	frequenciaNum := make(map[int]int)

	for i := 0; i < len(conj); i++ {

		frequenciaNum[conj[i]]++

	}

	for numeros, ranFrequenciaNum := range frequenciaNum {

		if ranFrequenciaNum > 1 {

			delete(frequenciaNum, numeros)

		}

	}

	newConj := []int{}

	for numeros := range frequenciaNum {

		newConj = append(newConj, numeros)

	}

	return newConj

}

func main() {

	conj := []int{5, 25, 35, 75, 75, 2, 5, 25}

	fmt.Println(tirarRepetidos(conj))

}