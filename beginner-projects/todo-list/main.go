package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type task struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	selectedAction := selectAction()

	switch selectedAction {
	case "display":
		err := displayTasks()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	case "delete":
		err := deleteTask()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	case "input":
		title := newTask("Please, enter the Title of the task: ")
		description := newTask("Please, enter the Description of the task: ")
		id := fmt.Sprintf("%d", time.Now().UnixNano())

		newTaskToSave, err := saveTask(id, title, description, time.Now())

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		saveDataToFile(newTaskToSave)
	}
}

// Helper functions
func loadDataFromFile() ([]task, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		return nil, errors.New("failed to open tasks.json: file not found or inaccessible")
	}
	defer file.Close()
	var tasks []task

	// Check if file is empty (if it is, NewDecoder() would fail)
	// Stat() method fetches file metadata (like size of the file)
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, errors.New("could not get file info")
	}
	// Size() method fetches file size
	if fileInfo.Size() == 0 {
		return []task{}, nil
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, errors.New("failed to decode tasks.json: invalid JSON format or empty file")
	}
	return tasks, nil
}

func writeDataToFile(tasks []task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return errors.New("error occured when creating file")
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(&tasks)
	if err != nil {
		return errors.New("error occured when encoding data")
	}
	return nil
}

func saveDataToFile(newTaskToSave task) error {
	tasks, err := loadDataFromFile()
	if err != nil {
		fmt.Println(err)
		return err
	}
	tasks = append(tasks, newTaskToSave)

	err = writeDataToFile(tasks)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return nil
}

// Main functions
func selectAction() string {
	var validActions = map[string]bool{
		"delete":  true,
		"input":   true,
		"display": true,
	}

	for {
		fmt.Println("Select the desired action (Input, Delete, Display): ")
		var value string
		fmt.Scan(&value)
		value = strings.ToLower(strings.TrimSpace(value))

		// Consume the leftover newline
		bufio.NewReader(os.Stdin).ReadString('\n')

		if validActions[value] {
			return value
		}
		fmt.Println("Invalid action. Please choose one of the following: Input, Delete, Display.")
	}
}

func displayTasks() error {
	tasks, err := loadDataFromFile()
	if err != nil {
		return err
	}
	fmt.Println(tasks)
	return nil
}

func deleteTask() error {
	fmt.Print("Type title of a task that you wish to delete: ")
	reader := bufio.NewReader(os.Stdin)

	taskTitle, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Could not read input")
	}
	taskTitle = strings.ToLower(strings.TrimSpace(taskTitle))

	tasks, err := loadDataFromFile()
	if err != nil {
		return err
	}

	for i, v := range tasks {
		if strings.ToLower(v.Title) == taskTitle {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	err = writeDataToFile(tasks)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return nil
}

func newTask(text string) string {
	var value string
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func saveTask(id, title, description string, createdAt time.Time) (task, error) {
	if title == "" || description == "" {
		fmt.Println("Title and Description cannot be empty")
		return task{}, errors.New("Title and Description cannot be empty")
	}
	return task{
		Id:          id,
		Title:       title,
		Description: description,
		CreatedAt:   createdAt,
	}, nil
}

// func findHighestId() (int, error) {
// 	tasks, err := loadDataFromFile()
// 	if err != nil {
// 		return 0, err
// 	}

// 	highestId := 0
// 	for _, v := range tasks {
// 		if v.Id > highestId {
// 			highestId = v.Id
// 		}
// 	}

// 	return highestId + 1, nil
// }
