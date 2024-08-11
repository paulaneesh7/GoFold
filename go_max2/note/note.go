package note

import (
	"bufio"
	"fmt"
	notestruct "go_max2/note_struct"
	"os"
	"strings"
)

func NoteCreate() {
	
	title, content := getNoteData()

	userNote, err := notestruct.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

}

func getNoteData() (string, string) {
	title := getUserInput("Enter Note title: ")
	

	content := getUserInput("Enter Note content: ")
	

	return title, content
}



func getUserInput(prompt string) string {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin) // creating a reader that listens on command line input

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	

	return text
}