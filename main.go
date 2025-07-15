package main

import (
	"strconv"

	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/layer"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var layers = []layer.Layer{}
var selectedLayer layer.Layer
var ui = lib.NewUi()
var drawHighlight func(lib.UIStruct, components.Components) = nil
var camera = rl.Camera2D{}
var c = components.NewComponents(&ui)

func AppendLayer(newLayer layer.Layer) {
	layers = append(layers, newLayer)
	startPosRec = nil
	currentPosRec = nil
	selectedLayer = newLayer
	ui.SelectedTool = lib.ToolSelection
}

func NewLayerId() string {
	id := 0
	if len(layers) > 0 {
		newId, _ := strconv.Atoi(layers[len(layers)-1].GetElement().Id)
		id = newId
	}
	return strconv.Itoa(id + 1)
}

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
		// for _, layer := range layers {
		// 	fmt.Println(layer)
		// }

		ui.ResetTabOrder()
		c.FrameReset()

		rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Black)

		if rl.IsKeyPressed(rl.KeySpace) {
			ui.TogglePlay()
		}

		if ui.IsPlaying {
			if visibleFrames[1] == ui.SelectedFrame {
				ui.SelectedFrame = visibleFrames[0]
			} else {
				ui.SelectedFrame++
			}
		}

		lib.NewLayout().
			Position(rl.NewVector2(0, 0)).
			Padding(lib.NewPadding().All(PANEL_GAP)).
			Gap(PANEL_GAP).
			Column().
			Width(float32(rl.GetScreenWidth())).
			Height(float32(rl.GetScreenHeight()),
				lib.ChildSize{SizeType: lib.SIZE_WEIGHT, Value: 1},
				lib.ChildSize{SizeType: lib.SIZE_ABSOLUTE, Value: BOTTOM_PANEL_HEIGHT}).
			Add(UpperPart()).
			Add(Timeline()).
			Draw()

		rl.EndDrawing()

		// TODO
		// IMPROVE THIS
		if rl.IsKeyPressed(rl.KeyTab) {
			nextTabIndex := -1
			for index, inputId := range ui.TabOrder {
				if inputId == ui.FocusedId {
					if index+1 > len(ui.TabOrder)-1 {
						nextTabIndex = 0
					} else {
						nextTabIndex = index + 1
					}
				}
			}

			if nextTabIndex != -1 {
				ui.FocusedId = ui.TabOrder[nextTabIndex]
				ui.SetCursors(0)
			}
		}
	}
}
