package main

import (
	"fmt"
	"strconv"

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

func RectangleTool() {
	RectangleSelectionActionHandler()

	if startPosRec == nil || currentPosRec == nil {
		return
	}

	rect := CalculateRectangle()

	// Draw the selection rectangle (semi-transparent fill + border)
	rl.DrawRectangleRec(rect, rl.Fade(rl.Black, 0.22))
	rl.DrawRectangleLinesEx(rect, 1, rl.Blue)
}

func RectangleSelectionActionHandler() {
	mousePos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
	if startPosRec != nil && rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		id := 0
		fmt.Println("len", len(layers))
		if len(layers) > 0 {
			newId, _ := strconv.Atoi(layers[len(layers)-1].GetElement().Id)
			id = newId
		}
		rect := CalculateRectangle()
		layers = append(layers, &layer.Rectangle{
			Width:  rect.ToInt32().Width,
			Height: rect.ToInt32().Height,
			Color:  rl.Fade(rl.Black, 0.32),
			Element: layer.Element{
				Id:       strconv.Itoa(id + 1),
				Position: rl.NewVector2(rect.X, rect.Y),
			},
		})
		startPosRec = nil
		currentPosRec = nil
		selectedValue = "selection"

		return
	}

	if startPosRec == nil && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		startPosRec = &mousePos
		return
	}

	if startPosRec != nil {
		currentPosRec = &mousePos
	}
}
