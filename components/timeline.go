package components

import (
	"github.com/dudubtw/figma/layout"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TimelinePanelLabel(text string) lib.MixComponent {
	var fontSize int32 = 14
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y, fontSize, rl.White)
		}, 0, float32(fontSize)
	}
}
func (c *Components) TimelineRow(label string, inputs lib.MixComponent) lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.Row(rect)
		row.Add(c.TimelinePanel(label, inputs))
		row.Add(c.TimelineFrames())
		return row.Draw, 0, row.CurrentRect.Height
	}
}
func (*Components) TimelinePanel(label string, inputs lib.MixComponent) lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.Panel(rect)
		row.Add(TimelinePanelLabel(label))
		row.Add(inputs)
		return row.Draw, 0, row.CurrentRect.Height
	}
}
func (c *Components) TimelineFrames() lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			y := rect.ToInt32().Y + 11
			rl.DrawLine(rect.ToInt32().X, y, rect.ToInt32().Width+rect.ToInt32().X, y, rl.White)
		}, 0, 0
	}
}
