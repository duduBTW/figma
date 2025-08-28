package main

import (
	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/layer"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ncruces/zenity"
)

func ImageTool(rect rl.Rectangle) {
	// User just clicked on the image tool
	if app.Apk.Workplace.DroppingImg == nil {
		dir, err := zenity.SelectFile(
			zenity.Title("Select an image"),
			zenity.FileFilter{
				Patterns: []string{"*.png", "*.jpg", ".jpeg"},
			},
		)

		if err != nil {
			// Goes back to the selection tool if nothing was selected
			app.Apk.Workplace.SelectedTool = app.ToolSelection
			return
		}

		app.Apk.Workplace.SetDroppingTexture(dir)
	}

	mousePoint := rl.GetMousePosition()
	if app.Apk.Workplace.DroppingImg == nil || !rl.CheckCollisionPointRec(mousePoint, rect) {
		return
	}

	texture := app.Apk.Workplace.DroppingImg.Texture
	textureRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	targetRect := rl.NewRectangle(mousePoint.X+8, mousePoint.Y+8, 80, 80)
	rl.DrawTexturePro(texture, textureRect, targetRect, rl.Vector2{}, 0, rl.White)

	if !rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		return
	}

	index := 0
	for _, l := range app.Apk.Workplace.Layers {
		_, isImage := l.(*layer.Image)
		if isImage {
			index++
		}
	}

	newLayer := layer.NewImage(app.Apk.Workplace.NewLayerId(), mousePoint, app.Apk.Workplace.DroppingImg.Path, index)
	app.Apk.Workplace.AppendLayer(&newLayer)
	app.Apk.Workplace.DroppingImg = nil
}
