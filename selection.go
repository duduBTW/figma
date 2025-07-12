package main

import (
	"math"

	"github.com/dudubtw/figma/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var startPos *rl.Vector2
var currentPos *rl.Vector2

var movingSartPos *rl.Vector2
var movingCurrentPos *rl.Vector2
var dragAnchor = rl.Rectangle{}

// -1 = horizontal 1 = vertical
var direction = 0

func Selection(canvas rl.Rectangle) {

	if rl.IsKeyReleased(rl.KeyLeftShift) {
		direction = 0
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}

	if !rl.CheckCollisionPointRec(rl.GetMousePosition(), canvas) {
		return
	}

	if selectedLayer != nil && rl.IsMouseButtonDown(rl.MouseButtonLeft) && startPos == nil {
		if selectedLayer.State() == components.STATE_INITIAL {
			selectedLayer = nil
			return
		}

		SelectionActionHandler(&movingSartPos, &movingCurrentPos)

		if movingSartPos != nil && movingCurrentPos != nil {
			element := selectedLayer.GetElement()
			x := movingCurrentPos.X - movingSartPos.X
			y := movingCurrentPos.Y - movingSartPos.Y

			if direction == 0 && rl.IsKeyDown(rl.KeyLeftShift) && (int(x) != 0 || int(y) > 0) {
				angleRadians := math.Atan2(float64(movingSartPos.Y-movingCurrentPos.Y), float64(movingSartPos.X-movingCurrentPos.X))
				if angleRadians < 0 {
					angleRadians += math.Pi * 2
				}
				angleRadians += math.Pi / 4
				if angleRadians < math.Pi/2 {
					direction = -1
				} else if angleRadians < math.Pi {
					direction = 1
				} else if angleRadians < 3*math.Pi/2 {
					direction = -1
				} else {
					direction = 1
				}

				if direction == -1 {
					rl.SetMouseCursor(rl.MouseCursorResizeEW)
				}
				if direction == 1 {
					rl.SetMouseCursor(rl.MouseCursorResizeNS)
				}
			}

			if direction != 1 && x != 0 {
				element.Position.X.Set(dragAnchor.X+x, &ui)
			}

			if direction != -1 && y != 0 {
				element.Position.Y.Set(dragAnchor.Y+y, &ui)
			}

		}
		return
	}

	SelectionActionHandler(&startPos, &currentPos)

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

func SelectionActionHandler(start, current **rl.Vector2) {
	mousePos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)

	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		*start = nil
		*current = nil
		return
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		*start = &mousePos
		*current = *start

		if selectedLayer != nil {
			dragAnchor = selectedLayer.Rect(ui.SelectedFrame)
		}
		return
	}

	if *start != nil {
		*current = &mousePos
	}
}
