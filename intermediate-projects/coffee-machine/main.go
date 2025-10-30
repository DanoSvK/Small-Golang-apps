package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type ingredients struct {
	Water  float64 `json:"water"`
	Milk   float64 `json:"milk"`
	Coffee float64 `json:"coffee"`
}

type drink struct {
	Ingredients ingredients `json:"ingredients"`
	Cost        float64     `json:"cost"`
}

type menu struct {
	Menu      map[string]drink   `json:"menu"`
	Resources map[string]float64 `json:"resources"`
}

func main() {
	data, err := loadData("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	// if _, exists := data.Resources["money"]; !exists {
	// 	data.Resources["money"] = 0
	// }

	for {
		customerChoice := selectDrink("\nWhat would you like? (espresso/latte/cappuccino): ")
		if customerChoice == "off" {
			return
		}
		if customerChoice == "report" {
			fmt.Println(data.Resources)
			continue
		}
		if customerChoice == "refill" {
			water := refillResources("Enter amount of water for refil in ml: ")
			milk := refillResources("Enter amount of milk for refil in ml: ")
			coffee := refillResources("Enter amount of coffee for refil in g: ")
			addToData("data.json", water, milk, coffee)
			continue
		}

		isSufficientIng := checkIngredients(&data, customerChoice)
		if !isSufficientIng {
			continue
		}

		isEnoughMoney := processCoins(&data, customerChoice)
		if !isEnoughMoney {
			continue
		}
		fmt.Printf("Here is your %s. Enjoy!", customerChoice)

		err = writeData("data.json", data)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}
}

// Prompts
func selectDrink(text string) string {
	var input string
	fmt.Print(text)
	for {
		fmt.Scan(&input)
		input = strings.TrimSpace(strings.ToLower(input))
		validInputs := map[string]bool{"espresso": true, "latte": true, "cappuccino": true, "report": true, "refill": true, "off": true}
		if !validInputs[input] {
			fmt.Print("Please, select a valid input (espresso/latter/cappuccino): ")
			continue
		}
		return input
	}
}

func insertCoins(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func refillResources(text string) float64 {
	var amount float64
	fmt.Print(text)
	for {
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Print("Please, enter a valid number: ")
			continue
		}
		return amount
	}

}

func loadData(path string) (menu, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return menu{}, errors.New("could not open file")
	}

	var menuData menu
	err = json.Unmarshal(file, &menuData)
	if err != nil {
		return menu{}, errors.New("something went wrong")
	}

	return menuData, nil
}

// Update coffee machine contents after each coffee dispensed
func writeData(path string, data menu) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return errors.New("could not parse JSON data")
	}
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return errors.New("could not write data to file")
	}
	return nil
}

// Refill resources
func addToData(path string, water, milk, coffee float64) error {
	data, err := loadData(path)
	if err != nil {
		return err
	}
	data.Resources["water"] += water
	data.Resources["milk"] += milk
	data.Resources["coffee"] += coffee

	err = writeData(path, data)
	if err != nil {
		return err
	}
	return nil
}

// Main logic
func checkIngredients(data *menu, choice string) bool {
	ing := data.Menu[choice].Ingredients
	pass := true

	if ing.Coffee >= data.Resources["coffee"] {
		fmt.Println("Sorry, there is not enough coffee")
		pass = false
	} else {
		data.Resources["coffee"] -= ing.Coffee
	}
	if ing.Milk >= data.Resources["milk"] {
		fmt.Println("Sorry, there is not enough milk")
		pass = false
	} else {
		data.Resources["milk"] -= ing.Milk
	}
	if ing.Water >= data.Resources["water"] {
		fmt.Println("Sorry, there is not enough water")
	} else {
		data.Resources["water"] -= ing.Water
	}

	return pass
}

func processCoins(data *menu, choice string) bool {
	quarters := insertCoins("How many quarters?: ")
	dimes := insertCoins("How many dimes?: ")
	nickels := insertCoins("How many nickels?: ")
	pennies := insertCoins("How many pennies?: ")

	totalAmount := 0.25*quarters + 0.1*dimes + 0.05*nickels + 0.01*pennies

	if totalAmount > data.Menu[choice].Cost {
		change := totalAmount - data.Menu[choice].Cost
		(*data).Resources["money"] += data.Menu[choice].Cost
		fmt.Printf("Here is $%.2f back as change.", change)
		fmt.Println(data)
		return true
	} else {
		fmt.Printf("Sorry, that's not enough money. $%.2f refunded.", totalAmount)
		return false
	}
}
