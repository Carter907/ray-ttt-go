package board

import (
	"testing"
)

func TestBoardContents(t *testing.T) {

	b := New()

	for i := 0; i < len(b.Mat); i++ {
		for j := 0; j < len(b.Mat[i]); j++ {

			if b.Mat[i][j] != NO_TEAM {
				t.Fatal("initial position of board is not correct")
			}
		}
	}

}

func TestBoardScores(t *testing.T) {
	b := New()

	if b.OScore != 0 || b.XScore != 0 {
		t.Fatal("scores of board are not initialized to 0")
	}
}

func TestCheckWinner(t *testing.T) {
	b := New()

	b.Mat[0][0] = TEAM_X
	b.Mat[0][1] = TEAM_X
	b.Mat[0][2] = TEAM_X

	t.Log(b.Mat)

	winner := CheckWinner(&b)

	if winner != TEAM_X {
		t.Fatalf("should've been team X but was %v", winner)
	}
}
