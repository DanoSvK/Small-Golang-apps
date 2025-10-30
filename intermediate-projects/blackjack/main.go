package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	dealerCards := generatecards(1)
	playerCards := generatecards(2)
	displayCards(dealerCards, playerCards)
	playerScore := cardCounter(playerCards)

	if playerScore == 0 {
		dealerCards = append(dealerCards, generatecards(1)...)
		fmt.Println(compare(playerCards, dealerCards))
		return
	}

	isPlayerTurn := hitOrStand()
	for isPlayerTurn {
		playerCards = append(playerCards, generatecards(1)...)
		displayCards(dealerCards, playerCards)
		playerScore = cardCounter(playerCards)
		if playerScore > 21 {
			fmt.Println(compare(playerCards, dealerCards))
			return
		} else if playerScore == 21 {
			return
		}
		compare(playerCards, dealerCards)
		isPlayerTurn = hitOrStand()
	}

	for cardCounter(dealerCards) < 17 {
		dealerCards = append(dealerCards, generatecards(1)...)
	}
	displayCards(dealerCards, playerCards)
	fmt.Println(compare(playerCards, dealerCards))
}

// Helper functions
func generatecards(iterations int) []int {
	cards := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
	selectedCards := make([]int, 0, iterations)
	for i := 0; i < iterations; i++ {
		index := rand.Intn(len(cards))
		selectedCards = append(selectedCards, cards[index])
	}
	return selectedCards
}

func cardCounter(cards []int) int {
	score, aceCount := 0, 0
	for _, v := range cards {
		if v == 11 { // Track Aces properly
			aceCount++
		}
		score += v
	}

	// Convert Aces from 11 to 1 until the score is safe
	for score > 21 && aceCount > 0 {
		score -= 10
		aceCount--
	}

	if score == 21 && len(cards) == 2 {
		return 0
	}
	return score
}

func displayCards(dealerCards []int, playerCards []int) {
	fmt.Println("\nCurrent turn")
	fmt.Println("Dealer's cards: ", dealerCards)
	fmt.Println("Your cards: ", playerCards)
	fmt.Println("_________________")
}

// Main logic functions
func hitOrStand() bool {
	var value string
	fmt.Print("Do you wish to Hit or Stand?: ")

	for {
		fmt.Scan(&value)
		valueToLower := strings.ToLower(value)
		if valueToLower == "hit" {
			return true
		} else if valueToLower == "stand" {
			return false
		} else {
			fmt.Print("Please, enter hit or stand: ")
			continue
		}
	}
}

func compare(playerCards []int, dealerCards []int) string {
	playerScore := cardCounter(playerCards)
	dealerScore := cardCounter(dealerCards)

	if playerScore == 0 && dealerScore == 0 {
		return "You both have Blackjack. Push!"
	} else if playerScore == 0 {
		return "You have Blackjack! You win!"
	} else if dealerScore == 0 {
		return "Dealer has Blackjack! You lose!"
	} else if playerScore > 21 {
		return "You lose!"
	} else if dealerScore > 21 {
		return "You win!"
	} else if playerScore == dealerScore {
		return "Push!"
	} else if playerScore > dealerScore {
		return "You win!"
	} else {
		return "You lose!"
	}
}
