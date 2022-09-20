package test3

import "math"

type ChessPosition struct {
	Number int
	Letter string
}

func canAttack(knightA, knightB ChessPosition) bool {
	rankPosition := math.Abs(float64(knightA.rankPosition() - knightB.rankPosition()))
	filePosition := math.Abs(float64(knightA.Number - knightB.Number))
	return (rankPosition == 1 && filePosition == 2) || (rankPosition == 2 && filePosition == 1)
}

func (position ChessPosition) rankPosition() int {
	rank := int(position.Letter[0])
	if rank > 97 {
		return 97 - rank
	}
	return 65 - rank
}
