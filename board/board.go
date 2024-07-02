package board

const (
	TEAM_X  = 1
	TEAM_O  = 0
	NO_TEAM = -3
)

type Board struct {
	Mat    [3][3]int
	OScore int
	XScore int
}

func New() *Board {

	return &Board{
		Mat: [3][3]int{
			{NO_TEAM, NO_TEAM, NO_TEAM},
			{NO_TEAM, NO_TEAM, NO_TEAM},
			{NO_TEAM, NO_TEAM, NO_TEAM},
		},
		OScore: 0,
		XScore: 0,
	}
}

func CheckWinner(board *Board) (team int, arr [3]int) {
	team = NO_TEAM
	arr = [3]int{-3, -3, -3}

	// check rows

	for i := 0; i < len(board.Mat); i++ {
		if team, arr = CheckSquaresMatch(board.Mat[i]); team != NO_TEAM {
			return
		}
	}

	// check cols

	for i := 0; i < len(board.Mat); i++ {
		if team, arr = CheckSquaresMatch(getCol(board.Mat, i)); team != NO_TEAM {

			return
		}
	}

	// check diagonals

	if team, arr = CheckSquaresMatch(getDiagonal(board.Mat, false)); team != NO_TEAM {

		return
	}
	if team, arr = CheckSquaresMatch(getDiagonal(board.Mat, true)); team != NO_TEAM {

		return
	}

	return
}

func IsDraw(board *Board) bool {

	return false
}

func CheckSquaresMatch(arr [3]int) (team int, _arr [3]int) {
	team = NO_TEAM
	sum := 0

	for _, tm := range arr {

		sum += tm

	}
	if sum == 3 {
		team = TEAM_X
	} else if sum == 0 {
		team = TEAM_O
	}

	_arr = arr

	return
}

func getCol(mat [3][3]int, index int) (col [3]int) {

	for i := 0; i < len(mat); i++ {
		col[i] = mat[i][index]
	}

	return
}

func getDiagonal(mat [3][3]int, rotate bool) (diag [3]int) {

	for i := 0; i < len(mat); i++ {

		if rotate {
			diag[i] = mat[i][2-i]
		} else {
			diag[i] = mat[i][i]
		}
	}

	return
}
