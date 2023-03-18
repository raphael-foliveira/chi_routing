package programminglanguages

import (
	"github.com/raphael-foliveira/chi_routing/database"
)

func FindAll() []ProgrammingLanguage {
	var pls []ProgrammingLanguage
	database.Manager(&pls).QueryRows("SELECT id, name, difficulty FROM programminglanguage")
	return pls
}

func FindOne(id string) ProgrammingLanguage {
	var pl ProgrammingLanguage
	database.Manager(&pl).QueryRow("SELECT id, name, difficulty FROM programminglanguage WHERE id = $1", id)
	return pl
}

func Create(pl ProgrammingLanguage) ProgrammingLanguage {
	var newPl ProgrammingLanguage
	database.Manager(&newPl).QueryRow("INSERT INTO programminglanguage (name, difficulty) VALUES ($1, $2) RETURNING id, name, difficulty", pl.Name, pl.Difficulty)
	return newPl
}

func Delete(id string) ProgrammingLanguage {
	var deletedPl ProgrammingLanguage
	database.Manager(&deletedPl).QueryRow("DELETE FROM programminglanguage WHERE id = $1 RETURNING id, name, difficulty", id)
	return deletedPl
}
