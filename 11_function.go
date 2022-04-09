package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World!")
}

func sayHelloTo(from string, to string) {
	fmt.Println(from + ", say hello to " + to)
}

func getHello(name string) string {
	return "Hello " + name
}

func hitung(a int, b int) int {
	return a + b
}

func multipleReturn() (string, string) {
	return "Hilih", "Hiyaaa"
}

func main() {
	sayHello()
	sayHelloTo("Udin", "Ucok")

	result := getHello("Egin")
	fmt.Println(result)

	hasil := hitung(2, 4)
	fmt.Println(hasil)

	fmt.Println(multipleReturn())

	firstName, lastName := multipleReturn()
	fmt.Println(firstName, lastName)
}
