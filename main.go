package main

import (
	"fmt"
	"strconv"

	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var ui = lib.UIStruct{}

func main() {
	rl.InitWindow(800, 450, "Figma")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var posX int32 = 0
	var posY int32 = 1
	var width int32 = 200
	var height int32 = 100
	var col = rl.Black

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(posX, posY, width, height, col)

		var newValue = components.Input(components.InputProps{
			X:           240,
			Y:           0,
			Id:          "pox-y",
			Width:       200,
			Placeholder: "Posicao y do retangulo.",
			Value:       fmt.Sprint(posY),
			MousePoint:  rl.GetMousePosition(),
			Ui:          &ui,
		})
		var newIntValue, _ = strconv.Atoi(newValue)
		posY = int32(newIntValue)
		posY -= int32(rl.GetMouseWheelMove() * 4)
		rl.EndDrawing()
	}
}
