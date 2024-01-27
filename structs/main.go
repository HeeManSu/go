package main

import "fmt"

// Composite data type that groups together variables under a single name. Similar to classed but they don't have methods attached to them.

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang; No super or parent class

	himanshu := User{"Himanshu", "Himashu@taghash.io", true, 18}
	fmt.Println(himanshu)
	fmt.Printf("hitesh details are: %+v\n", himanshu)
	fmt.Printf("Name is %v and email is %v\n", himanshu.Name, himanshu.Email)

	//create an instance of person struct

	Person := Person{
		FirstName: "Himanshu",
		LastName:  "Sharma",
		Age:       20,
	}

	fmt.Println("First Name: ", Person.FirstName)
	fmt.Println("Last Name: ", Person.LastName)
	fmt.Println("Age: ", Person.Age)

}
