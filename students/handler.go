package students

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/raphael-foliveira/chi_routing/utils"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	students := FindAll()

	utils.Json(w, students)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	student, err := getStudentFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	newStudent := Create(student)
	json.NewEncoder(w).Encode(newStudent)
}

func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	st := FindOneWithProgrammingLanguages(id)

	json.NewEncoder(w).Encode(st)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	deletedStudent := Delete(id)

	json.NewEncoder(w).Encode(map[string]any{
		"deleted": deletedStudent,
	})
}
