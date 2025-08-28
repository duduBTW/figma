package main

import (
	"fmt"

	"github.com/dudubtw/figma/app"
	ds "github.com/dudubtw/figma/design-system"
	"github.com/dudubtw/figma/home"
	newWorkplace "github.com/dudubtw/figma/new-workplace"
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

	app.Apk.CreateWorkplace = app.NewCreateWorkplace()
	app.Apk.TypographyMap = app.InitTypography()
	app.HomeLoad()

	for !rl.WindowShouldClose() {
		app.Apk.Frame()
		DrawBackground()

		switch app.Apk.SelectedPage {
		case app.PAGE_WORKPLACE:
			WorkplacePage()
		case app.PAGE_HOME:
			home.Page()
		case app.PAGE_NEW_WORKPLACE:
			newWorkplace.Page()
		}

		rl.EndDrawing()
	}
}

func DrawBackground() {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), ds.T2_COLOR_BODY)
}

func WorkplacePage() {
	if app.Apk.Workplace.Id == "" {
		fmt.Println("Workplace page opened with no Id.")
		panic(1)
	}

	if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyS) {
		app.Apk.Save()
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		app.Apk.Workplace.TogglePlay()
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
}

// TODO
// ui.ResetTabOrder()
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
