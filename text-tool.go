package main

import (
	"github.com/dudubtw/figma/layer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TextTool(container rl.Rectangle) {
	mousePos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), container) {
		newLayer := layer.NewText(NewLayerId(), mousePos)
		AppendLayer(&newLayer)
	}
}
