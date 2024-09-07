package todo

import (
	"bufio"
	"fmt"
	"go_max3/note"
	todostruct "go_max3/todo_struct"
	"os"
	"strings"
)

// Any struct that implements the saver interface must have a Saver method and Returns an error.
type saver interface{
	Save() error
}


// Embedding saver interface and Display method together in the outputtable interface.
type outputtable interface{
	saver
	Display()
}


func CreateTodo() {
	title, content := getNoteData()

	todo, err := todostruct.New(content)

	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)

}


func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

// Using the outputtable interface as a type to "data" parameter
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// Using the saver interface as a type to "data" parameter
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}