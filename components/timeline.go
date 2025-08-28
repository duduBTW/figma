package components

import (
	"github.com/dudubtw/figma/app"
	ds "github.com/dudubtw/figma/design-system"
	"github.com/dudubtw/figma/layout"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TimelinePanelLabel(text string) app.Component {
	return Typography(text, ds.FONT_SIZE, ds.FONT_WEIGHT_MEDIUM, ds.T2_COLOR_CONTENT)
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
					layerId := app.Apk.Workplace.SelectedLayer.GetElement().Id
					x := app.Apk.Workplace.GetXTimelineFrame(rect, keyframe)
					keyframeRect := rl.NewRectangle(x, float32(y), 10, 10)
					color := rl.Blue
					if app.Apk.Workplace.SelectedKeyframe.Keyframe == keyframe && app.Apk.Workplace.SelectedKeyframe.LayerId == layerId {
						color = rl.White
					}
					rl.DrawRectanglePro(
						keyframeRect,        // A 10×20 rectangle
						rl.NewVector2(5, 5), // Center of the rectangle
						45,
						color,
					)

					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(keyframeRect.X-5, keyframeRect.Y-5, keyframeRect.Width, keyframeRect.Height)) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
						app.Apk.Workplace.SelectedKeyframe = app.SelectedKeyframe{LayerId: layerId, Keyframe: keyframe}
						app.Apk.Workplace.SelectedAnimatedProp = &animatedProp
					}
				}
			}
		}, 0, 0
	}
}

func TimelinePanelTitle(iconName app.IconName, text string, layer app.Layer) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.NewLayout().
			Row().
			PositionRect(rect).
			Gap(ds.SPACING_1).
			Padding(app.NewPadding().Bottom(ds.SPACING_2)).
			Width(rect.Width).
			Add(Icon(iconName)).
			Add(Typography(text, ds.FONT_SIZE_LG, ds.FONT_WEIGHT_MEDIUM, ds.T2_COLOR_CONTENT_ACCENT))

		interractable := app.NewInteractable(layer.GetName() + "panel-item")
		elementRect := rl.NewRectangle(rect.X, rect.Y, rect.Width, layout.Size.Height)
		if interractable.Event(rl.GetMousePosition(), elementRect) {
			app.Apk.Workplace.SetSelectedLayer(layer)
		}

		return layout.Draw, elementRect.Width, elementRect.Height
	}
}
