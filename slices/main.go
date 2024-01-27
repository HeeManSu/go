package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Let's learn slices in array")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitLIST IS %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// Highscore
	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867

	highscores = append(highscores, 555, 666, 321)
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))

	//How to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses  = append(courses[:index], courses[index + 1:]...)
	fmt.Println(courses)

}
