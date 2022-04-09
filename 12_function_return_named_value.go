package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Egin"
	middleName = "Keren"
	lastName = "Banget"
	return
}

func main() {
	firstName, middleName, lastName := getFullName()

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
