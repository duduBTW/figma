package components

import (
	"github.com/dudubtw/figma/layout"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TimelinePanelLabel(text string) lib.Component {
	var fontSize int32 = 14
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y, fontSize, rl.White)
		}, 0, float32(fontSize)
	}
}
func (c *Components) TimelineRow(label string, inputs lib.Component, keyframes [][2]float32) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.Row(rect)
		row.Add(c.TimelinePanel(label, inputs))
		row.Add(c.TimelineFrames(c.ui, keyframes))
		return row.Draw, 0, row.Size.Height
	}
}
func (*Components) TimelinePanel(label string, inputs lib.Component) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.Panel(rect)
		row.Add(TimelinePanelLabel(label))
		row.Add(inputs)
		return row.Draw, 0, row.Size.Height
	}
}
func (c *Components) TimelineFrames(ui *lib.UIStruct, keyframes [][2]float32) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			y := rect.ToInt32().Y + 11
			rl.DrawLine(rect.ToInt32().X, y, rect.ToInt32().Width+rect.ToInt32().X, y, rl.Fade(rl.White, 0.32))

			if len(keyframes) > 0 {
				for _, keyframe := range keyframes {
					x := ui.GetXTimelineFrame(rect, keyframe[0])
					keyframeRect := rl.NewRectangle(x, float32(y), 10, 10)
					rl.DrawRectanglePro(
						keyframeRect,        // A 10×20 rectangle
						rl.NewVector2(5, 5), // Center of the rectangle
						45,
						rl.Blue,
					)

					// if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(keyframeRect.X-5, keyframeRect.Y-5, keyframeRect.Width, keyframeRect.Height)) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					// 	// clicked
					// }
				}
			}
		}, 0, 0
	}
}
