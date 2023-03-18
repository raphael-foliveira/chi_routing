package students

import (
	"github.com/raphael-foliveira/hot_reload/database"
	"github.com/raphael-foliveira/hot_reload/programminglanguages"
)

func FindAll() []Student {
	var sts []Student
	database.Manager(&sts).QueryRows("SELECT id, firstName, lastName FROM student")
	return sts
}

func FindOne(id string) Student {
	var st Student
	database.Manager(&st).QueryRow("SELECT id, firstName, lastName FROM student WHERE id = $1", id)
	return st
}

func FindStudentProgrammingLanguages(id string) []programminglanguages.ProgrammingLanguage {
	pls := []programminglanguages.ProgrammingLanguage{}
	manager := database.Manager(&pls)
	manager.QueryRows("select pl.id, pl.name, pl.difficulty from programminglanguage pl join student_programminglanguage sp on pl.id = sp.student where sp.student = $1", id)
	return pls

}

func FindOneWithProgrammingLanguages(id string) map[string]any {
	student := FindOne(id)
	studentPls := FindStudentProgrammingLanguages(id)
	return map[string]any{
		"id":                   student.Id,
		"firstName":            student.FirstName,
		"lastName":             student.LastName,
		"programmingLanguages": studentPls,
	}
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
