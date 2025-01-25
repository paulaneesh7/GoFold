package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
var courses []Course


// middleware - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == "" && c.CoursePrice == 0 && c.Author == nil
}

func main() {
	fmt.Println("Http Server In Golang")
}



// Controller - file


// sever home route
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home Page</h1>"))
}


func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}


func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	id := params["id"]

	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("No course found with id %v", id))
}


func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a Course")
	w.Header().Set("Content-Type", "application/json")


	// what if entire body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide input")
		return
	}

	// what about the data is send like this {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // decoding the data and storing it in course variable

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data found in the body")
		return
	}

	// generate a unique id (for course) and convert it into string (just generated it through autogeneration)
	course.CourseId = strconv.Itoa(len(courses) + 1)

	// append course into courses
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}