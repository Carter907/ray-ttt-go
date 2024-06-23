package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
	squareSize   = 50
	squareMargin = 3
)

func main() {

	rl.InitWindow(screenWidth, screenHeight, "Tic Tac Toe")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	timeElasped := 0.0
	board := [3][3]int{
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}
	const boardWidth = 3*squareSize + 4*squareMargin
	const boardHeight = 3*squareSize + 4*squareMargin
	isX := false

	for !rl.WindowShouldClose() {

		timeElasped += float64(rl.GetFrameTime())
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Tic Tac Toe!!", screenWidth/2, 30, 30, rl.Black)

		if timeElasped >= 2.0 {
			rl.DrawText("Time to Play!", screenWidth/2, 60, 20, rl.Blue)
		}

		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {

				rl.DrawRectangleLines(
					int32((screenWidth-boardWidth)/2+i*(squareSize+squareMargin)),
					int32((screenHeight-boardHeight)/2+j*(squareSize+squareMargin)),
					squareSize,
					squareSize,
					rl.Black,
				)

			}
		}

		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				x := (screenWidth-boardWidth)/2 + i*(squareSize+squareMargin)
				y := (screenHeight-boardHeight)/2 + j*(squareSize+squareMargin)
				if rl.IsMouseButtonPressed(0) {
					cursorPos := rl.GetMousePosition()
					fmt.Print(cursorPos)

					if cursorPos.X >= float32(x) && cursorPos.X <= float32(x+squareSize) {
						if cursorPos.Y >= float32(y) && cursorPos.Y <= float32(y+squareSize) {
							isX = !isX
							if isX {
								board[i][j] = 1
							} else {
								board[i][j] = 0
							}
						}
					}

				}

				if board[i][j] == 1 {
					rl.DrawLine(int32(x)+squareSize/6, int32(y)+squareSize/6,
						int32(x)+(83*squareSize/100), int32(y)+(83*squareSize/100), rl.Red)
					rl.DrawLine(int32(x)+(83*squareSize/100), int32(y)+squareSize/6,
						int32(x)+squareSize/6, int32(y)+(83*squareSize/100), rl.Red)
				} else if board[i][j] == 0 {

					rl.DrawCircleLines(int32(x)+squareSize/2, int32(y)+squareSize/2, (squareSize/2)*.8, rl.Blue)
				}
			}

		}

		rl.EndDrawing()
	}
}
