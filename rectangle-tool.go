package main

import (
	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/layer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var startPosRec *rl.Vector2
var currentPosRec *rl.Vector2

func CalculateRectangle() rl.Rectangle {
	width := currentPosRec.X - startPosRec.X
	height := currentPosRec.Y - startPosRec.Y

	// Handle dragging in any direction
	rect := rl.NewRectangle(startPosRec.X, startPosRec.Y, width, height)
	if width < 0 {
		rect.X += width
		rect.Width = -width
	}
	if height < 0 {
		rect.Y += height
		rect.Height = -height
	}

	return rect
}

func RectangleTool(container rl.Rectangle) {
	RectangleSelectionActionHandler(container)

	if startPosRec == nil || currentPosRec == nil {
		return
	}

	rect := CalculateRectangle()

	// Draw the selection rectangle (semi-transparent fill + border)
	rl.DrawRectangleRec(rect, rl.Fade(rl.White, 0.22))
	rl.DrawRectangleLinesEx(rect, 1, rl.Blue)
}

func RectangleSelectionActionHandler(container rl.Rectangle) {
	mousePos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
	if startPosRec != nil && rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		rect := CalculateRectangle()
		index := 0
		for _, l := range app.Apk.Workplace.Layers {
			_, isRec := l.(*layer.Rectangle)
			if isRec {
				index++
			}
		}

		newLayer := layer.NewRectangle(app.Apk.Workplace.NewLayerId(), rect, index)
		app.Apk.Workplace.AppendLayer(&newLayer)
		startPosRec = nil
		currentPosRec = nil
		return
	}

	if startPosRec == nil && rl.IsMouseButtonPressed(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), container) {
		startPosRec = &mousePos
		return
	}

	if startPosRec != nil {
		currentPosRec = &mousePos
	}
}
