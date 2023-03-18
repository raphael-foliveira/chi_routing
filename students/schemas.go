package students

import "github.com/raphael-foliveira/chi_routing/programminglanguages"

type StudentDetailed struct {
	Id                   int                                        `json:"id"`
	FirstName            string                                     `json:"firstName"`
	LastName             string                                     `json:"lastName"`
	ProgrammingLanguages []programminglanguages.ProgrammingLanguage `json:"programmingLanguages"`
}
