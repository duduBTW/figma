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
func TimelineRow(label string, inputs app.Component, animatedProp app.AnimatedProp) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.
			Row(rect).
			Add(TimelinePanel(label, inputs)).
			Add(TimelineFrames(animatedProp))
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
func TimelineFrames(animatedProp app.AnimatedProp) app.Component {
	keyframes := animatedProp.SortedKeyframesTimeline()
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			y := rect.ToInt32().Y + 11
			rl.DrawLine(rect.ToInt32().X, y, rect.ToInt32().Width+rect.ToInt32().X, y, rl.Fade(rl.White, 0.32))

			if len(keyframes) > 0 {
				for _, keyframe := range keyframes {
					layerId := app.Apk.SelectedLayer.GetElement().Id
					x := app.Apk.State.GetXTimelineFrame(rect, keyframe)
					keyframeRect := rl.NewRectangle(x, float32(y), 10, 10)
					color := rl.Blue
					if app.Apk.SelectedKeyframe.Keyframe == keyframe && app.Apk.SelectedKeyframe.LayerId == layerId {
						color = rl.White
					}
					rl.DrawRectanglePro(
						keyframeRect,        // A 10×20 rectangle
						rl.NewVector2(5, 5), // Center of the rectangle
						45,
						color,
					)

					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(keyframeRect.X-5, keyframeRect.Y-5, keyframeRect.Width, keyframeRect.Height)) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
						app.Apk.SelectedKeyframe = app.SelectedKeyframe{LayerId: layerId, Keyframe: keyframe}
						app.Apk.SelectedAnimatedProp = &animatedProp
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
			app.Apk.SetSelectedLayer(layer)
		}

		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y, fontSize, rl.White)
		}, width, height
	}
}
