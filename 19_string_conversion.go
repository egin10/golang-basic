package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error :", err.Error())
	}

	integerValue, err := strconv.ParseInt("20", 10, 32)
	if err == nil {
		fmt.Println(integerValue)
	} else {
		fmt.Println("Error :", err.Error())
	}

	floatValue, err := strconv.ParseFloat("2.5", 8)
	if err == nil {
		fmt.Println(floatValue)
	} else {
		fmt.Println("Error :", err.Error())
	}

	numberToString := strconv.FormatInt(10000, 10)
	fmt.Println(numberToString)
}
