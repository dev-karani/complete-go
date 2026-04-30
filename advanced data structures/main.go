package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main(){
	person := Person{Name: "John", Age: 25}
	fmt.Printf("THis si our person %v\n", person)
}