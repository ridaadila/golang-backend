package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	connection := database.GetConnection()
	fmt.Println(connection)
}
