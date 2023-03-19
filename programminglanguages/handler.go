package programminglanguages

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/raphael-foliveira/chi_routing/utils"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	pls := FindAll()
	utils.Json(w, pls)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	pl, err := getPlFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	newPl := Create(pl)
	utils.Json(w, newPl)
}

func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	pl := FindOne(chi.URLParam(r, "id"))
	utils.Json(w, pl)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	deletedPl := Delete(chi.URLParam(r, "id"))
	utils.Json(w, map[string]any{
		"deleted": deletedPl,
	})
}
