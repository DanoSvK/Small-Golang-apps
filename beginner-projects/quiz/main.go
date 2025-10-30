package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type questions struct {
	Type              string
	Difficulty        string
	Category          string
	Question          string
	Correct_answer    string
	Incorrect_answers []string
}

func main() {
	data, err := loadQuestions("questions.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	quizStart(data)
	fmt.Println("See you next time!")
}

func loadQuestions(path string) ([]questions, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Could not read file: %w", err)
	}

	var questions []questions
	err = json.Unmarshal(file, &questions)
	if err != nil {
		return nil, fmt.Errorf("Issue processing file data: %w", err)
	}
	return questions, nil
}

func quizStart(data []questions) {
	score := 0
	lastQuestion := false
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	for i, v := range data {
		var answer string
		fmt.Printf("Question %d: %s\n", i+1, v.Question)
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Select your answer (true/false): ")
			scanner.Scan()
			answer = scanner.Text()
			if len(answer) == 0 {
				fmt.Println("Input cannot be empty.")
				continue
			}
			answer = strings.ToUpper(answer[:1]) + strings.ToLower(answer[1:])
			if answer == "True" || answer == "False" {
				break
			} else {
				fmt.Println("Invalid input. Please enter 'true' or 'false'.")
				continue
			}
		}
		if answer == v.Correct_answer {
			score += 1
			fmt.Println("Correct. Your current score is:", score)
		} else {
			fmt.Println("Wrong. Your current score is:", score)
		}
		if lastQuestion {
			fmt.Println("The quiz is over. Your final score is:", score)
			break
		}
	}
}
