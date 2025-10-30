package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
)

type searchTerm struct {
	Name           string
	Follower_count int
	Description    string
	Country        string
}

func main() {
	data, err := loadData("data.json")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	selectedTerms := generateSearchTerms(&data, 2)
	score := 0

	for {
		displayUi(selectedTerms)
		playerGuess := higherOrLower()
		correctAnswer := compare(selectedTerms)
		clearTerminal()
		if handleGameState(playerGuess, correctAnswer, &score) {
			return
		}
		newSearchTerm(&selectedTerms, &data)
	}
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

// Display UI
func displayUi(selectedTerms []searchTerm) {
	fmt.Println(Logo)
	fmt.Printf("Compare A: %s, %s, from %s", selectedTerms[0].Name, selectedTerms[0].Description, selectedTerms[0].Country)
	fmt.Println(Vs)
	fmt.Printf("Against B: %s, %s, from %s\n", selectedTerms[1].Name, selectedTerms[1].Description, selectedTerms[1].Country)
}

// Load json data
func loadData(path string) ([]searchTerm, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("could not read the file")
	}
	var terms []searchTerm
	err = json.Unmarshal(file, &terms)
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	return terms, nil
}

// Generate 2 random numbers to fetch 2 random search terms
func generateSearchTerms(data *[]searchTerm, n int) []searchTerm {
	var searchTerms = make([]searchTerm, 0, n)
	for i := 0; i < n; i++ {
		randomIndex := rand.Intn(len(*data))
		randomItem := (*data)[randomIndex]
		searchTerms = append(searchTerms, randomItem)
		*data = slices.Delete(*data, randomIndex, randomIndex+1)
	}
	return searchTerms
}

// Ask player if the second has higher or lower count of followers
func higherOrLower() int {
	var value string
	fmt.Print("\nWho has more followers? Type A or B: ")

	for {
		fmt.Scan(&value)
		valueLower := strings.ToLower(value)
		if valueLower == "a" {
			return 0
		} else if valueLower == "b" {
			return 1
		}
		fmt.Print("Please, type A or B: ")
		continue
	}
}

// Compare follower counts
func compare(searchTerm []searchTerm) int {
	followerCountA := searchTerm[0].Follower_count
	followerCountB := searchTerm[1].Follower_count

	if followerCountA > followerCountB {
		return 0
	} else {
		return 1
	}
}

// Update score
func handleGameState(guess, answer int, score *int) bool {
	if guess == answer {
		*score += 1
		fmt.Println("Correct! Your current score is: ", *score)
		return false
	} else {
		fmt.Println("You lost. Your final score is: ", *score)
		return true
	}
}

// Keep the second search term for comparison and generate a new random one
func newSearchTerm(selectedTerms *[]searchTerm, data *[]searchTerm) {
	if len(*data) == 0 {
		fmt.Println("No more terms left. You win!")
		os.Exit(0) // Safe exit
	}
	randomIndex := rand.Intn(len(*data))
	(*selectedTerms)[0] = (*selectedTerms)[1]
	(*selectedTerms)[1] = (*data)[randomIndex]
	*data = slices.Delete(*data, randomIndex, randomIndex+1)
}
