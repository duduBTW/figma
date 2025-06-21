package main

import (
	"math"
	"strconv"

	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/layer"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Timeline(rect rl.Rectangle) {
	DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))

	padding := lib.Padding{}
	padding.All(12)
	layout := lib.NewConstrainedLayout(lib.ContrainedLayout{
		Padding:   padding,
		Contrains: rect,
		Direction: lib.DIRECTION_ROW,
		Gap:       12,
		ChildrenSize: []lib.ChildSize{
			{
				SizeType: lib.SIZE_ABSOLUTE,
				Value:    280,
			},
			{
				SizeType: lib.SIZE_WEIGHT,
				Value:    1,
			},
		},
	})

	layout.Add(LayersPanel)
	layout.Add(TimelineFrames)
	layout.Draw()
}

func LayersPanel(rect rl.Rectangle) {
	padding := lib.Padding{}
	padding.Top(20)
	layout := lib.NewLayout(lib.PublicLayouyt{
		Padding:   padding,
		Direction: lib.DIRECTION_COLUMN,
		Gap:       2,
	}, rl.NewVector2(rect.X, rect.Y))

	for index, l := range layers {
		layout.Add(LayerTimelineItem(l, index, rect))
	}

	layout.Draw()
}

func LayerTimelineItem(layer layer.Layer, index int, containerRect rl.Rectangle) lib.Component {
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
		padding.All(8)
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

var visibleFrames = [2]int{0, 240}

func TimelineFrames(rect rl.Rectangle) {
	layout := lib.NewConstrainedLayout(lib.ContrainedLayout{
		Contrains: rect,
		Direction: lib.DIRECTION_COLUMN,
		ChildrenSize: []lib.ChildSize{
			{
				SizeType: lib.SIZE_ABSOLUTE,
				Value:    20,
			},
			{
				SizeType: lib.SIZE_WEIGHT,
				Value:    1,
			},
		},
	})
	layout.Add(TimelineFramesHeader)
	layout.Add(TimelineFramesContent)
	layout.Draw()
}
func TimelineFramesHeader(rect rl.Rectangle) {
	totalFrames := int(float32((visibleFrames[1] - visibleFrames[0])) * 0.1)
	frameWidth := rect.Width * 0.1

	for i := 0; i <= 10; i++ {
		var x int32 = int32(rect.X + frameWidth*float32(i) + 1)
		label := strconv.Itoa(totalFrames*i) + "F"
		var fontSize int32 = 10
		rl.DrawText(label, lib.MinInt32((rect.ToInt32().X+rect.ToInt32().Width)-rl.MeasureText(label, fontSize), x+4), rect.ToInt32().Y, fontSize, rl.White)
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

	layout := lib.NewLayout(lib.PublicLayouyt{
		Direction: lib.DIRECTION_COLUMN,
		Gap:       2,
	}, rl.NewVector2(rect.X, rect.Y))

	for _, l := range layers {
		layout.Add(TimelinePropery(l, frameWidth, rect))
	}

	layout.Draw()
}

func TimelinePropery(layer layer.Layer, frameWidth float32, containerRect rl.Rectangle) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		rect := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, containerRect.Width, 32)
		keyframes := layer.GetElement().Position.X.Keyframes

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
