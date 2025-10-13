package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "rida,adila,tes\n" +
		"kana,rekha\n" +
		"devi,hainun"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
