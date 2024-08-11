package notestruct

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

/* Constructor function*/
func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("title and content are required")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

/* Method of Note struct to display note */
func (n Note) Display(){
	fmt.Printf("Your note **Titled** %v has the following **Content**: %v\n", n.Title, n.Content)
}


/* Save the notes to a file */
func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	/* json is already in byte[] format, if file is saved then fine but if not then "Error" will be returned */
	return os.WriteFile(fileName, json, 0644)
}
