package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Welcome to temperature converter app")
	temp := getTemperatureValue("Please, enter the temperature: ")
	originalScale := getTemperatureScale("Please, enter the original temperature scale (C | F | K): ")
	targetScale := getTemperatureScale("Please, enter the target temperature scale (C | F | K): ")

	result := convertTemperature(temp, originalScale, targetScale)
	formattedTemp := formatNumber(temp)
	formattedResult := formatNumber(result)

	fmt.Printf("%s°%s converted to %s is %s°%s", formattedTemp, originalScale, targetScale, formattedResult, targetScale)
}

func getTemperatureValue(text string) float64 {
	var value float64
	for {
		fmt.Print(text)
		_, err := fmt.Scan(&value)
		if err == nil {
			return value
		}
		fmt.Println("Invalid input. Please enter a valid number.")
	}
}

func getTemperatureScale(text string) string {
	var value string
	for {
		fmt.Print(text)
		fmt.Scan(&value)
		value = strings.ToUpper(value)
		if value != "C" && value != "F" && value != "K" {
			continue
		}
		break
	}

	return value
}

func convertTemperature(temp float64, source, target string) float64 {
	if source == target {
		return temp
	}
	var calcs = map[string]float64{
		"CtoF": (temp * 9 / 5) + 32,
		"CtoK": temp + 273.15,
		"FtoC": (temp - 32) * 5 / 9,
		"FtoK": (temp-32)*5/9 + 273.15,
		"KtoC": temp - 273.15,
		"KtoF": (temp-273.15)*9/5 + 32,
	}
	return calcs[source+"to"+target]
}

func formatNumber(n float64) string {
	if math.Mod(n, 1) == 0 {
		return fmt.Sprintf("%.0f", n)
	}
	return fmt.Sprintf("%.2f", n)
}
