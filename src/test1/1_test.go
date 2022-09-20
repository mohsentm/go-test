/*
 * Case sensitivity:
 *  I coded with this assumption that First and Middle name should be in capital letter format ( as the way exists in the expected).
 *  except for the Family name that should be exactly the same as entered
 *
 * BUG: invalid Family case format
 * `WAHLGREN` case letter wasn't matching with expected family name (`Wahlgren`)
 *
 */
package test1

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

type testCase struct {
	raw      string
	expected person
}

func strPtr(v string) *string { return &v }

func TestNameParsing(t *testing.T) {
	testCases := []testCase{
		{raw: "Michael Daniel Jäger", expected: person{FirstName: "Michael", MiddleNames: []string{"Daniel"}, LastName: strPtr("Jäger")}},
		{raw: "LINUS HARALD christer Wahlgren", expected: person{FirstName: "Linus", MiddleNames: []string{"Harald", "Christer"}, LastName: strPtr("Wahlgren")}},
		{raw: "Pippilotta Viktualia Rullgardina Krusmynta Efraimsdotter LÅNGSTRUMP", expected: person{FirstName: "Pippilotta", MiddleNames: []string{"Viktualia", "Rullgardina", "Krusmynta", "Efraimsdotter"}, LastName: strPtr("LÅNGSTRUMP")}},
		{raw: "Kalle Anka", expected: person{FirstName: "Kalle", MiddleNames: []string{}, LastName: strPtr("Anka")}},
		{raw: "Ghandi", expected: person{FirstName: "Ghandi", MiddleNames: []string{}}},
	}
	for _, test := range testCases {
		t.Run(test.raw, func(t *testing.T) {
			actual := parser(test.raw)
			if !cmp.Equal(test.expected, actual) {
				t.Log(cmp.Diff(test.expected, actual))
				t.Fail()
			}
		})
	}
}
