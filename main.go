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

var selectedValue = "selection"

var drawHighlight func() = nil
var camera = rl.Camera2D{}

func main() {
	rl.InitWindow(1200, 720, "Figma")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	camera.Zoom = 1

	for !rl.WindowShouldClose() {
		drawHighlight = nil

		if selectedValue == "selection" && rl.IsKeyDown(rl.KeySpace) && rl.IsMouseButtonDown(rl.MouseLeftButton) {
			delta := rl.GetMouseDelta()
			delta = rl.Vector2Scale(delta, -1.0/camera.Zoom)
			camera.Target = rl.Vector2Add(camera.Target, delta)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		layout := lib.NewConstrainedLayout(lib.ContrainedLayout{
			Direction: lib.DIRECTION_ROW,
			Contrains: rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())),
			ChildrenSize: []lib.ChildSize{
				{
					SizeType: lib.SIZE_WEIGHT,
					Value:    1,
				},
				{
					SizeType: lib.SIZE_ABSOLUTE,
					Value:    240,
				},
			},
		})

		rl.BeginMode2D(camera)
		targetRect := layout.Render(nil)
		intT := targetRect.ToInt32()
		rl.BeginScissorMode(intT.X, intT.Y, intT.Width, intT.Height)
		Canvas(targetRect)

		switch selectedValue {
		case "selection":
			if drawHighlight == nil {
				Selection()
			}
			if drawHighlight != nil {
				drawHighlight()
			}
		case "rectangle":
			RectangleTool()
		}

		rl.EndScissorMode()
		rl.EndMode2D()
		layout.Render(PanelStyles)

		BottomNavigation()

		rl.EndDrawing()
	}
}

func Canvas(rect rl.Rectangle) {
	for _, l := range layers {
		isClicked := l.DrawComponent(&ui, rl.GetScreenToWorld2D(rl.GetMousePosition(), camera))
		if isClicked {
			selectedLayer = l
		}

		if l.State() == components.STATE_HOT || l.State() == components.STATE_ACTIVE {
			drawHighlight = l.DrawHighlight
		}
	}
}

func PanelStyles(rect rl.Rectangle) {
	if selectedLayer == nil {
		return
	}

	selectedLayer.DrawControls(&ui, rect)
}

var items = [2]string{"selection", "rectangle"}

func BottomNavigation() {
	padding := lib.Padding{}
	padding.All(8)
	layout := lib.NewLayout(lib.Layout{
		Direction: lib.DIRECTION_ROW,
		Gap:       8,
		Padding:   padding,
	}, rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())))

	for _, item := range items {
		layout.Render(Item(item))
	}
}

func Item(item string) lib.Component {
	return func(avaliablePosition lib.Position, next lib.Next) {
		var size float32 = 32
		rect := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Contrains.Height-size, size, size)
		interactable := components.NewInteractable(item, &ui)
		if interactable.Event(rl.GetMousePosition(), rect) {
			selectedValue = item

			startPos = nil
			currentPos = nil
		}

		var color = rl.Fade(rl.Blue, 0.32)
		if selectedValue == item {
			color = rl.Blue
		}

		rl.DrawRectanglePro(rect, rl.Vector2{}, 0, color)

		next(rect)
	}
}
