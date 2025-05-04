package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Figma")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var posX int32 = 0
	var posY int32 = 0
	var width int32 = 0
	var height int32 = 0
	var col = rl.Black

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(posX, posY, width, height, col)
		rl.EndDrawing()
	}
}
