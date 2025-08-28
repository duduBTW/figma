package main

import (
	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/components"
	ds "github.com/dudubtw/figma/design-system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func LeftPart() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.
			NewLayout().
			PositionRect(rect).
			Column().
			Gap(PANEL_GAP).
			Width(rect.Width).
			Height(rect.Height,
				app.ChildSize{Value: 48, SizeType: app.SIZE_ABSOLUTE},
				app.ChildSize{Value: 1, SizeType: app.SIZE_WEIGHT}).
			Add(LeftPartTabs()).
			Add(Canvas())

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}

func LeftPartTabs() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.NewLayout().
			Row().
			PositionRect(rect).
			Width(rect.Width).
			Height(rect.Height).
			Padding(app.NewPadding().Start(ds.SPACING_2).Top(ds.SPACING_2)).
			Add(components.OpenTabs())

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}

func Canvas() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			rl.BeginMode2D(camera)
			intT := rect.ToInt32()
			rl.BeginScissorMode(intT.X, intT.Y, intT.Width, intT.Height)

			CanvasContent(rect)

			switch app.Apk.Workplace.SelectedTool {
			case app.ToolSelection:
				Selection(rect)
			case app.ToolRectangle:
				RectangleTool(rect)
			case app.ToolText:
				TextTool(rect)
			case app.ToolImage:
				ImageTool(rect)
			}

			if app.Apk.Workplace.DrawFrameHighlight != nil {
				app.Apk.Workplace.DrawFrameHighlight()
			}
			if app.Apk.Workplace.SelectedLayer != nil {
				app.Apk.Workplace.SelectedLayer.DrawHighlight()
			}

			rl.EndScissorMode()
			rl.EndMode2D()
		}, 0, 0
	}
}

func CanvasContent(rect rl.Rectangle) {
	for _, l := range app.Apk.Workplace.Layers {
		isClicked := l.DrawComponent(rl.GetScreenToWorld2D(rl.GetMousePosition(), camera), rect)
		if isClicked {
			app.Apk.Workplace.SetSelectedLayer(l)
		}

		if l.State() == app.STATE_HOT || l.State() == app.STATE_ACTIVE {
			app.Apk.Workplace.DrawFrameHighlight = l.DrawHighlight
		}
	}
}
