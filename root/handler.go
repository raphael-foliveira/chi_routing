package root

import (
	"net/http"

	"github.com/raphael-foliveira/hot_reload/utils"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	utils.Json(w, map[string]string{
		"message": "Hello World"})
}
