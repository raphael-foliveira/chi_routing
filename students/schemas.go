package students

import "github.com/raphael-foliveira/hot_reload/programminglanguages"

type StudentDetailed struct {
	Id                   int                                        `json:"id"`
	FirstName            string                                     `json:"firstName"`
	LastName             string                                     `json:"lastName"`
	ProgrammingLanguages []programminglanguages.ProgrammingLanguage `json:"programmingLanguages"`
}
