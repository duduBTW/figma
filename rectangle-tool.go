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
		id := 0
		if len(layers) > 0 {
			newId, _ := strconv.Atoi(layers[len(layers)-1].GetElement().Id)
			id = newId
		}
		rect := CalculateRectangle()
		fmt.Println(rect)
		index := 0
		for _, l := range layers {
			_, isRec := l.(*layer.Rectangle)
			if isRec {
				index++
			}
		}

		newLayer := layer.NewRectangle(strconv.Itoa(id+1), rect, index)

		layers = append(layers, &newLayer)
		startPosRec = nil
		currentPosRec = nil
		selectedLayer = &newLayer
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
