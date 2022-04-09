package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "Ini adalah sebuah kalimat tanya?"
	if strings.Contains(kalimat, "?") {
		fmt.Println("Terdapat tanda baca ?")
	}

	sliceKalimat := strings.Split(kalimat, " ")
	for _, value := range sliceKalimat {
		fmt.Println(value)
	}

	menjadiHurufBesar := strings.ToUpper(kalimat)
	fmt.Println(menjadiHurufBesar)

	menjadiHurufKecil := strings.ToLower(kalimat)
	fmt.Println(menjadiHurufKecil)

	trimKalimat := strings.Trim("         "+kalimat+"         ", " ")
	fmt.Println(trimKalimat)

	gantiKata := strings.ReplaceAll(kalimat, "kalimat", "kata-kata")
	fmt.Println(gantiKata)
}
