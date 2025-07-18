package main

import (
	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/components"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ncruces/zenity"
)

func RightPart() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.
			NewLayout().
			PositionRect(rect).
			Column().
			Gap(PANEL_GAP).
			Width(rect.Width).
			Height(rect.Height,
				app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: TOOL_DOCK_HEIGHT},
				app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1}).
			Add(ToolDock()).
			Add(PropertiesPanel())
		return layout.Draw, 0, 0
	}
}

var tools = []app.Tool{
	app.ToolSelection,
	app.ToolRectangle,
	app.ToolText,
	app.ToolImage,
}

func ToolDock() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.
			NewLayout().
			PositionRect(rect).
			Row().
			Padding(app.NewPadding().Axis(12, 8)).
			Gap(8)

		for _, tool := range tools {
			layout.Add(ToolButton(tool))
		}

		return func() {
			DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))
			layout.Draw()
		}, 0, 0
	}
}

var toolButtonLabels = map[app.Tool]string{
	app.ToolSelection: "S",
	app.ToolRectangle: "R",
	app.ToolText:      "T",
	app.ToolImage:     "I",
}

func ToolButton(tool app.Tool) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := components.Button("tool-"+string(tool), rl.NewVector2(rect.X, rect.Y), []app.Component{Content(tool)})

		if button.Clicked {
			app.Apk.SelectedTool = tool
		}

		// User just clicked on the image tool
		if app.Apk.SelectedTool == app.ToolImage && app.Apk.DroppingTexture == nil {
			dir, err := zenity.SelectFile(
				zenity.Title("Select the osu! lazer folder"),
				zenity.FileFilter{
					Patterns: []string{"*.png", "*.jpg", ".jpeg"},
				},
			)

			if err == nil {
				app.Apk.SetDroppingTexture(dir)
			} else {
				app.Apk.SelectedTool = app.ToolSelection
			}
		}

		return button.Draw, button.Rect.Width, button.Rect.Height
	}
}

func Content(tool app.Tool) app.Component {
	return func(avaliablePosition rl.Rectangle) (func(), float32, float32) {
		textContet := toolButtonLabels[tool]
		var fontSize int32 = 16
		return func() {
			rl.DrawText(textContet, int32(avaliablePosition.X), int32(avaliablePosition.Y), fontSize, rl.White)
		}, float32(rl.MeasureText(textContet, fontSize)), float32(fontSize)
	}
}

func PropertiesPanel() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {

		return func() {
			DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))
			if app.Apk.SelectedLayer == nil {
				return
			}

			app.Apk.SelectedLayer.DrawControls(rl.NewRectangle(rect.X+12, rect.Y+12, rect.Width-24, rect.Height-24))
		}, 0, 0
	}
}
