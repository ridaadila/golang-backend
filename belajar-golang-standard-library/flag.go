package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "localhost", "database username")
	var password *string = flag.String("password", "12345", "database password")
	var host *string = flag.String("host", "root", "database root")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()
	fmt.Println("Username: ", *username)
	fmt.Println("Password: ", *password)
	fmt.Println("Host: ", *host)
	fmt.Println("Port: ", *port)
}
