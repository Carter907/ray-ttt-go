package main

import (
	"fmt"
	boardUtils "github.com/Carter907/ray-ttt-go/board"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
	squareSize   = 70
	squareMargin = 3
	boardWidth   = 3*squareSize + 4*squareMargin
	boardHeight  = 3*squareSize + 4*squareMargin
)

func main() {

	rl.InitWindow(screenWidth, screenHeight, "Tic Tac Toe")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	timeElasped := 0.0
	board := boardUtils.New()
	isX := false
	teamWon := boardUtils.NO_TEAM

	for !rl.WindowShouldClose() {

		timeElasped += float64(rl.GetFrameTime())
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Tic Tac Toe!!", screenWidth/2, 30, 30, rl.Black)
		rl.DrawText(fmt.Sprintf("O Score: %d", board.OScore), 30, 30, 20, rl.Blue)
		rl.DrawText(fmt.Sprintf("X Score: %d", board.XScore), 30, 60, 20, rl.Red)

		if timeElasped >= 2.0 {
			rl.DrawText("Time to Play!", screenWidth/2, 60, 20, rl.Blue)
		}

		for i := 0; i < len(board.Mat); i++ {
			for j := 0; j < len(board.Mat[i]); j++ {
				x := int32((screenWidth-boardWidth)/2 + i*(squareSize+squareMargin))
				y := int32((screenHeight-boardHeight)/2 + j*(squareSize+squareMargin))

				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					cursorPos := rl.GetMousePosition()
					fmt.Print(cursorPos)

					if cursorPos.X >= float32(x) && cursorPos.X <= float32(x+squareSize) {
						if cursorPos.Y >= float32(y) && cursorPos.Y <= float32(y+squareSize) {
							isX = !isX
							if board.Mat[i][j] == boardUtils.NO_TEAM {
								if isX {
									board.Mat[i][j] = boardUtils.TEAM_X
								} else {
									board.Mat[i][j] = boardUtils.TEAM_O
								}
							}
						}
					}

				}
				rl.DrawRectangleLines(
					int32((screenWidth-boardWidth)/2+i*(squareSize+squareMargin)),
					int32((screenHeight-boardHeight)/2+j*(squareSize+squareMargin)),
					squareSize,
					squareSize,
					rl.Black,
				)
				if board.Mat[i][j] == 1 {
					rl.DrawLine(int32(x)+squareSize/6, int32(y)+squareSize/6,
						int32(x)+(83*squareSize/100), int32(y)+(83*squareSize/100), rl.Red)
					rl.DrawLine(int32(x)+(83*squareSize/100), int32(y)+squareSize/6,
						int32(x)+squareSize/6, int32(y)+(83*squareSize/100), rl.Red)
				} else if board.Mat[i][j] == 0 {

					rl.DrawCircleLines(int32(x)+squareSize/2, int32(y)+squareSize/2, (squareSize/2)*.8, rl.Blue)
				}

			}

		}

		team, _ := boardUtils.CheckWinner(board)

		teamWon = team

		if teamWon != boardUtils.NO_TEAM || boardUtils.IsDraw(board) {

			for i := 0; i < len(board.Mat); i++ {
				for j := 0; j < len(board.Mat[i]); j++ {
					board.Mat[i][j] = boardUtils.NO_TEAM
				}
			}
			switch teamWon {
			case boardUtils.TEAM_O:
				board.OScore += 1
			case boardUtils.TEAM_X:
				board.XScore += 1

			}

			teamWon = boardUtils.NO_TEAM
		}

		rl.EndDrawing()
	}
}
