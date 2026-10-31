package main

import(
	"fmt"
	"encoding/csv"
	"strings"
	"io"
)

func main() {
	csvString := "eko, kurniawan, khannedy\n" +
	"budi, pratama, pratama\n" +
	"joko, wi, dodo"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record,err := reader.Read()
		if err ==  io.EOF {
			break
		}

		fmt.Println(record)
	}
}