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

func New() Board {

	return Board{
		Mat: [3][3]int{
			{NO_TEAM, NO_TEAM, NO_TEAM},
			{NO_TEAM, NO_TEAM, NO_TEAM},
			{NO_TEAM, NO_TEAM, NO_TEAM},
		},
		OScore: 0,
		XScore: 0,
	}
}
func CheckWinner(board *Board) (team int) {
	team = NO_TEAM

	vsum, hsum := 0, 0
	for i := 0; i < len(board.Mat); i++ {
		vsum, hsum = 0, 0
		for j := 0; j < len(board.Mat[i]); j++ {

			vsum += board.Mat[i][j]
			hsum += board.Mat[j][i]
		}
		if vsum == 3 || hsum == 3 {
			team = TEAM_X
			break
		}
		if vsum == 0 || hsum == 0 {
			team = TEAM_O
			break
		}

	}

	return
}
