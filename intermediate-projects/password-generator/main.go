package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	lowercase := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z"}
	uppercase := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	symbols := []string{"!", "#", "$", "%", "&", "(", ")", "*", "+"}
	password := genPwd(uppercase, lowercase, numbers, symbols)

	fmt.Println("Result:", password)
}

func getPwdLen() int {
	var length int
	fmt.Print("Please, enter length of your password between 1 and 128: ")
	for {
		_, err := fmt.Scan(&length)
		if err == nil && length > 0 && length < 64 {
			return length
		}
		fmt.Print("Please, enter a valid whole number between 1 and 128: ")
	}
}

func getCharacters(text string) string {
	var answer string
	fmt.Printf("Do you want to include %s in your password? (Y | N): ", text)
	for {
		fmt.Scan(&answer)
		formattedAns := strings.ToLower(answer)
		if formattedAns == "y" || formattedAns == "n" {
			return formattedAns
		}
		fmt.Print("Please, enter either Y or N: ")
	}
}

func genPwd(upper, lower, nums, symbols []string) string {
	pwdLength := getPwdLen()
	includeUppercase := getCharacters("upper case characters")
	includeLowercase := getCharacters("lower case characters")
	includeNums := getCharacters("numbers")
	includeSymbols := getCharacters("symbols")

	var charsSelected []string
	if includeUppercase == "y" {
		charsSelected = append(charsSelected, upper...)
	}
	if includeLowercase == "y" {
		charsSelected = append(charsSelected, lower...)
	}
	if includeNums == "y" {
		charsSelected = append(charsSelected, nums...)
	}
	if includeSymbols == "y" {
		charsSelected = append(charsSelected, symbols...)
	}
	if len(charsSelected) == 0 {
		fmt.Println("Error: You must select at least one character type!")
		return "Error, could not generate password!"
	}
	var password []string
	rand.Shuffle(len(charsSelected), func(i, j int) { charsSelected[i], charsSelected[j] = charsSelected[j], charsSelected[i] })
	for i := 0; i < int(pwdLength); i++ {
		randNum := rand.Intn(len(charsSelected))
		password = append(password, charsSelected[randNum])
	}
	formattedPwd := strings.Join(password, "")

	return formattedPwd
}
