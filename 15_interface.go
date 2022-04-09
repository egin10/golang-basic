package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHallo(hasName HasName) {
	fmt.Println("Hallo", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

//Interface Kosong
func Test(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Unknown"
	}
}

func CobaInterfaceKosong() interface{} {
	return true
}

func main() {
	egin := Person{
		Name: "Egin",
	}
	sayHallo(egin)

	kucing := Animal{
		Name: "Kucing",
	}
	sayHallo(kucing)

	test := Test(2)
	fmt.Println(test)

	interfaceKosong := CobaInterfaceKosong()
	switch value := interfaceKosong.(type) {
	case string:
		fmt.Println(value)
	case int:
		fmt.Println(value)
	case bool:
		fmt.Println(value)
	default:
		fmt.Println("Unknown!")
	}
}
