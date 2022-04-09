package main

import "fmt"

func main() {
	name := "egin"

	switch name {
	case "Egin":
		fmt.Println("Hallo Egin")
	case "Eko":
		fmt.Println("Hallo Eko")

	default:
		fmt.Println("Hallo boleh kenalan?")
	}

	// Switch with short satement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah OKE")
	}

	// Switch without expression
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama cukup panjang")
	default:
		fmt.Println("Nama sudah OKE")
	}
}
