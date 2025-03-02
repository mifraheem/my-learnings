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
	CourseId    string  `json:courseid`
	CourseName  string  `json:coursename`
	CoursePrice int     `json:courseprice`
	Author      *Author `json:author`
}

type Author struct {
	Fullname string `json:fullname`
	Email    string `json:email`
	Website  string `json:website`
}

var courses []Course

// middlewares, helpers -> in file
func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {
	fmt.Println("Welcome to our course platform")
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")
	r.HandleFunc("/courses/{id}", getCourse).Methods("GET")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers -> in file

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Our Courses Platform</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all courses")
	// w.Write([]byte("All Courses"))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting Single Course")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found")
	return

}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating New Course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please submit the actual data")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}
	for _, crs := range courses {
		if crs.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("This course already exists")
			return
		}
	}
	// generate the unique id

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating an existing course")
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses := append(courses[:index], courses[index+1:]...)
			var course Course
			fmt.Println(courses)
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting an existing course")
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("The course has been deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found")
}
func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting all courses")
	w.Header().Set("Content-Type", "application/json")
	courses = []Course{}
	json.NewEncoder(w).Encode("All courses have been deleted")
}

func errorHandler(err error) {
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

}
