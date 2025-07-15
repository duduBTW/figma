package main

import (
	"math"
	"strconv"

	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/layout"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Timeline() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := lib.
			NewLayout().
			PositionRect(rect).
			Column().
			Padding(lib.NewPadding().All(12)).
			Width(rect.Width).
			Height(rect.Height,
				lib.ChildSize{SizeType: lib.SIZE_ABSOLUTE, Value: 32},
				lib.ChildSize{SizeType: lib.SIZE_ABSOLUTE, Value: 16},
				lib.ChildSize{SizeType: lib.SIZE_WEIGHT, Value: 1}).
			Add(TimelineUpperPart()).
			Add(TimelineHeader()).
			Add(TimelineBotttomPart())

		return func() {
			DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))
			layout.Draw()
		}, 0, 0
	}
}

func TimelineUpperPart() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.Row(rect)
		row.Add(TimelineControls())
		// row.Add(TimelineVisibleSlider())
		return row.Draw, 0, 0
	}
}

var visibleFrames = [2]int{0, 240}

func TimelineHeader() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.Row(rect)
		row.Add(Blank())
		row.Add(TimelineFrameController())
		return row.Draw, 0, 0
	}
}

func Blank() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {}, 0, 0
	}
}

func TimelineFrameController() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		totalFrames := int(float32((visibleFrames[1] - visibleFrames[0])) * 0.1)
		totalFrames2 := visibleFrames[1] - visibleFrames[0]

		frameWidth := rect.Width * 0.1
		ui.FrameWidth = rect.Width / float32(totalFrames2)

		smallestMousePos := math.MaxFloat64
		smallestFrame := 0
		for i := 0; i <= totalFrames2; i++ {
			var x int32 = int32(rect.X + ui.FrameWidth*float32(i) + 1)
			mousePos := math.Abs(float64(rl.GetMousePosition().X) - float64(x))
			if mousePos < smallestMousePos {
				smallestMousePos = mousePos
				smallestFrame = i
			}
		}

		rect2 := rect
		rect2.Height = 16
		interactable := components.NewInteractable("timeline-frame-controller", &ui)
		if interactable.Event(rl.GetMousePosition(), rect2) ||
			interactable.State() == components.STATE_ACTIVE {
			ui.SelectedFrame = smallestFrame
		}

		return func() {
			for i := 0; i <= 10; i++ {
				var x int32 = int32(rect.X + frameWidth*float32(i) + 1)
				label := strconv.Itoa(totalFrames*i) + "F"
				var fontSize int32 = 10
				rl.DrawText(label, lib.MinInt32((rect.ToInt32().X+rect.ToInt32().Width)-rl.MeasureText(label, fontSize), x+4), rect.ToInt32().Y, fontSize, rl.White)
			}
		}, 0, 0
	}
}

func TimelineBotttomPart() lib.Component {
	return func(current rl.Rectangle) (func(), float32, float32) {
		ui.ScrollTimeline()

		position := current
		position.Y -= ui.TimelineScroll

		layout := lib.
			NewLayout().
			PositionRect(position).
			Column().
			Gap(16).
			Width(current.Width)

		for _, layer := range layers {
			layout.Add(layer.DrawTimeline(&ui, c))
		}

		return func() {
			framesRect := rl.NewRectangle(current.X+280+12, current.Y, current.Width-280+12, current.Height)
			rl.DrawRectanglePro(framesRect, rl.NewVector2(0, 0), 0, rl.Fade(rl.Black, 0.4))

			rl.BeginScissorMode(current.ToInt32().X, current.ToInt32().Y, current.ToInt32().Width, current.ToInt32().Height)
			layout.Draw()
			rl.EndScissorMode()

			x := ui.GetXTimelineFrame(framesRect, float32(ui.SelectedFrame))
			rl.DrawLine(int32(x), framesRect.ToInt32().Y-10, int32(x), framesRect.ToInt32().Y+framesRect.ToInt32().Height, rl.Blue)
			rl.DrawRectanglePro(
				rl.NewRectangle(x, framesRect.Y-15, 11, 11), // A 10×20 rectangle
				rl.NewVector2(5, 5),                         // Center of the rectangle
				45,
				rl.Blue,
			)
		}, 0, 0
	}
}

func TimelineControls() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		buttonListLayout := lib.
			NewLayout().
			PositionRect(rect).
			Row().
			Gap(8).
			Add(PlayButton())
		return buttonListLayout.Draw, buttonListLayout.Size.Width, buttonListLayout.Size.Height
	}
}
func PlayButton() lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := c.Button("play-button", rl.NewVector2(rect.X, rect.Y), []lib.Component{})

		if button.Clicked {
			ui.TogglePlay()
		}

		return button.Draw, button.Rect.Width, button.Rect.Height
	}
}
