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

	//if else
	if age >= 18 {
		fmt.Println("you are an adult")
	} else if age >=13 {
		fmt.Println("you are a teenager")
	}else {
		fmt.Println("you are a child")
	}

	//switches
	day := "tuesday"

	switch day {
	case "monday":
		fmt.Println("start the week")
	case "tuesday", "wednesday","thursday":
		fmt.Println("midweek")
	case "friday":
		fmt.Println("tgif")
	default:
		fmt.Println("its the weekend")
	}
	
	//for loop
	for i := 0; i <3; i++  {
		fmt.Println("this is i," ,i)
	}

	//"while loop"
	counter := 0  //counter
	for counter < 3 {
		fmt.Println("this is the counter", counter)
		counter ++ //increment  
	}
	//explicit infinite loop
	iterations :=0 
	for{
		//some condition is met
		if iterations > 3 {
			break
		}
		fmt.Println(iterations)
		iterations ++
	}

	//arrays & slices
	numbers := [5] int{20,30,40,50,60}
	fmt.Printf("this is our arrray %v\n", numbers)
	fmt.Printf("this is our arrray %v\n", numbers[0])

	//multidimensional arrays
	matrix := [2][3] int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Printf("this is our matrix: %v\n", matrix)

	//slices


	fruits := []string{"apple", "banana", "strawberry"}
	fmt.Printf("these are my fruits %v\n", fruits)

	fruits = append(fruits, "kiwi")
	fmt.Printf("these are my fruits %v\n", fruits)

	morefruits := []string{"blueberries", "tomato"}
	fruits = append(fruits, morefruits...)
	fmt.Printf("these are my fruits %v\n", fruits)

	// nums := []int{1,2,3,4}
	// numsMixed := append(nums, morefruits)

	//iterating on arrays & slices
	for index, value := range numbers{
		fmt.Printf("index %d and value %d \n", index, value)
	}


	//maps
	capitalCities := map[string]string {
		"USA": "Washington D.C",
		"India": "New Delhi",
		"UK": "London",
	}
	fmt.Println(capitalCities["USA"])  //print value
	fmt.Println(capitalCities["Germany"])  //checks but return ""

	//checks for porperties and values
	capital, exist := capitalCities["Germany"]

	if exist {
		fmt.Println("this is the capital", capital)
	} else{
		fmt.Println("does not exist")
		print(exist)
	}

	//to delete key
	delete(capitalCities, "UK")
	// fmt.Printf("this is the new deleted map %V\n", capitalCities)
}


	//functions
	func add(a int,b int) int {
		return a + b
	}

	func calculatesumproduct(a, b int) (int, int) {
		return a+b, a*b
	}

