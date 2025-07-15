package main

import (
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func UpperPart() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := lib.
			NewLayout().
			PositionRect(rect).
			Row().
			Gap(PANEL_GAP).
			Width(rect.Width,
				lib.ChildSize{SizeType: lib.SIZE_WEIGHT, Value: 1},
				lib.ChildSize{SizeType: lib.SIZE_ABSOLUTE, Value: SIDE_PANEL_WIDTH}).
			Height(rect.Height).
			Add(Canvas()).
			Add(RightPart())
		return layout.Draw, 0, 0
	}
}

func Canvas() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			rl.BeginMode2D(camera)
			intT := rect.ToInt32()
			rl.BeginScissorMode(intT.X, intT.Y, intT.Width, intT.Height)

			CanvasContent(rect)

			switch ui.SelectedTool {
			case lib.ToolSelection:
				Selection(rect)
			case lib.ToolRectangle:
				RectangleTool(rect)
			case lib.ToolText:
				TextTool(rect)
			}

			// if drawHighlight != nil {
			// 	drawHighlight(ui, c)
			// }
			if selectedLayer != nil {
				selectedLayer.DrawHighlight(ui, c)
			}

			rl.EndScissorMode()
			rl.EndMode2D()
		}, 0, 0
	}
}

func CanvasContent(rect rl.Rectangle) {
	for _, l := range layers {
		isClicked := l.DrawComponent(&ui, rl.GetScreenToWorld2D(rl.GetMousePosition(), camera))
		if isClicked {
			selectedLayer = l
		}

		if l.State() == components.STATE_HOT || l.State() == components.STATE_ACTIVE {
			drawHighlight = l.DrawHighlight
		}
	}
}
