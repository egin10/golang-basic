package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	//fmt.Println(args)

	for i, value := range args {
		if i != 0 {
			fmt.Println(value)
		}
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", hostname)
	} else {
		fmt.Println("Error :", err.Error())
	}

	dbUsername := os.Getenv("APP_DB_USERNAME")
	dbPassword := os.Getenv("APP_DB_PASSWORD")

	fmt.Println(dbUsername, dbPassword)
}
