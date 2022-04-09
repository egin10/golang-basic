package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke", i)
	}

	slice := []string{"Egin", "Emang", "Yihaa"}

	// Cara 1
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Cara 2 (For Range)
	for index, value := range slice {
		fmt.Println("Index", index, "=", value)
	}

	// Print hanya value saja
	for _, value := range slice {
		fmt.Println("Name", value)
	}

	// For Range untuk Map
	person := make(map[string]string)
	person["name"] = "Egin"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
