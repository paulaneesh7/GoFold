package main

import (
	"fmt"
)

// Model for courses - file

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	Author      *Author `json:"author"` // we specified a pointer here because we don't want to create a copy of the Author: Firstname Lastname when we add data to it later
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}


// fake DB
var course []Course


// middleware - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == "" && c.CoursePrice == 0 && c.Author == nil
}

func main() {
	fmt.Println("Http Server In Golang")
}
