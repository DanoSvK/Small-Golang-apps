package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type HabitsData struct {
	Name           string   `json:"name"`
	DatesCompleted []string `json:"dates_completed"`
	Goal           int      `json:"goal"`
}

type Habits struct {
	Habits []HabitsData `json:"habits"`
}

func main() {
	data, err := loadDataFromFile()
	if err != nil {
		fmt.Println("Error loading data: ", err)
	}
	fmt.Println(data.Habits[0].DatesCompleted)
	name := habitName("Please, enter the name of the habit: ")
	var isHabit bool
	var date string
	var goal int
	for i := range data.Habits {
		nameToLower := strings.ToLower(data.Habits[i].Name)
		if nameToLower == name {
			isHabit = true
			break
		} else {
			isHabit = false
		}
	}
	if isHabit {
		date = habitDate("Please, enter the date of the habit completion: ")
	} else {
		date = habitDate("Please, enter the date of the habit completion in YYYY-MM-DD format: ")
		goal = habitGoal("Please, enter the goal in number of days: ")
	}
	saveDataToHabits(name, date, goal)
}

func habitName(text string) string {
	var input string
	var err error
	fmt.Print(text)
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil || strings.TrimSpace(input) == "" {
			fmt.Print("Invalid name! Please, enter a valid name (must not be empty): ")
			continue
		}
		fmt.Println("habit name for loop:", input)
		break
	}
	input = strings.TrimSpace(strings.ToLower(input))

	return input
}

func habitDate(text string) string {
	var input string
	var err error
	fmt.Print(text)
	for {
		fmt.Scan(&input)
		_, err = time.Parse("2006-01-02", input)
		if err != nil {
			fmt.Print("Invalid date format! Please use the format YYYY-MM-DD: ")
			continue
		}
		break
	}
	return input
}

func habitGoal(text string) int {
	var input string
	var num int
	fmt.Print(text)
	for {
		fmt.Scan(&input)
		var err error
		num, err = strconv.Atoi(input)
		if err != nil {
			fmt.Print("Invalid number! Please, enter a valid integer: ")
			continue
		}
		break
	}
	return num
}

func loadDataFromFile() (Habits, error) {
	data, err := os.ReadFile("habits.json")
	if err != nil {
		return Habits{}, fmt.Errorf("Unable to read data: %w", err)
	}

	var habits Habits
	err = json.Unmarshal(data, &habits)
	if err != nil {
		return Habits{}, fmt.Errorf("Unable to unmarshal data: %w", err)
	}

	return habits, nil
}

func writeDataToFile(data Habits) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("Could not marshal data: %w", err)
	}

	os.WriteFile("habits.json", jsonData, 0644)

	return nil
}

func saveDataToHabits(name, date string, goal int) {
	data, err := loadDataFromFile()
	if err != nil {
		fmt.Println("Error loading data: ", err)
	}

	for i := range data.Habits {
		if strings.ToLower(data.Habits[i].Name) == name {
			data.Habits[i].DatesCompleted = append(data.Habits[i].DatesCompleted, date)
			writeDataToFile(data)
			return
		}
	}

	habit := HabitsData{
		Name:           name,
		DatesCompleted: []string{date},
		Goal:           goal,
	}
	data.Habits = append(data.Habits, habit)
	writeDataToFile(data)
}
