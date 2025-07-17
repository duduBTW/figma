package components

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BoxInstance struct {
	Draw func()
	Rect rl.Rectangle
}

type BoxProps struct {
	Rect         rl.Rectangle
	Padding      app.Padding
	BorderRadius float32
	Color        rl.Color
	Children     []app.Component
	Direction    app.Direction
	Gap          float32
}

func Box(props BoxProps) BoxInstance {
	boxInstance := BoxInstance{}
	direction := props.Direction
	gap := props.Gap
	children := props.Children
	rect := props.Rect
	padding := props.Padding
	color := props.Color
	borderRadius := props.BorderRadius

	// Ensure radius is not negative
	radiusPixels := maxF(0, borderRadius)

	// padding
	layout := app.
		NewLayout().
		PositionRect(rect).
		Padding(&padding).
		Gap(gap).
		Direction(direction)

	for _, component := range children {
		layout.Add(component)
	}

	drawRect := rl.NewRectangle(rect.X, rect.Y, rect.Width, rect.Height)
	if rect.Width == 0 {
		drawRect.Width = layout.Size.Width
	}
	if rect.Height == 0 {
		drawRect.Height = layout.Size.Height
	}

	boxInstance.Draw = func() {
		DrawRectangleRoundedPixels(drawRect, radiusPixels, color)
		layout.Draw()
	}

	boxInstance.Rect = drawRect
	return boxInstance
}
