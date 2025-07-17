package main

import (
	"math"
	"strconv"

	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/fmath"
	"github.com/dudubtw/figma/layout"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Timeline() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.
			NewLayout().
			PositionRect(rect).
			Column().
			Padding(app.NewPadding().All(12)).
			Width(rect.Width).
			Height(rect.Height,
				app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: 32},
				app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: 16},
				app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1}).
			Add(TimelineUpperPart()).
			Add(TimelineHeader()).
			Add(TimelineBotttomPart())

		return func() {
			DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))
			layout.Draw()
		}, 0, 0
	}
}

func TimelineUpperPart() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.Row(rect)
		row.Add(TimelineControls())
		// row.Add(TimelineVisibleSlider())
		return row.Draw, 0, 0
	}
}

func TimelineHeader() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := layout.Timeline.Row(rect)
		row.Add(Blank())
		row.Add(TimelineFrameController())
		return row.Draw, 0, 0
	}
}

func Blank() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {}, 0, 0
	}
}

// TODO
// FIX THIS!!
func TimelineFrameController() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		totalFrames := int(float32((app.Apk.VisibleFrames[1] - app.Apk.VisibleFrames[0])) * 0.1)
		totalFrames2 := app.Apk.VisibleFrames[1] - app.Apk.VisibleFrames[0]

		frameWidth := rect.Width * 0.1
		app.Apk.FrameWidth = rect.Width / float32(totalFrames2)

		smallestMousePos := math.MaxFloat64
		smallestFrame := 0
		for i := 0; i <= totalFrames2; i++ {
			var x int32 = int32(rect.X + app.Apk.FrameWidth*float32(i) + 1)
			mousePos := math.Abs(float64(rl.GetMousePosition().X) - float64(x))
			if mousePos < smallestMousePos {
				smallestMousePos = mousePos
				smallestFrame = i
			}
		}

		rect2 := rect
		rect2.Height = 16
		interactable := app.NewInteractable("timeline-frame-controller")
		if interactable.Event(rl.GetMousePosition(), rect2) ||
			interactable.State() == app.STATE_ACTIVE {
			app.Apk.SelectedFrame = smallestFrame
		}

		return func() {
			for i := 0; i <= 10; i++ {
				var x int32 = int32(rect.X + frameWidth*float32(i) + 1)
				label := strconv.Itoa(totalFrames*i) + "F"
				var fontSize int32 = 10
				rl.DrawText(label, fmath.MinInt32((rect.ToInt32().X+rect.ToInt32().Width)-rl.MeasureText(label, fontSize), x+4), rect.ToInt32().Y, fontSize, rl.White)
			}
		}, 0, 0
	}
}

func TimelineBotttomPart() app.Component {
	return func(current rl.Rectangle) (func(), float32, float32) {
		app.Apk.ScrollTimeline()

		position := current
		position.Y -= app.Apk.TimelineScroll

		layout := app.
			NewLayout().
			PositionRect(position).
			Column().
			Gap(16).
			Width(current.Width)

		for _, layer := range app.Apk.Layers {
			layout.Add(layer.DrawTimeline())
		}

		return func() {
			framesRect := rl.NewRectangle(current.X+280+12, current.Y, current.Width-280-12, current.Height)
			rl.DrawRectanglePro(framesRect, rl.NewVector2(0, 0), 0, rl.Fade(rl.Black, 0.4))

			rl.BeginScissorMode(current.ToInt32().X, current.ToInt32().Y, current.ToInt32().Width, current.ToInt32().Height)
			layout.Draw()
			rl.EndScissorMode()

			x := app.Apk.GetXTimelineFrame(framesRect, float32(app.Apk.SelectedFrame))
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

func TimelineControls() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		buttonListLayout := app.
			NewLayout().
			PositionRect(rect).
			Row().
			Gap(8).
			Add(PlayButton())
		return buttonListLayout.Draw, buttonListLayout.Size.Width, buttonListLayout.Size.Height
	}
}
func PlayButton() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := components.Button("play-button", rl.NewVector2(rect.X, rect.Y), []app.Component{})

		if button.Clicked {
			app.Apk.TogglePlay()
		}

		return button.Draw, button.Rect.Width, button.Rect.Height
	}
}
