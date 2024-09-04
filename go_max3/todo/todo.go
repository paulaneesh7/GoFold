package todo

import (
	"bufio"
	"fmt"
	todostruct "go_max3/todo_struct"
	"os"
	"strings"
)

func CreateTodo() {
	content := getTodoData()

	todo, err := todostruct.New(content)

	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	todo.Display()

	err = todo.Save()

	if err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}
}

func getTodoData() string {

	content := getUserInput("Enter Todo content: ")

	return content
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