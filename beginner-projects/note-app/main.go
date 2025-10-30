package main

import (
	"bufio"
	"fmt"
	"note-app/note"
	"note-app/todo"
	"os"
	"strings"
)

type outputtable interface {
	Save() error
	Display()
}

func main() {
	title, content := getNoteData()
	text := getUserInput("Enter the todo text: ")

	note, err := note.NewNote(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.NewTodo(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(note)
	outputData(todo)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data outputtable) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving data failed: ", err)
		return err
	}
	fmt.Println("Data saved successfully")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the note title: ")
	content := getUserInput("Enter the note content: ")

	return title, content
}

func test[T int | float64 | string](a, b T) T {
	return a + b
}
