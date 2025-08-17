package main

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Canvas() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			rl.BeginMode2D(camera)
			intT := rect.ToInt32()
			rl.BeginScissorMode(intT.X, intT.Y, intT.Width, intT.Height)

			CanvasContent(rect)

			switch app.Apk.SelectedTool {
			case app.ToolSelection:
				Selection(rect)
			case app.ToolRectangle:
				RectangleTool(rect)
			case app.ToolText:
				TextTool(rect)
			case app.ToolImage:
				ImageTool(rect)
			}

			if app.Apk.DrawFrameHighlight != nil {
				app.Apk.DrawFrameHighlight()
			}
			if app.Apk.SelectedLayer != nil {
				app.Apk.SelectedLayer.DrawHighlight()
			}

			rl.EndScissorMode()
			rl.EndMode2D()
		}, 0, 0
	}
}

func CanvasContent(rect rl.Rectangle) {
	for _, l := range app.Apk.Layers {
		isClicked := l.DrawComponent(rl.GetScreenToWorld2D(rl.GetMousePosition(), camera), rect)
		if isClicked {
			app.Apk.SetSelectedLayer(l)
		}

		if l.State() == app.STATE_HOT || l.State() == app.STATE_ACTIVE {
			app.Apk.DrawFrameHighlight = l.DrawHighlight
		}
	}
}
