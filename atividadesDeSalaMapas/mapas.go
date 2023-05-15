package main

import (
	"fmt"
	"strings"
)

func contarCaracteres(frase string) map[string]int {

	conjFrase := strings.Split(frase, "")
	frequeciaCaracteres := make(map[string]int)

	for _, ranConjFrase := range conjFrase {

		frequeciaCaracteres[ranConjFrase]++

	}

	return frequeciaCaracteres

}

func main() {

	s := "Hello World!"

	mapa := contarCaracteres(s)
	fmt.Println(mapa["l"])

}
