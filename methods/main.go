package main

import "fmt"

func main() {

	User := User{
		Name:   "Himanshu",
		Email:  "Himanshu@taghash.io",
		Status: true,
		Age:    18,
	}

	User.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
