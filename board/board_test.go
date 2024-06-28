package board

import (
	"reflect"
	"testing"
)

func TestBoardContents(t *testing.T) {

	b := NewBoard()

	for i := 0; i < len(b.Mat); i++ {
		for j := 0; j < len(b.Mat[i]); j++ {

			if b.Mat[i][j] != NO_TEAM {
				t.Fatal("initial position of board is not correct")
			}
		}
	}

}
func TestBoardSlices(t *testing.T) {

	b := NewBoard()
	b.Mat[0][0] = TEAM_X
	b.Mat[1][0] = TEAM_X
	b.Mat[2][0] = TEAM_X

	t.Log(b.Mat)

	t.Log(b.Mat[0:])

}

func TestBoardScores(t *testing.T) {

	b := NewBoard()

	if b.OScore != 0 || b.XScore != 0 {
		t.Fatal("scores of board are not initialized to 0")
	}
}

func TestCheckWinner(t *testing.T) {
	b := NewBoard()

	t.Log(b.Mat)

	winner, _ := CheckWinner(&b)

	if winner != NO_TEAM {
		t.Fatalf("should've been no team but was %v", winner)
	}
}

func TestTypes(t *testing.T) {
	b := 3.2

	t.Log(reflect.TypeOf(b))
}
