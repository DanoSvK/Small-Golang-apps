package main

import (
	"errors"
	"fmt"
	"structs/user"
	"time"
)

func main() {
	firstName := getStructData("Enter your first name: ")
	lastName := getStructData("Enter your last name: ")
	birthDate := getStructData("Enter your birth date: ")

	user, err := populateStruct(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.FirstName)
	fmt.Println(user.LastName)
	fmt.Println(user.BirthDate)
}

func getStructData(text string) string {
	var value string
	fmt.Println(text)
	fmt.Scan(&value)
	return value
}

func populateStruct(firstName, lastName, birthDate string) (*user.User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("You must enter all fields")
	}

	u := &user.User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		UpdatedAt: time.Now(),
	}

	return u, nil
}
