package student

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB *sqlx.DB
}

type createStudentRequest struct {
	Name string `json:"name"`
}

type updateStudentRequest struct {
	Name string `json:"name"`
}

// HandleStudents handles:
// GET  /students
// POST /students
func (h *Handler) HandleStudents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getStudents(w, r)

	case http.MethodPost:
		h.createStudent(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleStudent handles:
// GET    /students/{id}
// PUT    /students/{id}
// DELETE /students/{id}
func (h *Handler) HandleStudent(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/students/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid student id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getStudent(w, r, id)

	case http.MethodPut:
		h.updateStudent(w, r, id)

	case http.MethodDelete:
		h.deleteStudent(w, r, id)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// POST /students
func (h *Handler) createStudent(w http.ResponseWriter, r *http.Request) {
	var request createStudentRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(request.Name) == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	student, err := CreateStudent(h.DB, request.Name)
	if err != nil {
		http.Error(w, "failed to create student", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, student)
}

// GET /students
func (h *Handler) getStudents(w http.ResponseWriter, r *http.Request) {
	students, err := GetStudents(h.DB)
	if err != nil {
		http.Error(w, "failed to get students", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, students)
}

// GET /students/{id}
func (h *Handler) getStudent(w http.ResponseWriter, r *http.Request, id int) {
	student, err := GetStudent(h.DB, id)
	if err != nil {
		http.Error(w, "student not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, student)
}

// PUT /students/{id}
func (h *Handler) updateStudent(w http.ResponseWriter, r *http.Request, id int) {
	var request updateStudentRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(request.Name) == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	student, err := UpdateStudent(h.DB, id, request.Name)
	if err != nil {
		http.Error(w, "student not found", http.StatusNotFound)
		return
	}

	fmt.Println("Student updated successfully")

	writeJSON(w, http.StatusOK, student)
}

// DELETE /students/{id}
func (h *Handler) deleteStudent(w http.ResponseWriter, r *http.Request, id int) {
	err := DeleteStudent(h.DB, id)
	if err != nil {
		http.Error(w, "failed to delete student", http.StatusInternalServerError)
		return
	}

	fmt.Println("Student deleted successfully")

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
