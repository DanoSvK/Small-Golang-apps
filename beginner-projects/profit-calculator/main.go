package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, please follow instructions to proceed.")

	revenue, err1 := getUserInput("revenue")

	// if err != nil {
	// 	fmt.Println("Invalid input")
	// 	fmt.Println(err)
	// 	return
	// }

	expenses, err2 := getUserInput("expenses")

	// if err != nil {
	// 	fmt.Println("Invalid input")
	// 	fmt.Println(err)
	// 	return
	// }

	taxRate, err3 := getUserInput("taxRate")

	// if err != nil {
	// 	fmt.Println("Invalid input")
	// 	fmt.Println(err)
	// 	return
	// panic(err)
	// }

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	var ebt, profit, ratio = calculate(revenue, expenses, taxRate)
	storeResults(ebt, profit, ratio)
	fmt.Printf(`Earnings before tax: %.2f
				Earnings after tax: %.2f
				Ratio: %.2f`, ebt, profit, ratio)

}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	var ebt = revenue - expenses
	var profit = ebt - expenses/100*taxRate
	var ratio = ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Printf("Enter %s:", infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid input")
	}

	return userInput, nil
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("Earnings before tax: %.2f\nEarnings after tax: %.2f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
