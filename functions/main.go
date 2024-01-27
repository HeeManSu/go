package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 5, 33, 2, 34, 5, 53)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Message with pro result function is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"

}

func greeter() {
	fmt.Println("Namaste from Himanshu")
}
