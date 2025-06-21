package main

import (
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/layer"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var layers = []layer.Layer{}
var selectedLayer layer.Layer
var ui = lib.UIStruct{}
var drawHighlight func() = nil
var camera = rl.Camera2D{}
var c = components.NewComponents(&ui)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagBorderlessWindowedMode)

	rl.InitWindow(1600, 900, "Figma")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	camera.Zoom = 1

	IconTexture = LoadSVGAsTexture("D:\\Peronal\\figma\\assets\\icons\\type.svg", 16, 16)

	for !rl.WindowShouldClose() {
		rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Black)

		if rl.IsKeyPressed(rl.KeySpace) {
			ui.IsPlaying = !ui.IsPlaying
		}

		padding := lib.Padding{}
		padding.All(PANEL_GAP)
		bodyLayout := lib.NewConstrainedLayout(lib.ContrainedLayout{
			Direction: lib.DIRECTION_COLUMN,
			Gap:       PANEL_GAP,
			Padding:   padding,
			Contrains: rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())),
			ChildrenSize: []lib.ChildSize{
				{
					SizeType: lib.SIZE_WEIGHT,
					Value:    1,
				},
				{
					SizeType: lib.SIZE_ABSOLUTE,
					Value:    BOTTOM_PANEL_HEIGHT,
				},
			},
		})
		bodyLayout.Add(UpperPart)
		bodyLayout.Add(Timeline)
		bodyLayout.Draw()

		rl.EndDrawing()
	}
}
