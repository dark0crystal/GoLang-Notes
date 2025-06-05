package main



// spin up the server first : go run main.go 
// then send request here i used curl : 
// 1.curl http://localhost:8080/students
// output:[{"id":1,"name":"Ali","grade":3},{"id":2,"name":"Sara","grade":2}]
// 2.http://localhost:8080/students/1
// output:{"id":1,"name":"Ali","grade":3}
// 3.curl -X POST http://localhost:8080/students \
//   -H "Content-Type: application/json" \
//   -d '{"name": "Omar", "grade": 4}'
// output: {"id":3,"name":"Omar","grade":4}

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

var students = []Student{
	{ID: 1, Name: "Ali", Grade: 3},
	{ID: 2, Name: "Sara", Grade: 2},
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func getStudentByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/students/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, student := range students {
		if student.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(student)
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var newStudent Student
	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newStudent.ID = len(students) + 1
	students = append(students, newStudent)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newStudent)
}

func main() {
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getStudents(w, r)
		} else if r.Method == http.MethodPost {
			addStudent(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/students/", getStudentByID)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
