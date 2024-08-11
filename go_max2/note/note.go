package note

import (
	"fmt"
	notestruct "go_max2/note_struct"
)

func NoteCreate() {
	
	title, content := getNoteData()

	userNote, err := notestruct.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	userNote.Display()

}

func getNoteData() (string, string) {
	title := getUserInput("Enter Note title: ")
	

	content := getUserInput("Enter Note content: ")
	

	return title, content
}



func getUserInput(prompt string) string {
	fmt.Println(prompt)

	var value string
	fmt.Scanln(&value)

	

	return value
}