package main

import (
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func RightPart() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := lib.
			NewLayout().
			PositionRect(rect).
			Column().
			Gap(PANEL_GAP).
			Width(rect.Width).
			Height(rect.Height,
				lib.ChildSize{SizeType: lib.SIZE_ABSOLUTE, Value: TOOL_DOCK_HEIGHT},
				lib.ChildSize{SizeType: lib.SIZE_WEIGHT, Value: 1}).
			Add(ToolDock()).
			Add(PropertiesPanel())
		return layout.Draw, 0, 0
	}
}

var tools = []lib.Tool{
	lib.ToolSelection,
	lib.ToolRectangle,
	lib.ToolText,
}

func ToolDock() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))
		layout := lib.
			NewLayout().
			PositionRect(rect).
			Row().
			Padding(lib.NewPadding().Axis(12, 8)).
			Gap(8)

		for _, tool := range tools {
			layout.Add(ToolButton(tool))
		}

		return layout.Draw, 0, 0
	}
}

var toolButtonLabels = map[lib.Tool]string{
	lib.ToolSelection: "S",
	lib.ToolRectangle: "R",
	lib.ToolText:      "T",
}

func ToolButton(tool lib.Tool) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := c.Button("tool-"+string(tool), rl.NewVector2(rect.X, rect.Y), []lib.Component{Content(tool)})

		if button.Clicked {
			ui.SelectedTool = tool
		}

		return button.Draw, button.Rect.Width, button.Rect.Height
	}
}

func Content(tool lib.Tool) lib.Component {
	return func(avaliablePosition rl.Rectangle) (func(), float32, float32) {
		textContet := toolButtonLabels[tool]
		var fontSize int32 = 16
		return func() {
			rl.DrawText(textContet, int32(avaliablePosition.X), int32(avaliablePosition.Y), fontSize, rl.White)
		}, float32(rl.MeasureText(textContet, fontSize)), float32(fontSize)
	}
}

func PropertiesPanel() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))

		rect.X += 12
		rect.Y += 12
		rect.Width -= 24
		rect.Height -= 24

		return func() {
			if selectedLayer == nil {
				return
			}

			selectedLayer.DrawControls(&ui, rect, c)
		}, 0, 0
	}
}
