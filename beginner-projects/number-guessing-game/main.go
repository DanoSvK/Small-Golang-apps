package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println(`
	Welcome to the number guessing game!
	You will try to guess a number between 1-100.
	After an incorrect guess, you will recieve a hint saying if the target number is higher or lower.
	You have 10 tries.
	Good luck!`)

	for {
		difficulty := choosingDifficulty()
		targetNumber := rand.Intn(100) + 1
		gameLogic(targetNumber, difficulty)
		if !restart() {
			break
		}
	}
	fmt.Println("Thanks for playing! Goodbye.")
}

func choosingDifficulty() int {
	var difficulty string
	fmt.Print("Before you start, choose a difficulty (easy - 10 lives or hard - 5 lives): ")

	for {
		fmt.Scan(&difficulty)
		difficulty = strings.ToLower(difficulty)

		if difficulty == "hard" {
			return 5
		} else if difficulty == "easy" {
			return 10
		}
		fmt.Print("Error: incorrect difficulty selected. Choose Easy or Hard: ")
	}
}

func playerInput() int {
	var value int
	fmt.Print("Please, select a number between 1-100: ")
	for {
		fmt.Scan(&value)
		if value < 1 || value > 100 {
			fmt.Print("Error: incorrect number selected. Select a number between 1-100: ")
		} else {
			return value
		}
	}
}

func gameLogic(target, lives int) {
	livesLeft := lives
	for {
		if livesLeft == 0 {
			fmt.Printf("You have %d lives left. Game over!\n", livesLeft)
			return
		}
		guessedNumber := playerInput()
		if guessedNumber > target {
			fmt.Printf("Incorrect, try lower. You have %d lives left.\n", livesLeft)
			livesLeft -= 1
		} else if guessedNumber < target {
			fmt.Printf("Incorrect, try higher. You have %d lives left.\n", livesLeft)
			livesLeft -= 1
		} else if guessedNumber == target {
			fmt.Println("You guessed it!")
			return
		}
	}
}

func restart() bool {
	var value string
	fmt.Print("Would you like to play again? Y/N: ")
	for {
		fmt.Scan(&value)
		value = strings.ToLower(value)
		if value == "y" {
			return true
		} else if value == "n" {
			return false
		} else {
			fmt.Print("Please, type in Y or N: ")
		}
	}
}
