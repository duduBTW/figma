package main

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var camera = rl.Camera2D{}

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagBorderlessWindowedMode)

	rl.InitWindow(1280, 800, "Figma")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	camera.Zoom = 1

	// IconTexture = LoadSVGAsTexture("D:\\Peronal\\figma\\assets\\icons\\type.svg", 16, 16)

	for !rl.WindowShouldClose() {
		// ui.ResetTabOrder()
		app.Apk.Frame()

		rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Black)
		if rl.IsKeyPressed(rl.KeySpace) {
			app.Apk.TogglePlay()
		}

		app.NewLayout().
			Position(rl.NewVector2(0, 0)).
			Padding(app.NewPadding().All(PANEL_GAP)).
			Gap(PANEL_GAP).
			Column().
			Width(float32(rl.GetScreenWidth())).
			Height(float32(rl.GetScreenHeight()),
				app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1},
				app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: BOTTOM_PANEL_HEIGHT}).
			Add(UpperPart()).
			Add(Timeline()).
			Draw()

		rl.EndDrawing()

		// TODO
		// IMPROVE THIS
		// if rl.IsKeyPressed(rl.KeyTab) {
		// 	nextTabIndex := -1
		// 	for index, inputId := range ui.TabOrder {
		// 		if inputId == ui.FocusedId {
		// 			if index+1 > len(ui.TabOrder)-1 {
		// 				nextTabIndex = 0
		// 			} else {
		// 				nextTabIndex = index + 1
		// 			}
		// 		}
		// 	}

		// 	if nextTabIndex != -1 {
		// 		ui.FocusedId = ui.TabOrder[nextTabIndex]
		// 		ui.SetCursors(0)
		// 	}
		// }
	}
}
