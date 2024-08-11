package notestruct

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

/* Constructor function*/
func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("title and content are required")
	}

	return Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}

/* Method of Note struct to display note */
func (n Note) Display(){
	fmt.Printf("Your note titled %v has the following content: %v\n", n.title, n.content)
}