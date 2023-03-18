package programminglanguages

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	pls := FindAll()
	json.NewEncoder(w).Encode(pls)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	pl := FindOne(id)
	json.NewEncoder(w).Encode(pl)
}

func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	pl := FindOne(id)
	json.NewEncoder(w).Encode(pl)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	pl := FindAll()
	json.NewEncoder(w).Encode(pl)
}
