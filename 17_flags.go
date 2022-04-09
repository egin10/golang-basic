package main

import (
	"flag"
	"fmt"
)

func main() {
	hostname := flag.String("host", "localhost", "Put your database host")
	username := flag.String("user", "root", "Put your database user")
	password := flag.String("pass", "", "Put your database password")
	port := flag.Int("port", 3036, "Put your database port number")

	flag.Parse()

	fmt.Println("Hostname :", *hostname)
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
	fmt.Println("Port :", *port)
}
