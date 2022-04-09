package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Function / Method Struct
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello,", name, "my name is", customer.Name)
}

func main() {
	var budi Customer
	budi.Name = "Budi"
	budi.Address = "Jakarta"
	budi.Age = 20

	joko := Customer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     23,
	}

	joko.sayHello("Udin")

	fmt.Println(budi)
	fmt.Println(joko)
}
