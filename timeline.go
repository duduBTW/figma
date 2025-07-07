package main

import (
	"math"
	"strconv"

	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/layer"
	"github.com/dudubtw/figma/layout"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Timeline() lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))

		padding := lib.Padding{}
		padding.All(12)
		layout := lib.NewConstrainedLayout(lib.ContrainedLayout{
			Padding:   padding,
			Contrains: rect,
			Direction: lib.DIRECTION_COLUMN,
			Gap:       12,
			ChildrenSize: []lib.ChildSize{
				{
					SizeType: lib.SIZE_ABSOLUTE,
					Value:    32,
				},
				{
					SizeType: lib.SIZE_WEIGHT,
					Value:    1,
				},
			},
		})

		layout.Add(TimelineUpperPart())
		layout.Add(TimelineBotttomPart())

		return layout.Draw, 0, 0
	}
}

func TimelineUpperPart() lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		row := layout.Timeline.NewTilelineRowLayout(rect)
		row.Add(TimelineControls())
		// row.Add(TimelineVisibleSlider())
		row.Draw()
	}
}

func TimelineBotttomPart() lib.ContrainedComponent {
	return func(current rl.Rectangle) {
		layout := lib.NewMixLayout(lib.PublicMixLayouyt{
			Direction: lib.DIRECTION_COLUMN,
			Gap:       16,
			InitialRect: lib.MixLayouytRect{
				Position: rl.NewVector2(current.X, current.Y),
				Width: lib.ContrainedSize{
					Value: current.Width,
					Contrains: []lib.ChildSize{
						{
							SizeType: lib.SIZE_WEIGHT,
							Value:    1,
						},
					},
				},
			},
		})
		for _, layer := range layers {
			layout.Add(layer.DrawTimeline(&ui, c))
		}
		layout.Draw()
	}
}

func R(height float32) lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			rl.DrawRectangle(rect.ToInt32().X, rect.ToInt32().Y, rect.ToInt32().Width, int32(height), rl.Red)
		}, 0, height
	}
}
func B(height float32) lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			rl.DrawRectangle(rect.ToInt32().X, rect.ToInt32().Y, rect.ToInt32().Width, int32(height), rl.Blue)
		}, 0, height
	}
}

func TimelineControls() lib.Component {
	buttonListLayout := lib.NewLayout(lib.PublicLayouyt{
		Direction: lib.DIRECTION_ROW,
		Gap:       8,
	}, rl.NewVector2(0, 0))

	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		buttonListLayout.SetPosition(avaliablePosition)
		buttonListLayout.Add(PlayButton())
		rect := buttonListLayout.Rect()
		rect.Width = 280
		return buttonListLayout.Draw, rect
	}
}
func PlayButton() lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		button := c.Button("play-button", avaliablePosition, []lib.Component{})

		if button.Clicked {
			ui.TogglePlay()
		}

		return button.Draw, button.Rect
	}
}

var visibleFrames = [2]int{0, 240}

func TimelineHeader(rect rl.Rectangle) {
	totalFrames := int(float32((visibleFrames[1] - visibleFrames[0])) * 0.1)
	frameWidth := rect.Width * 0.1

	for i := 0; i <= 10; i++ {
		var x int32 = int32(rect.X + frameWidth*float32(i) + 1)
		label := strconv.Itoa(totalFrames*i) + "F"
		var fontSize int32 = 10
		rl.DrawText(label, lib.MinInt32((rect.ToInt32().X+rect.ToInt32().Width)-rl.MeasureText(label, fontSize), x+4), rect.ToInt32().Y, fontSize, rl.White)
	}
}

func LayerPositionLine(containerRect rl.Rectangle, x layer.AnimatedProp, y layer.AnimatedProp) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		padding := lib.Padding{}
		padding.Start(32)
		padding.Top(12)
		row := lib.NewLayout(lib.PublicLayouyt{
			Padding:   padding,
			Direction: lib.DIRECTION_ROW,
			Gap:       16,
		}, avaliablePosition)

		row.Add(LayerPositionLineLabel(containerRect))
		row.Add(LayerPositionLineProperties(containerRect, x, y))
		return row.Draw, rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, row.Size.Width, row.Size.Height)
	}
}

func LayerPositionLineLabel(containerRect rl.Rectangle) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		text := "Position"
		fontSize := 14
		rect := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, containerRect.Width-100-32, 14)

		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y, int32(fontSize), rl.White)
		}, rect
	}
}

func LayerPositionLineProperties(containerRect rl.Rectangle, x layer.AnimatedProp, y layer.AnimatedProp) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		padding := lib.Padding{}
		padding.Start(32)
		column := lib.NewLayout(lib.PublicLayouyt{
			Padding:   padding,
			Direction: lib.DIRECTION_COLUMN,
			Gap:       12,
		}, avaliablePosition)

		if len(x.SortedKeyframes) > 0 {
			column.Add(LayerPositionLineProperty("x", x))
		}

		if len(y.SortedKeyframes) > 0 {
			column.Add(LayerPositionLineProperty("y", y))
		}

		return column.Draw, rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, column.Size.Width, column.Size.Height)
	}
}

func LayerPositionLineProperty(text string, animatedProp layer.AnimatedProp) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		var fontSize int32 = 14
		rect := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, float32(rl.MeasureText(text, fontSize)), float32(fontSize))

		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y, fontSize, rl.White)
		}, rect
	}
}

func LayerTimelineLabel(layer layer.Layer, index int, containerRect rl.Rectangle) lib.Component {
	return lib.NewComponent(func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		interactable := components.NewInteractable("layer"+strconv.Itoa(index), &ui)

		color := rl.Fade(rl.White, 0)

		isSelected := selectedLayer != nil && layer.GetElement().Id == selectedLayer.GetElement().Id
		if isSelected {
			color = rl.Fade(rl.White, 0.2)
		} else {
			switch interactable.State() {
			case components.STATE_HOT:
				color = rl.Fade(rl.White, 0.1)
			case components.STATE_ACTIVE:
				color = rl.Fade(rl.White, 0.2)
			}
		}

		padding := lib.Padding{}
		padding.All(4)
		box := c.Box(components.BoxProps{
			Padding:      padding,
			Rect:         rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, containerRect.Width, 0),
			Direction:    lib.DIRECTION_ROW,
			Children:     []lib.Component{LayerTimelineItemText(layer)},
			Color:        color,
			BorderRadius: 4,
		})

		if interactable.Event(rl.GetMousePosition(), box.Rect) {
			selectedLayer = layer
		}

		return box.Draw, box.Rect
	})
}

func LayerTimelineItemText(layer layer.Layer) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		textContent := layer.GetName()
		fontSize := 16

		return func() {
			rl.DrawText(textContent, int32(avaliablePosition.X), int32(avaliablePosition.Y), int32(fontSize), rl.White)
		}, rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, float32(rl.MeasureText(textContent, int32(fontSize))), float32(fontSize))
	}
}

func GetXTimelineFrame(frame int, frameWidth float32, timelineRect rl.Rectangle) float32 {
	return timelineRect.X + frameWidth*float32(frame) + 1
}

func TimelineFramesContent(rect rl.Rectangle) {
	DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(41, 41, 41, 255))

	totalFrames := visibleFrames[1] - visibleFrames[0]
	frameWidth := rect.Width / float32(totalFrames)
	smallestMousePos := math.MaxFloat64
	smallestFrame := 0
	for i := 0; i <= totalFrames; i++ {
		var x int32 = int32(rect.X + frameWidth*float32(i) + 1)
		mousePos := math.Abs(float64(rl.GetMousePosition().X) - float64(x))
		if mousePos < smallestMousePos {
			smallestMousePos = mousePos
			smallestFrame = i
		}
	}

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), rect) {
		ui.SelectedFrame = smallestFrame
	}

	x := GetXTimelineFrame(ui.SelectedFrame, frameWidth, rect)
	rl.DrawLine(int32(x), rect.ToInt32().Y-10, int32(x), rect.ToInt32().Y+rect.ToInt32().Height, rl.Blue)
	rl.DrawRectanglePro(
		rl.NewRectangle(x, rect.Y-15, 11, 11), // A 10×20 rectangle
		rl.NewVector2(5, 5),                   // Center of the rectangle
		45,
		rl.Blue,
	)

	if ui.IsPlaying {
		if visibleFrames[1] == ui.SelectedFrame {
			ui.SelectedFrame = visibleFrames[0]
		} else {
			ui.SelectedFrame++
		}
	}

	// layout := lib.NewLayout(lib.PublicLayouyt{
	// 	Direction: lib.DIRECTION_COLUMN,
	// 	Gap:       2,
	// }, rl.NewVector2(rect.X, rect.Y))

	// for _, l := range layers {
	// 	layout.Add(TimelineSpacer())
	// 	layout.Add(TimelinePropery(l.GetElement().Position.X, frameWidth, rect))
	// 	layout.Add(TimelinePropery(l.GetElement().Position.Y, frameWidth, rect))
	// }

	// layout.Draw()
}

func TimelineSpacer() lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		rect := rl.NewRectangle(avaliablePosition.X, avaliablePosition.X, 24+8, 24+8)
		return func() {
			rl.DrawRectanglePro(rect, rl.NewVector2(0, 0), 0, rl.Fade(rl.White, 0))
		}, rect
	}
}
func TimelinePropery(animatedProp layer.AnimatedProp, frameWidth float32, containerRect rl.Rectangle) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		if len(animatedProp.SortedKeyframes) == 0 {
			return func() {}, rl.NewRectangle(0, 0, 0, 0)
		}

		rect := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, containerRect.Width, 14+12+2)
		keyframes := animatedProp.SortedKeyframes

		return func() {
			y := int32(rect.Y + (rect.Height / 2))
			rl.DrawLine(rect.ToInt32().X, y, rect.ToInt32().X+rect.ToInt32().Width, y, rl.NewColor(68, 68, 68, 255))

			if len(keyframes) > 0 {
				for _, keyframe := range keyframes {
					x := GetXTimelineFrame(int(keyframe[0]), frameWidth, containerRect)
					keyframeRect := rl.NewRectangle(x, float32(y), 10, 10)
					rl.DrawRectanglePro(
						keyframeRect,        // A 10×20 rectangle
						rl.NewVector2(5, 5), // Center of the rectangle
						45,
						rl.Blue,
					)

					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(keyframeRect.X-5, keyframeRect.Y-5, keyframeRect.Width, keyframeRect.Height)) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
						// clicked
					}
				}
			}
		}, rect
	}
}
