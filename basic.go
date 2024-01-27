package main

import (
	"bufio"
	"fmt"
	"os"
)

const LoginToken string = "assbshshshs"

func main() {
	var username string = "Himanshu"
	var isLoggedIn bool = false
	var smallVal uint8 = 255
	var smallFloat float64 = 255.232323232323

	//Print is used to print but not in the new line.
	//Println is used print in the next line. 
	//Printf is used to printing the formatted output based on format string. 
	// % are used to format string with placeholder.
	// %t is used for formatting boolean values. 
	// %v used to format any type of value. 

	fmt.Println(smallFloat)
	fmt.Println(smallVal)
	fmt.Println(isLoggedIn)
	fmt.Println(username)
	fmt.Printf("Variables if of type: %T \n", username)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implict type
	var website = "cyborg.t5"
	fmt.Println(website)

	//no var style (Allowed only inside any function)
	numberOfUser := 300000.00
	fmt.Println(numberOfUser)
	

	//const can't be changed but var can be changed.
	fmt.Println(LoginToken)
	fmt.Printf("Variable if of type: %T \n", LoginToken)

	//Taking user input

	welcome := "Welcome himanshu sharama"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating for out pizza: ")

	//Comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is:  %T \n", input)
}
