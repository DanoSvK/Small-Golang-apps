package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"unicode"
)

func main() {
	lives := 6
	randomWord, display := generateRandomWord()
	for {
		letter := guessALetter()
		isLetterInWord := checkWordForLetter(randomWord, letter)
		display = updateDisplay(randomWord, letter, display, isLetterInWord, &lives)
		fmt.Println(Stages[lives])
		fmt.Println(display)
		isGameOver := handleGameOver(display, lives)
		if isGameOver {
			break
		}
	}
	if lives > 0 {
		fmt.Println("You won!")
	} else {
		fmt.Print("You lost! The correct word was: ", randomWord)
	}
}

func guessALetter() string {
	var letter string
	fmt.Print("Guess a letter: ")
	for {
		_, err := fmt.Scan(&letter)
		if err != nil || len(letter) != 1 || !unicode.IsLetter(rune(letter[0])) {
			fmt.Print("Please, enter a valid letter: ")
			continue
		}
		return strings.ToLower(letter)
	}
}

func checkWordForLetter(randomWord string, letter string) bool {
	return strings.Contains(randomWord, letter)
}

func generateRandomWord() (string, []string) {
	wordList := []string{"aardvark", "baboon", "camel"}
	randomWord := wordList[rand.Intn(3)]

	var display []string
	for range randomWord {
		display = append(display, "_")
	}

	return randomWord, display
}

func updateDisplay(randomWord, letter string, display []string, isLetterInWord bool, lives *int) []string {
	if isLetterInWord {
		for i, v := range randomWord {
			if string(v) == letter {
				display[i] = letter
			}
		}
	} else {
		*lives -= 1
	}
	return display
}

func handleGameOver(display []string, lives int) bool {
	return !slices.Contains(display, "_") || lives == 0
}
