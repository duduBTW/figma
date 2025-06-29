package main

import (
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func RightPart(rect rl.Rectangle) {
	layout := lib.NewConstrainedLayout(lib.ContrainedLayout{
		Direction: lib.DIRECTION_COLUMN,
		Contrains: rect,
		Gap:       PANEL_GAP,
		ChildrenSize: []lib.ChildSize{
			{
				SizeType: lib.SIZE_ABSOLUTE,
				Value:    TOOL_DOCK_HEIGHT,
			},
			{
				SizeType: lib.SIZE_WEIGHT,
				Value:    1,
			},
		},
	})
	layout.Add(ToolDock)
	layout.Add(PropertiesPanel)
	layout.Draw()
}

var tools = []lib.Tool{
	lib.ToolSelection,
	lib.ToolRectangle,
	lib.ToolText,
}

func ToolDock(rect rl.Rectangle) {
	DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))

	padding := lib.Padding{}
	padding.Axis(12, 8)
	layout := lib.NewLayout(lib.PublicLayouyt{
		Direction: lib.DIRECTION_ROW,
		Padding:   padding,
		Gap:       8,
	}, rl.NewVector2(rect.X, rect.Y))

	for _, tool := range tools {
		layout.Add(ToolButton(tool))
	}

	layout.Draw()
}

var toolButtonLabels = map[lib.Tool]string{
	lib.ToolSelection: "S",
	lib.ToolRectangle: "R",
	lib.ToolText:      "T",
}

func ToolButton(tool lib.Tool) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		button := c.Button("tool-"+string(tool), avaliablePosition, []lib.Component{Content(tool)})

		if button.Clicked {
			ui.SelectedTool = tool
		}

		return button.Draw, button.Rect
	}
}

func Content(tool lib.Tool) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		textContet := toolButtonLabels[tool]
		fontSize := 16
		rect := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, float32(rl.MeasureText(textContet, int32(fontSize))), float32(fontSize))
		return func() {
			rl.DrawText(textContet, int32(avaliablePosition.X), int32(avaliablePosition.Y), int32(fontSize), rl.White)
		}, rect
	}
}

func PropertiesPanel(rect rl.Rectangle) {
	DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))

	if selectedLayer == nil {
		return
	}

	rect.X += 12
	rect.Y += 12
	rect.Width -= 24
	rect.Height -= 24

	selectedLayer.DrawControls(&ui, rect, c)
}
