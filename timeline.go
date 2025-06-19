package main

import (
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

	layout.Render(LayersPanel)
	layout.Render(TimelineFrames)
}

func LayersPanel(rect rl.Rectangle) {
	layout := lib.NewLayout(lib.PublicLayouyt{
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

func TimelineFrames(rect rl.Rectangle) {
	DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(41, 41, 41, 255))
}
