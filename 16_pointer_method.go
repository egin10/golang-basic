package main

import "fmt"

type Man struct {
	Name string
}

// Agar data asli dari Struct berubah, maka harus gunakan Pointer
// Jika tidak pakai Pointer maka data hanya akan tercopy dan tidak mengubah data asli
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

// Agar data bisa berubah dalam function yang dibuat, maka harus menggunakan Pointer
func GantiNama(man *Man, name string) {
	man.Name = name
}

func main() {
	egin := Man{
		Name: "Egin",
	}
	egin.Married()
	fmt.Println(egin.Name)

	GantiNama(&egin, "Ucok")
	egin.Married()
	fmt.Println(egin.Name)
}
