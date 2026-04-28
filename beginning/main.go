package main

import (
	"fmt"
)

func main() {

	var name string = "melkey"

	age := 27
	var city string
	city = "seattle"
	var country, continent string = "usa", "north america"

	var (
		isEmployed bool = true
		salary     int     = 4000
		position   string  = "developer"
	)

	fmt.Printf("my name is %s , and i am %d years old, living in %s, %s, %s\n", name, age, city, country, continent)
	fmt.Printf("i am employed is %t and this is my salary %d , and this is my position %s\n", isEmployed, salary, position)

	//zero values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("%d, %f, %s, %t\n", defaultInt,defaultFloat,defaultString,defaultBool)

	//consts
	const pi =3.144

	const (
		monday =1
		tuesday =2
		wednesday =3
	)
	fmt.Printf("Monday: %d -Tuessday:%d - wednesday:%d\n", monday, tuesday, wednesday)

	const typed int =25
	const untyped =25

	fmt.Println(typed == untyped)

	//enums
	const (
		jan int = iota + 1 //1
		feb //2
		mar //3
		apr //4
	)

	fmt.Printf("jan -%d, feb -%d, mar -%d, april -%d\n", jan, feb,mar,apr)
	
	result := add(3,4)
	fmt.Println(result)

	sum, product := calculatesumproduct(2,3)
	fmt.Println(sum, product)
}
	//functions
	func add(a int,b int) int {
		return a + b
	}

	func calculatesumproduct(a, b int) (int, int) {
		return a+b, a*b
	}

