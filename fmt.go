package main

import "fmt"

func main(){
	firstName := "Danar"
	lastName := "Rais"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}