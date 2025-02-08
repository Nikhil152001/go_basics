package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {

	a := Person{
		Name:    "nikhil",
		Age:     24,
		Country: "india",
	}
	fmt.Println("Name:", a.Name)
	fmt.Println("Age:", a.Age)
	fmt.Println("Country:", a.Country)
}
