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
		totalFrames := int(float32((app.Apk.Workplace.VisibleFrames[1] - app.Apk.Workplace.VisibleFrames[0])) * 0.1)
		totalFrames2 := app.Apk.Workplace.VisibleFrames[1] - app.Apk.Workplace.VisibleFrames[0]

		frameWidth := rect.Width * 0.1
		app.Apk.Workplace.FrameWidth = rect.Width / float32(totalFrames2)

		smallestMousePos := math.MaxFloat64
		smallestFrame := 0
		for i := 0; i <= totalFrames2; i++ {
			var x int32 = int32(rect.X + app.Apk.Workplace.FrameWidth*float32(i) + 1)
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
			app.Apk.Workplace.SelectedFrame = smallestFrame
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

func FramesRect(rect rl.Rectangle) rl.Rectangle {
	return rl.NewRectangle(rect.X+280+12, rect.Y, rect.Width-280-12, rect.Height)
}

func TimelineBotttomPart() app.Component {
	if app.Apk.Workplace.TimelineCurveSelected {
		return TimelineBotttomPartCurve()
	}

	return TimelineBotttomPartFrames()
}
func TimelineBotttomPartCurve() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			DrawSelectedFrameIndicator(rect)

			if app.Apk.Workplace.SelectedLayer == nil {
				return
			}

			animatedProp := app.Apk.Workplace.SelectedAnimatedProp
			if len(animatedProp.SortedKeyframes) < 1 {
				return
			}

			framesRect := FramesRect(rect)
			framesRect.Width -= 24
			framesRect.Height -= 24
			framesRect.X += 12
			framesRect.Y += 12
			positions := []rl.Vector2{}
			var min float64 = 9999
			var max float64 = -9999
			for _, frame := range animatedProp.SortedKeyframes {
				value := float64(frame[1])
				if value > max {
					max = value
				}

				if value < min {
					min = value
				}
			}

			y := app.NewLinear().Domain(float64(framesRect.Y), float64(framesRect.Y+framesRect.Height)).Range(min, max)
			mousePos := rl.GetMousePosition()

			app.Apk.Workplace.FramesRect = framesRect

			for index, frame := range animatedProp.SortedKeyframes {
				pos := GetCirclePosition(frame, y)
				positions = append(positions, pos)

				interactable := app.NewInteractable("corner-control" + strconv.Itoa(index))
				isClicked := interactable.Event(mousePos, rl.NewRectangle(pos.X-4, pos.Y-4, 8, 8))
				if isClicked || interactable.State() == app.STATE_ACTIVE {
					app.Apk.Workplace.SelectedLayer.GetElement().Position.X.Set(frame[0], float32(y.Scale(float64(mousePos.Y))))
				}

				rl.DrawCircle(int32(pos.X), int32(pos.Y), 4, rl.Red)

				if index+1 != len(animatedProp.SortedKeyframes) {
					nextFrame := animatedProp.SortedKeyframes[index+1]
					nextFramePos := GetCirclePosition(nextFrame, y)
					positions = append(
						positions,
						// Start handle
						DrawHandler("start-curve-control", frame, index, pos, animatedProp.KeyFrameCurveStart, 16, pos.X, nextFramePos.X),
						// End handler
						DrawHandler("end-curve-control", nextFrame, index, nextFramePos, animatedProp.KeyFrameCurveEnd, -16, pos.X, nextFramePos.X),
					)
				}
			}

			rl.DrawSplineBezierCubic(positions, 4, rl.Blue)
		}, 0, 0
	}
}

func GetCirclePosition(frame [2]float32, y *app.Linear) rl.Vector2 {
	return rl.NewVector2(app.Apk.Workplace.GetXTimelineFrame(app.Apk.Workplace.FramesRect, frame[0]), float32(y.Invert(float64(frame[1]))))
}

func DrawHandler(id string, frame [2]float32, index int, circlePosition rl.Vector2, keyFrameCurves map[float32]rl.Vector2, positionOffset, startLimit, endLimit float32) rl.Vector2 {
	offset := keyFrameCurves[frame[0]]
	positionWithOffset := rl.NewVector2(circlePosition.X+offset.X, circlePosition.Y+offset.Y)

	mousePos := rl.GetMousePosition()
	startX := int32(positionWithOffset.X + positionOffset)
	startInteractable := app.NewInteractable(id + strconv.Itoa(index))
	isClicked := startInteractable.Event(mousePos, rl.NewRectangle(float32(startX-4), positionWithOffset.Y-4, 8, 8))
	if isClicked || startInteractable.State() == app.STATE_ACTIVE {
		keyFrameCurve := keyFrameCurves[frame[0]]
		x := keyFrameCurve.X
		if mousePos.X >= startLimit && mousePos.X <= endLimit {
			x = mousePos.X - circlePosition.X - positionOffset
		}

		y := keyFrameCurve.Y
		if mousePos.Y >= app.Apk.Workplace.FramesRect.Y && mousePos.Y <= app.Apk.Workplace.FramesRect.Y+app.Apk.Workplace.FramesRect.Height {
			y = mousePos.Y - circlePosition.Y
		}

		keyFrameCurves[frame[0]] = rl.NewVector2(x, y)
	}

	rl.DrawCircle(startX, int32(positionWithOffset.Y), 2, rl.Yellow)
	rl.DrawLine(int32(circlePosition.X), int32(circlePosition.Y), startX, int32(positionWithOffset.Y), rl.Yellow)

	return positionWithOffset
}

func TimelineBotttomPartFrames() app.Component {
	return func(current rl.Rectangle) (func(), float32, float32) {
		app.Apk.Workplace.ScrollTimeline()

		position := current
		position.Y -= app.Apk.Workplace.TimelineScroll

		layout := app.
			NewLayout().
			PositionRect(position).
			Column().
			Gap(16).
			Width(current.Width)

		for _, layer := range app.Apk.Workplace.Layers {
			layout.Add(layer.DrawTimeline())
		}

		return func() {
			rl.BeginScissorMode(current.ToInt32().X, current.ToInt32().Y, current.ToInt32().Width, current.ToInt32().Height)
			layout.Draw()
			rl.EndScissorMode()
			DrawSelectedFrameIndicator(current)
		}, 0, 0
	}
}

func DrawSelectedFrameIndicator(rect rl.Rectangle) {
	framesRect := FramesRect(rect)
	rl.DrawRectanglePro(framesRect, rl.NewVector2(0, 0), 0, rl.Fade(rl.Black, 0.4))

	x := app.Apk.Workplace.GetXTimelineFrame(framesRect, float32(app.Apk.Workplace.SelectedFrame))
	rl.DrawLine(int32(x), framesRect.ToInt32().Y-10, int32(x), framesRect.ToInt32().Y+framesRect.ToInt32().Height, rl.Blue)
	rl.DrawRectanglePro(
		rl.NewRectangle(x, framesRect.Y-15, 11, 11), // A 10×20 rectangle
		rl.NewVector2(5, 5),                         // Center of the rectangle
		45,
		rl.Blue,
	)
}

func TimelineControls() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		buttonListLayout := app.
			NewLayout().
			PositionRect(rect).
			Row().
			Gap(8).
			Add(PlayButton()).
			Add(CurveButton())
		return buttonListLayout.Draw, buttonListLayout.Size.Width, buttonListLayout.Size.Height
	}
}

func CanvasButtonVariant() components.ButtonVariant {
	if app.Apk.Workplace.TimelineCurveSelected {
		return components.BUTTON_VARIANT_PRIMARY
	}

	return components.BUTTON_VARIANT_GHOST
}
func CurveButton() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := components.Button("curve-button", CanvasButtonVariant(), rl.NewVector2(rect.X, rect.Y), []app.Component{components.Icon(app.ICON_CHART_SPLINE)})
		if button.Clicked {
			app.Apk.Workplace.ToggleCurveSelected()
		}

		return button.Draw, button.Rect.Width, button.Rect.Height
	}
}

func PlayButtonIcon() app.IconName {
	if app.Apk.Workplace.IsPlaying {
		return app.ICON_PAUSE
	}

	return app.ICON_PLAY
}
func PlayButton() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := components.Button("play-button", components.BUTTON_VARIANT_GHOST, rl.NewVector2(rect.X, rect.Y), []app.Component{components.Icon(PlayButtonIcon())})
		if button.Clicked {
			app.Apk.Workplace.TogglePlay()
		}

		return button.Draw, button.Rect.Width, button.Rect.Height
	}
}
