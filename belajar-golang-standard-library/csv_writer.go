package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	writer.Write([]string{"rida", "tes", "adila"})
	writer.Write([]string{"golang", "java", "ruby"})

	writer.Flush()
}
