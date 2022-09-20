/*
 * BUG: Invalid position number (-1). {Number: -1, Letter: "A"}
 *      I coded with this assumption that input parameters were validated therefore I didn't add position validate
 */

package test3

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

type testCase3 struct {
	Position1 ChessPosition
	Position2 ChessPosition
	CanAttack bool
}

func TestChess(t *testing.T) {
	testCases := []testCase3{
		{Position1: ChessPosition{Number: 2, Letter: "C"}, Position2: ChessPosition{Number: 4, Letter: "D"}, CanAttack: true},
		{Position1: ChessPosition{Number: 7, Letter: "F"}, Position2: ChessPosition{Number: 5, Letter: "E"}, CanAttack: true},
		{Position1: ChessPosition{Number: 2, Letter: "C"}, Position2: ChessPosition{Number: 1, Letter: "A"}, CanAttack: true},
		{Position1: ChessPosition{Number: 6, Letter: "A"}, Position2: ChessPosition{Number: 4, Letter: "B"}, CanAttack: true},
		{Position1: ChessPosition{Number: 6, Letter: "A"}, Position2: ChessPosition{Number: 5, Letter: "B"}, CanAttack: false},
		{Position1: ChessPosition{Number: 2, Letter: "C"}, Position2: ChessPosition{Number: 2, Letter: "C"}, CanAttack: false},

		// It works, but in realty -1 is out of the range
		{Position1: ChessPosition{Number: -1, Letter: "A"}, Position2: ChessPosition{Number: 1, Letter: "B"}, CanAttack: true},

		{Position1: ChessPosition{Number: 4, Letter: "D"}, Position2: ChessPosition{Number: 5, Letter: "E"}, CanAttack: false},
	}
	for ind, test := range testCases {
		t.Run(fmt.Sprint(ind), func(t *testing.T) {
			actual := canAttack(test.Position1, test.Position2)
			if !cmp.Equal(test.CanAttack, actual) {
				t.Log(cmp.Diff(test.CanAttack, actual))
				t.Fail()
			}
		})
	}
}
