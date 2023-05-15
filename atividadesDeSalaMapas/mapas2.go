package main

import "fmt"

func mediaDeNotas(mapa map[string][]float64) map[string]float64 {

	newMap := make(map[string]float64)

	for aluno, conjNotas := range mapa {

		soma := 0.0

		for _, ranConjNotas := range conjNotas {

			soma += ranConjNotas
			media := soma / float64(len(conjNotas))
			newMap[aluno] = media

		}

	}

	return newMap

}

func main() {

	alunos := map[string][]float64{

		"Gabriel": {10, 9, 8},
		"João":    {5, 10, 3},
	}

	fmt.Println(mediaDeNotas(alunos))

	fmt.Println(mediaDeNotas(alunos)["Gabriel"])
}
