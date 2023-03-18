package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func Json(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetId(r *http.Request) (int, error) {
	stringId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(stringId, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
