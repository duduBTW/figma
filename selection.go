package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var startPos *rl.Vector2
var currentPos *rl.Vector2

func Selection() {
	SelectionActionHandler()

	if startPos == nil || currentPos == nil {
		return
	}

	width := currentPos.X - startPos.X
	height := currentPos.Y - startPos.Y

	// Handle dragging in any direction
	rect := rl.NewRectangle(startPos.X, startPos.Y, width, height)
	if width < 0 {
		rect.X += width
		rect.Width = -width
	}
	if height < 0 {
		rect.Y += height
		rect.Height = -height
	}

	// Draw the selection rectangle (semi-transparent fill + border)
	rl.DrawRectangleRec(rect, rl.Fade(rl.SkyBlue, 0.3))
	rl.DrawRectangleLinesEx(rect, 1, rl.Blue)
}

func SelectionActionHandler() {
	mousePos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
	if startPos != nil && rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		startPos = nil
		currentPos = nil
		return
	}

	if startPos == nil && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		startPos = &mousePos
		return
	}

	if startPos != nil {
		currentPos = &mousePos
	}
}
