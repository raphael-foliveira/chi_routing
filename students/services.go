package students

import (
	"encoding/json"
	"errors"
	"net/http"
)

func validateStudentData(student Student) error {
	if student.FirstName == "" {
		return errors.New("firstName is required")
	}
	if student.LastName == "" {
		return errors.New("lastName is required")
	}
	return nil
}

func getStudentFromBody(r *http.Request) (Student, error) {
	var student Student
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		return Student{}, err
	}

	err = validateStudentData(student)
	if err != nil {
		return Student{}, err
	}

	return student, err
}
