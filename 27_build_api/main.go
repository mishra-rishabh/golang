package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

/* Model for courses */
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

/* Fake DB */
var courses []Course

/* Middleware or Helper functions => should be in a separate file */
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

/* Controllers => should be in a separate file */
// Serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API building tutorial in golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through the courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)

			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id!")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")

	w.Header().Set("Content-Type:", "application/json")

	// what if entire body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data!")

	}

	// what if data is {}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON!")

		return
	}

	// challenges: generate unique id, convert this id to string
	// append course into courses
	rand.NewSource(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")

	w.Header().Set("Content-Type", "application/json")

	// 1. grab id from req
	params := mux.Vars(r)

	// loop through the value, once we get the id, remove it, and then add the value again in courses with my Id (from params)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var courseVar Course
			_ = json.NewDecoder(r.Body).Decode(&courseVar)

			courseVar.CourseId = params["id"]
			courses = append(courses, courseVar)

			json.NewEncoder(w).Encode(course)

			return
		}
	}

	// also send a response when id is not found (not taken care of here)
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop through the value, find the id and remove it
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}

func main() {
	fmt.Println("Building API in golang")

	router := mux.NewRouter()

	// seeding
	courses = append(courses,
		Course{
			CourseId: "2", CourseName: "Golang", CoursePrice: 299, Author: &Author{
				FullName: "Rishabh Mishra", Website: "rishabh.go.dev"}},
	)

	courses = append(courses,
		Course{
			CourseId: "3", CourseName: "Javascript", CoursePrice: 199, Author: &Author{
				FullName: "Rishabh Mishra", Website: "rishabh.go.dev"}},
	)

	// routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", router))
}
