package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(10, 20, 30, 10)
	fmt.Println(total)

	slice := []int{12, 35, 12, 23}
	newTotal := sumAll(slice...)
	fmt.Println(newTotal)
}
