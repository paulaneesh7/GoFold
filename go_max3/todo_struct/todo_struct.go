package todostruct

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json: "text"`
}

func (t Todo) Display() {
	fmt.Printf("Your Todo is: %v", t.Text)
}

func (t Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(t) // converting the struct to json

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("content is required")
	}

	return Todo{
		Text: content,
	}, nil
}