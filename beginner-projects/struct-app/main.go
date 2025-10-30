package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName   string
	lastName    string
	dateOfBirth string
	created     time.Time
}

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userDateOfBirth := getUserData("Enter your date of birth: ")

	var user *user
	user, err := populateUser(userFirstName, userLastName, userDateOfBirth)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputUser(user)

}

func populateUser(firstName, lastName, dateOfBirth string) (*user, error) {
	if firstName == "" || lastName == "" || dateOfBirth == "" {
		return nil, errors.New("All fields are required")
	}

	user := &user{
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: dateOfBirth,
		created:     time.Now(),
	}

	return user, nil
}

func outputUser(u *user) {
	fmt.Println(u.firstName, u.lastName, u.dateOfBirth, u.created)
}

func getUserData(text string) string {
	var userValue string
	fmt.Println(text)
	fmt.Scanln(&userValue)
	return userValue
}
