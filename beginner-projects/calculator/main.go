package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome in calculator app")
	operation := selectOperation()
	num1, num2 := selectNumbers()
	result := handleOperation(operation, num1, num2)
	// resultStr := fmt.Sprintf("%f", result)
	var resultStr string
	resultStr = strconv.FormatFloat(result, 'f', -1, 64)

	for {
		var cont string
		fmt.Printf("The result is: %s. Do you want to continue? Y/N: ", resultStr)
		fmt.Scan(&cont)
		if cont == "y" {
			resultStr = contCalc(resultStr)
		}
		if cont == "n" {
			return
		}
	}

}

func selectOperation() string {
	var operation string
	fmt.Print("Select an operation ('+' | '-' | '*' | '/'): ")
	fmt.Scan(&operation)
	return operation
}

func handleOperation(operator string, num1, num2 float64) float64 {
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	return result
}

func selectNumbers() (float64, float64) {
	var num1 float64
	var num2 float64
	fmt.Print("Select the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Select the second number: ")
	fmt.Scan(&num2)
	return num1, num2
}

func contCalc(firstNumStr string) string {
	result, err := strconv.ParseFloat(firstNumStr, 64)
	if err != nil {
		fmt.Println("Error occured when parsing float")
		return ""
	}
	operation := selectOperation()
	var num2 float64
	fmt.Print("Select the second number: ")
	fmt.Scan(&num2)
	result = handleOperation(operation, result, num2)
	resultStr := strconv.FormatFloat(result, 'f', -1, 64)
	return resultStr
}
