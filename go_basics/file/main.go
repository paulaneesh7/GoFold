package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	

	// create a file
	createFile()


	// read a file
	readFile()
}

func createFile(){
	// create a file -> it returns a file pointer and an error
	file, err := os.Create("example.txt")
	if err != nil{
		fmt.Println("Error while creating a file: ", err)
		return
	}
	defer file.Close() // this will close the file after the main function ends, it's like sc.close() of "Java"

	fmt.Println("Successfully created a file")

	// Creating content in a file
	createContentInFile(file)
}

func createContentInFile(file *os.File){
	// Writing content to  a file
	content := "My name is Aneesh Paul"
	byte, errors := io.WriteString(file, content+"\n") // this returns the number of bytes written and an error
	fmt.Println("Byte written: ", byte)
	if errors != nil{
		fmt.Println("Error while writing to a file: ", errors)
		return
	}

}

func readFile(){
	file, err := os.Open("example.txt")
	if err != nil{
		fmt.Println("Error while opening a file: ", err)
		return
	}

	defer file.Close()

	// to read the data we need to create a buffer (which is a temporary storage area)
	buffer := make([]byte, 1024) // created a slice of size 1024

	// read the content of the file
	for {
		cnt, error := file.Read(buffer)
		if error == io.EOF{ // if while reading the it reaches till the end, then just break it
			break
		}

		if error != nil{
			fmt.Println("Error while reading a file: ", error)
			return
		}

		// printing the content of the file from 0 till the number of bytes read
		fmt.Println("Content of the file: ", string(buffer[:cnt]))
	}

	// shortcut to read a file
	shortCutToReadFile()
}
/*
Whatever content is written/present inside the file, it pushes them in the buffer in small chunks (as buffer
is a temporary kind of memory) and this keeps going till it reaches the end of the file.
The moment it reaches the end of the file, it returns an error called io.EOF (End of File) and then it breaks
Then we read the buffer from 0 till the end (whatever content is present in the buffer)

Check the notes.pdf file for more detail
*/

func shortCutToReadFile(){
	content, error := ioutil.ReadFile("example.txt") // it's deprecated, but still can be used
	// use os.ReadFile() instead of ioutil.ReadFile()

	if error != nil{
		fmt.Println("Error while reading a file: ", error)
		return
	}

	fmt.Println("Content of the file: ", string(content))
}

func deleteFile(){
	err := os.Remove("example.txt")
	if err != nil{
		fmt.Println("Error while deleting a file: ", err)
		return
	}
	fmt.Println("Successfully deleted a file")
}