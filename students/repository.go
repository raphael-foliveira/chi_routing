package students

import (
	"sync"

	"github.com/raphael-foliveira/chi_routing/database"
	"github.com/raphael-foliveira/chi_routing/programminglanguages"
)

func FindAll() []Student {
	var sts []Student
	database.Manager(&sts).QueryRows("SELECT id, firstName, lastName FROM student")
	return sts
}

func FindOne(id string) (Student, error) {
	var st Student
	err := database.Manager(&st).QueryRow("SELECT id, firstName, lastName FROM student WHERE id = $1", id)
	if err != nil {
		return st, err
	}
	return st, nil
}

func FindStudentProgrammingLanguages(id string) ([]programminglanguages.ProgrammingLanguage, error) {
	pls := []programminglanguages.ProgrammingLanguage{}
	manager := database.Manager(&pls)
	err := manager.QueryRows("select pl.id, pl.name, pl.difficulty from programminglanguage pl join student_programminglanguage sp on pl.id = sp.student where sp.student = $1", id)
	if err != nil {
		return nil, err
	}
	return pls, nil
}

func FindOneWithProgrammingLanguages(id string) (map[string]any, error) {
	sChannel := make(chan Student)
	plsChannel := make(chan []programminglanguages.ProgrammingLanguage)
	errChannel := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		pls, err := FindStudentProgrammingLanguages(id)
		if err != nil {
			errChannel <- err
			plsChannel <- nil
			return
		}
		plsChannel <- pls
	}()
	go func() {
		defer wg.Done()
		s, err := FindOne(id)
		if err != nil {
			errChannel <- err
			sChannel <- Student{}
			return
		}
		sChannel <- s
	}()
	student := <-sChannel
	studentPls := <-plsChannel
	if len(errChannel) > 0 {
		return nil, <-errChannel
	}
	wg.Wait()
	return map[string]any{
		"id":                   student.Id,
		"firstName":            student.FirstName,
		"lastName":             student.LastName,
		"programmingLanguages": studentPls,
	}, nil
}

func Create(student Student) Student {
	var newStudent Student
	database.Manager(&newStudent).QueryRow("INSERT INTO student (firstName, lastName) VALUES ($1, $2) RETURNING id, firstName, lastName", student.FirstName, student.LastName)
	return newStudent

}

func Delete(id string) Student {
	var deletedStudent Student
	database.Manager(&deletedStudent).QueryRow("DELETE FROM student WHERE id = $1 RETURNING id, firstName, lastName", id)
	return deletedStudent
}
