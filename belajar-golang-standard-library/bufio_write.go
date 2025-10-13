package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Ini rida\n")
	writer.WriteString("Asal surabaya\n")
	writer.Flush()
}
