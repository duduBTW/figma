package main

import (
	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/layer"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ncruces/zenity"
)

func ImageTool(rect rl.Rectangle) {
	// User just clicked on the image tool
	if app.Apk.DroppingTexture == nil {
		dir, err := zenity.SelectFile(
			zenity.Title("Select an image"),
			zenity.FileFilter{
				Patterns: []string{"*.png", "*.jpg", ".jpeg"},
			},
		)

		if err != nil {
			// Goes back to the selection tool if nothing was selected
			app.Apk.SelectedTool = app.ToolSelection
			return
		}

		app.Apk.SetDroppingTexture(dir)
	}

	mousePoint := rl.GetMousePosition()
	if app.Apk.DroppingTexture == nil || !rl.CheckCollisionPointRec(mousePoint, rect) {
		return
	}

	textureRect := rl.NewRectangle(0, 0, float32(app.Apk.DroppingTexture.Width), float32(app.Apk.DroppingTexture.Height))
	targetRect := rl.NewRectangle(mousePoint.X+8, mousePoint.Y+8, 80, 80)
	rl.DrawTexturePro(*app.Apk.DroppingTexture, textureRect, targetRect, rl.Vector2{}, 0, rl.White)

	if !rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		return
	}

	index := 0
	for _, l := range app.Apk.Layers {
		_, isImage := l.(*layer.Image)
		if isImage {
			index++
		}
	}

	newLayer := layer.NewImage(app.Apk.NewLayerId(), mousePoint, *app.Apk.DroppingTexture, index)
	app.Apk.AppendLayer(&newLayer)
	app.Apk.DroppingTexture = nil
}
