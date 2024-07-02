package board_test

import (
	board "github.com/Carter907/ray-ttt-go/board"
	"testing"
)

// Testing that the contents of the board are correctly updating
func TestBoardContents(t *testing.T) {

	b := board.New()

	for i := 0; i < len(b.Mat); i++ {
		for j := 0; j < len(b.Mat[i]); j++ {

			if b.Mat[i][j] != board.NO_TEAM {
				t.Fatalf("incorrect board contents: %v", b.Mat[i][j])
			}
		}
	}

}

// Testing that the board IsDraw function correctly identifies drawed positions
func TestBoardDraw(t *testing.T) {
	b := board.New()

	if board.IsDraw(b) {
		t.Fatalf("board can't draw it was just initialized.")
	}

	b.Mat = [3][3]int{
		{0, 1, 0},
		{0, 1, 0},
		{1, 0, 1},
	}

	if !board.IsDraw(b) {
		t.Errorf("board should draw but did not\nboard state: %v", b.Mat)
	}
	b.Mat = [3][3]int{
		{0, 0, 1},
		{1, 0, 0},
		{1, 0, 1},
	}
	if !board.IsDraw(b) {
		t.Errorf("board should draw but did not\nboard state: %v", b.Mat)
	}

}

// Test that the board scores are initialized properly
func TestBoardScores(t *testing.T) {

	b := board.New()

	if b.OScore != 0 || b.XScore != 0 {
		t.Fatal("scores of board are not initialized to 0")
	}
}

// Test CheckWinner function that should correctly identifiy a winning position and return
// the squares of the board associated.
func TestCheckWinner(t *testing.T) {
	b := board.New()

	winner, squares := board.CheckWinner(b)

	if winner != board.NO_TEAM {

		t.Fatalf("should've been team %v, but was %v", board.TEAM_O, winner)
	}
	if team, _ := board.CheckSquaresMatch(squares); team != board.NO_TEAM {

		t.Fatalf("should've been team %v, but was %v", board.TEAM_O, winner)
	}

	b.Mat = [3][3]int{

		{0, 0, 0},
		{1, -3, -3},
		{-3, 1, -3},
	}

	winner, squares = board.CheckWinner(b)

	if winner != board.TEAM_O {

		t.Fatalf("should've been team %v, but was %v", board.TEAM_O, winner)
	}
	if team, _ := board.CheckSquaresMatch(squares); team != board.TEAM_O {

		t.Fatalf("should've been team %v, but was %v", board.TEAM_O, winner)
	}
}
