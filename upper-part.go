package main

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func UpperPart() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.
			NewLayout().
			PositionRect(rect).
			Row().
			Gap(PANEL_GAP).
			Width(rect.Width,
				app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1},
				app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: SIDE_PANEL_WIDTH}).
			Height(rect.Height).
			Add(Canvas()).
			Add(RightPart())
		return layout.Draw, 0, 0
	}
}

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
			}

			// if drawHighlight != nil {
			// 	drawHighlight(ui, c)
			// }
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
		isClicked := l.DrawComponent(rl.GetScreenToWorld2D(rl.GetMousePosition(), camera))
		if isClicked {
			app.Apk.SelectedLayer = l
		}

		// TODO REIMPLEMENT
		// if l.State() == components.STATE_HOT || l.State() == components.STATE_ACTIVE {
		// drawHighlight = l.DrawHighlight
		// }
	}
}
