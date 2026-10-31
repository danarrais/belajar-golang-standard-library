package main

import ("fmt"
		"strings"
	)

func main () {
	fmt.Println(strings.Contains("Danar Rais", "Danar"))
	fmt.Println(strings.Split("Danar Rais", " "))
	fmt.Println(strings.ToLower("Danar Rais"))
	fmt.Println(strings.ToUpper("Danar Rais"))
	fmt.Println(strings.Trim("      Danar Rais     ", " "))
	fmt.Println(strings.ReplaceAll("Danar Rais Danar Natasya", "Danar", "Aca"))

}