package board_test

import (
	boardUtils "github.com/Carter907/ray-ttt-go/board"
	"testing"
)

func TestBoardContents(t *testing.T) {

	b := boardUtils.NewBoard()

	for i := 0; i < len(b.Mat); i++ {
		for j := 0; j < len(b.Mat[i]); j++ {

			if b.Mat[i][j] != boardUtils.NO_TEAM {
				t.Fatal("initial position of board is not correct")
			}
		}
	}

}
func TestBoardDraw(t *testing.T) {
	b := boardUtils.NewBoard()

	if boardUtils.IsDraw(&b) {
		t.Fatalf("board can't draw it was just initialized.")
	}

	b.Mat = [3][3]int{
		{0, 1, 0},
		{0, 1, 0},
		{1, 0, 1},
	}

	if !boardUtils.IsDraw(&b) {
		t.Errorf("board should draw but did not\n")
		t.Logf("board state: %v", b.Mat)
	}

}

func TestBoardSlices(t *testing.T) {

	b := boardUtils.NewBoard()
	b.Mat[0][0] = boardUtils.TEAM_X
	b.Mat[1][0] = boardUtils.TEAM_X
	b.Mat[2][0] = boardUtils.TEAM_X

	t.Log(b.Mat)

	t.Log(b.Mat[0:])

}

func TestBoardScores(t *testing.T) {

	b := boardUtils.NewBoard()

	if b.OScore != 0 || b.XScore != 0 {
		t.Fatal("scores of board are not initialized to 0")
	}
}

func TestCheckWinner(t *testing.T) {
	b := boardUtils.NewBoard()

	t.Log(b.Mat)

	winner, _ := boardUtils.CheckWinner(&b)

	if winner != boardUtils.NO_TEAM {
		t.Fatalf("should've been no team but was %v", winner)
	}
}

func TestTypes(t *testing.T) {
	b := 3.2

	t.Logf("%T\n", b)
}
