package root

import (
	"net/http"

	"github.com/raphael-foliveira/chi_routing/utils"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	utils.Json(w, map[string]string{
		"message": "Hello World"})
}
