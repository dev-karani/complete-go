package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main(){
	person := Person{Name: "John", Age: 25}
	fmt.Printf("THis is00 our person %+v\n", person)

	//anonymous structs defined in place
	employee := struct{
		name string
		id int
	}{
		name: "alice",
		id: 3,
	}
	fmt.Println("this is employee", employee)

	//nested structs
	type Address struct {
		Street string
		City string
	}

	type Contact struct {
		Name string
		Address Address
		Phone string
	}
	
	contact := Contact{
		Name: "june",
		Address: Address{
			Street: "123 Main Street",
			City: "New Delhi",
		},
	}
	fmt.Println("heres my contact", contact)

	//Pointers
	fmt.Println("name before", person.Name)
	person.modPersonName("ed")
	fmt.Println("name after:", person.Name)

	
}
//pass by reference
//method for structs with method receivers
func (p *Person)modPersonName (name string) {
	p.Name = name
	fmt.Println("inside scope new name," ,p.Name)
}