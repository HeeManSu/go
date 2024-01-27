package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("If else in golang")

	loginCount := 24
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Non identifiable"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}

	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move to spot 3")
	case 4:
		fmt.Println("You can move to spot 4")
	case 5:
		fmt.Println("You can move to spot 5")
	case 6:
		fmt.Println("You can move to spot 6")
	}

	// Loops in Golang

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	for _, day := range days {
		fmt.Printf("Index is and value is %v\n", day)
	}

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 5 {
			rougeValue++
			continue
		}

		if rougeValue == 2 {
			goto lco
		}

		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at loo[ps]")

}
