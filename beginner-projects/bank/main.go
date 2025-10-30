package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var choice int
	var deposit float64
	var withdraw float64

	var balance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
		// panic("Cannot continue without balance")
	}

	// Switch is not good in combination with for as break within switch does not break the loop
	for {
		welcomeMsg()
		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("Your balance is %.2f\n", balance)
		case 2:
			fmt.Printf("Enter the amount you want to deposit: ")
			fmt.Scan(&deposit)
			if deposit >= 0 {
				fmt.Println("Invalid amount")
			}
			balance += deposit
			writeToFile(balance)
			fmt.Printf("\nYou succesfully deposited %.2f. Your currecnt balance is %.2f\n", deposit, balance)
		case 3:
			fmt.Println("Withdraw money:")
			fmt.Scan(&withdraw)
			if withdraw > balance {
				fmt.Println("Insufficient balance")
			}
			balance -= withdraw
			writeToFile(balance)
			fmt.Printf("\nYou succesfully withdrew %.2f. Your currecnt balance is %.2f\n", withdraw, balance)
		default:
			fmt.Println("Thank you for using Go Bank")
			return
		}
	}
}

const accountBalanceFile = "balance.txt"

func writeToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Error reading file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Error parsing balance")
	}
	return balance, nil
}

func welcomeMsg() {
	fmt.Println("Welcome to Go Bank")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
