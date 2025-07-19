package main

import (
	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/components"
	rl "github.com/gen2brain/raylib-go/raylib"
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

var toolButtonIcons = map[app.Tool]app.IconName{
	app.ToolSelection: app.ICON_MOUSE_POINTER,
	app.ToolRectangle: app.ICON_SQUARE,
	app.ToolText:      app.ICON_TYPE,
	app.ToolImage:     app.ICON_IMAGE,
}

func ToolButton(tool app.Tool) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		buttonVariant := components.BUTTON_VARIANT_GHOST
		if tool == app.Apk.SelectedTool {
			buttonVariant = components.BUTTON_VARIANT_PRIMARY
		}

		button := components.Button("tool-"+string(tool), buttonVariant, rl.NewVector2(rect.X, rect.Y), []app.Component{components.Icon(toolButtonIcons[tool])})

		if button.Clicked {
			app.Apk.SelectedTool = tool
		}

		return button.Draw, button.Rect.Width, button.Rect.Height
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
