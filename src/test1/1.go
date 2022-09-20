package test1

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type person struct {
	FirstName   string
	MiddleNames []string
	LastName    *string
}

func toCapital(v string) string {
	return cases.Title(language.Und, cases.NoLower).String(strings.ToLower(v))
}

func parser(fullName string) person {
	parsedName := fullName

	personName := person{
		FirstName:   "",
		MiddleNames: []string{},
	}

	lastNameIndex := strings.LastIndex(fullName, " ")
	if lastNameIndex != -1 {
		fullName = fullName[lastNameIndex+1:]
		personName.LastName = &fullName

		parsedName = toCapital(parsedName[:lastNameIndex])
	}

	personName.FirstName, parsedName, _ = strings.Cut(parsedName, " ")

	if parsedName != "" {
		personName.MiddleNames = strings.Split(parsedName, " ")
	}
	return personName
}
