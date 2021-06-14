// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book - Our struct for all books
type Student struct {
	StudentId string `json:"StudentId"`
	Branch    string `json:"Branch"`
	College   string `json:"College"`
}

var Students []Student

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Students)
}

func getStudentByStudentId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["studentid "]

	for _, Student := range Students {
		if Student.StudentId == key {
			json.NewEncoder(w).Encode(Student)
		}
	}
}

func createNewStudent(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Student Student
	json.Unmarshal(reqBody, &Student)
	Students = append(Students, Student)
	json.NewEncoder(w).Encode(Student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentid := vars["studentid"]

	for index, Student := range Students {
		if Student.StudentId == studentid {
			Students = append(Students[:index], Students[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/Students", getStudents)
	myRouter.HandleFunc("/Student", createNewStudent).Methods("POST")
	myRouter.HandleFunc("/Student/{studentid }", deleteStudent).Methods("DELETE")
	myRouter.HandleFunc("/Student/{studentid }", getStudentByStudentId)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	Students = []Student{
		Student{StudentId: "1", Branch: "ECE", College: "SITAMS"},
		Student{StudentId: "2", Branch: "EEE", College: "SVCET"},
		Student{StudentId: "3", Branch: "CSE", College: "VEMU"},
	}
	handleRequests()
}
