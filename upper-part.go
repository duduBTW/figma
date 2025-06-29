package main

import (
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func UpperPart(rect rl.Rectangle) {
	layout := lib.NewConstrainedLayout(lib.ContrainedLayout{
		Direction: lib.DIRECTION_ROW,
		Gap:       PANEL_GAP,
		Contrains: rect,
		ChildrenSize: []lib.ChildSize{
			{
				SizeType: lib.SIZE_WEIGHT,
				Value:    1,
			},
			{
				SizeType: lib.SIZE_ABSOLUTE,
				Value:    SIDE_PANEL_WIDTH,
			},
		},
	})
	layout.Add(Canvas)
	layout.Add(RightPart)
	layout.Draw()
}

func Canvas(rect rl.Rectangle) {
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

	if selectedLayer != nil {
		selectedLayer.DrawHighlight(ui, c)
	}

	rl.EndScissorMode()
	rl.EndMode2D()
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
