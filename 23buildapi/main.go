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

type Course struct {
	CourseID   string  `json:"courseID"`
	CourseName string  `json:"courseName"`
	CourseFee  float64 `json:"courseFee"`
	Author     *Author `json:"author"`
}

type Author struct {
	AuthorName string `json:"authorName"`
	AuthorSite string `json:"authorSite"`
}

var courses []Course

func main() {

	fmt.Println("welcome to build api tutorial...")

	router := mux.NewRouter()

	router.HandleFunc("/", getHome).Methods("GET")
	router.HandleFunc("/course/all", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course/create", createOneCourse).Methods("POST")
	router.HandleFunc("/course/update/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/delete/{id}", deleteOneCourse).Methods("DELETE")

	//keep the server up and running always

	//http.ListenAndServe(":8080", nil)

	courses = append(courses, Course{
		CourseID:   "123",
		CourseName: "Java Springboot course",
		CourseFee:  12000,
		Author: &Author{
			AuthorName: "Ajay Kumar",
			AuthorSite: "https://www.google.com/search?q=ajay+kumar",
		},
	})

	courses = append(courses, Course{
		CourseID:   "223",
		CourseName: "Golang course",
		CourseFee:  10000,
		Author: &Author{
			AuthorName: "Raghu Kumar",
			AuthorSite: "https://www.google.com/search?q=raghu+kumar",
		},
	})

	log.Fatal(http.ListenAndServe(":8080", router))

}

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

//controllers

func getHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside getHome method")
	w.Write([]byte(`"<h1> Welcome to the home of apis tutorial </h1>"`))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")
	//w.Write()
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//get course id from url
	params := mux.Vars(r)

	//loop throug the course and find that course

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No courses found...")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating new course")
	w.Header().Set("content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Kindly provide a valid Course object - nil body request")
		return
	}

	var course Course

	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Kindly provide a valid Course Object -Empty request")
		return
	}
	//create a new course id based on the random number
	rand.Seed(time.Now().UnixMilli())
	randomNumber := rand.Intn(100)

	randomNumberString := strconv.Itoa(randomNumber)
	course.CourseID = randomNumberString

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating existing course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var course Course

	if r.Body == nil {
		json.NewEncoder(w).Encode("Kindly provide a valid Course Object to Update")
	}
	json.NewDecoder(r.Body).Decode(&course)

	for index, courseObject := range courses {

		if courseObject.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
		}
	}

	course.CourseID = params["id"]
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course is deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("Course is not Avilable into the database")
}
