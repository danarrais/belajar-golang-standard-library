package main

import(
	"os"
	"encoding/csv"
)

func main () {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"eko", "kuriawan", "khannedy"})
	_ = writer.Write([]string{"joko", "wi", "dodo"})
	_ = writer.Write([]string{"prab", "owo", "teddy"})
	
	writer.Flush()
}