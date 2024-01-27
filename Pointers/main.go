package main

import "fmt"

func main() {

	// A pointer is a variable that stores the memory of another variable. 
	// Allows you to indirectly reference or access the value of a variable by using the memory address.

	fmt.Println("Learning pointers")
	// var ptr *int
	// fmt.Println("Value of pointer is: ",ptr)
	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is: ", ptr)
	fmt.Println("Value of actual pointer is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is: ", myNumber)

	//Declartion of pointer
	
	var x int
	// var ptr *int

	//Initialization of pointers 
	// It is initialized using the & operator which returns the memory address of a variable. 
	// The zero value of pointer is nil. 

	// x := 42
	// ptr := &x

	fmt.Println("Value of x:", x)


}
