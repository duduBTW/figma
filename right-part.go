package main

import (
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func RightPart(rect rl.Rectangle) {
	layout := lib.NewConstrainedLayout(lib.ContrainedLayout{
		Direction: lib.DIRECTION_COLUMN,
		Contrains: rect,
		Gap:       PANEL_GAP,
		ChildrenSize: []lib.ChildSize{
			{
				SizeType: lib.SIZE_ABSOLUTE,
				Value:    TOOL_DOCK_HEIGHT,
			},
			{
				SizeType: lib.SIZE_WEIGHT,
				Value:    1,
			},
		},
	})
	layout.Add(ToolDock)
	layout.Add(PropertiesPanel)
	layout.Draw()
}

func ToolDock(rect rl.Rectangle) {
	DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))

	padding := lib.Padding{}
	padding.Axis(12, 8)
	layout := lib.NewLayout(lib.PublicLayouyt{
		Direction: lib.DIRECTION_ROW,
		Padding:   padding,
		Gap:       8,
	}, rl.NewVector2(rect.X, rect.Y))
	layout.Add(Test)
	layout.Add(Square)
	layout.Add(Square)
	layout.Add(Square)
	layout.Draw()
}

var Square = lib.NewComponent(func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
	rect := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, 32, 32)
	return func() {
		DrawRectangleRoundedPixels(rect, 4, rl.NewColor(12, 159, 233, 255))
	}, rect
})

func Test(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
	button := c.Button("test", avaliablePosition, []lib.Component{Content})
	return button.Draw, button.Rect
}

func Content(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
	textContet := "Select"
	fontSize := 16
	rect := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, float32(rl.MeasureText(textContet, int32(fontSize))), float32(fontSize))
	return func() {
		rl.DrawText(textContet, int32(avaliablePosition.X), int32(avaliablePosition.Y), int32(fontSize), rl.White)
	}, rect
}

func PropertiesPanel(rect rl.Rectangle) {
	DrawRectangleRoundedPixels(rect, PANEL_ROUNDNESS, rl.NewColor(34, 34, 34, 255))

	if selectedLayer == nil {
		return
	}

	rect.X += 12
	rect.Y += 12
	rect.Width -= 24
	rect.Height -= 24

	selectedLayer.DrawControls(&ui, rect, c)
}
