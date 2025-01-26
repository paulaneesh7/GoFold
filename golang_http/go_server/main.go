package main

import (
	"encoding/json"
	"fmt"
	"log"
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

// PORT
const PORT = 8080


// middleware - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == "" && c.CoursePrice == 0 && c.Author == nil
}

func main() {
	fmt.Printf("Http Server In Golang running on port %v\n", PORT)

	r := mux.NewRouter()

	// routes
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")
	

	// seeding some dummy data into courses
	courses = append(courses, Course{"1", "React", 299, &Author{"Aneesh", "https://paulaneesh7.in"}})
	courses = append(courses, Course{"2", "Angular", 199, &Author{"Kunal", "https://paulaneesh7.in"}})
	courses = append(courses, Course{"3", "Golang", 199, &Author{"Aneesh", "https://paulaneesh7.in"}})
	courses = append(courses, Course{"4", "Python", 199, &Author{"Kunal", "https://paulaneesh7.in"}})

	// listening to a port
	log.Fatal(http.ListenAndServe(":8080", r))

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


func getCourseById(w http.ResponseWriter, r *http.Request) {
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

	// what if course with same name already exists
	for _, course := range courses {
		if course.CourseName == course.CourseName {
			json.NewEncoder(w).Encode(fmt.Sprintf("Course with name %v already exists", course.CourseName))
			return
		}
	}

	// generate a unique id (for course) and convert it into string (just generated it through autogeneration)
	course.CourseId = strconv.Itoa(len(courses) + 1)

	// append course into courses
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}


func updateCourse(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// extracting the id from params
	params := mux.Vars(r)
	id := params["id"]

	// decode the request body
	var updatedCourse Course
	_ = json.NewDecoder(r.Body).Decode(&updatedCourse)

	for index, course := range courses {
		if course.CourseId == id {
			courses[index] = updatedCourse
			courses[index].CourseId = id
			json.NewEncoder(w).Encode(courses[index])
			return
		}
	}


	// Send a reponse when id is not found
	json.NewEncoder(w).Encode(fmt.Sprintf("No course found with id %v", id))
}


func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")

	// extracting the id from params
	params := mux.Vars(r)
	id := params["id"]

	for index, course := range courses {
		if course.CourseId == id {
			courses = append(courses[:index], courses[index+1:]...)  // based on index, delete the course
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("No course found with id %v", id))
}
