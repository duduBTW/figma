package components

import (
	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/layout"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TimelinePanelLabel(text string) app.Component {
	var fontSize int32 = 14
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y, fontSize, rl.White)
		}, 0, float32(fontSize)
	}
}
func TimelineRow(label string, inputs app.Component, keyframes [][2]float32) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.
			Row(rect).
			Add(TimelinePanel(label, inputs)).
			Add(TimelineFrames(keyframes))
		return row.Draw, 0, row.Size.Height
	}
}
func TimelinePanel(label string, inputs app.Component) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.Panel(rect)
		row.Add(TimelinePanelLabel(label))
		row.Add(inputs)
		return row.Draw, 0, row.Size.Height
	}
}
func TimelineFrames(keyframes [][2]float32) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			y := rect.ToInt32().Y + 11
			rl.DrawLine(rect.ToInt32().X, y, rect.ToInt32().Width+rect.ToInt32().X, y, rl.Fade(rl.White, 0.32))

			if len(keyframes) > 0 {
				for _, keyframe := range keyframes {
					x := app.Apk.State.GetXTimelineFrame(rect, keyframe[0])
					keyframeRect := rl.NewRectangle(x, float32(y), 10, 10)
					rl.DrawRectanglePro(
						keyframeRect,        // A 10×20 rectangle
						rl.NewVector2(5, 5), // Center of the rectangle
						45,
						rl.Blue,
					)

					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(keyframeRect.X-5, keyframeRect.Y-5, keyframeRect.Width, keyframeRect.Height)) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
						// ui.SelectedKeyframe = app.SelectedKeyframe{LayerId: selectedLayer.GetElement().Id, Keyframe: keyframe}
					}
				}
			}
		}, 0, 0
	}
}

func TimelinePanelTitle(text string, layer app.Layer) app.Component {
	var fontSize int32 = 16
	var height = float32(fontSize)
	return func(rect rl.Rectangle) (func(), float32, float32) {
		width := float32(rl.MeasureText(text, fontSize))
		interractable := app.NewInteractable(layer.GetName() + "panel-item")
		if interractable.Event(rl.GetMousePosition(), rl.NewRectangle(rect.X, rect.Y, width, height)) {
			app.Apk.SelectedLayer = layer
		}

		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y, fontSize, rl.White)
		}, width, height
	}
}
