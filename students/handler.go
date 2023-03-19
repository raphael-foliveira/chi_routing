package students

import (
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
	utils.Json(w, newStudent)
}

func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	st := FindOneWithProgrammingLanguages(id)

	utils.Json(w, st)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	deletedStudent := Delete(id)

	utils.Json(w, map[string]any{
		"deleted": deletedStudent,
	})
}
