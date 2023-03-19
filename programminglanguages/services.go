package programminglanguages

import (
	"encoding/json"
	"errors"
	"net/http"
)

func validatePlData(pl ProgrammingLanguage) error {
	if pl.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

func getPlFromBody(r *http.Request) (ProgrammingLanguage, error) {
	var pl ProgrammingLanguage
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&pl)

	err := validatePlData(pl)
	return pl, err
}
