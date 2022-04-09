package main

import "fmt"

func main() {
	name := "Egin1"

	if name == "Egin" {
		fmt.Println("Hello", name)
	} else if name == "Ucok" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hello, boleh kenal ?")
	}

	if len(name) > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah OKE")
	}

	// If with short statement
	if name_len := len(name); name_len > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah OKE")
	}
}
