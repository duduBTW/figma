package components

import (
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BoxInstance struct {
	Draw func()
	Rect rl.Rectangle
}

type BoxProps struct {
	Rect         rl.Rectangle
	Padding      lib.Padding
	BorderRadius float32
	Color        rl.Color
	Children     []lib.Component
	Direction    lib.Direction
	Gap          float32
}

func (*Components) Box(props BoxProps) BoxInstance {
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
	layout := lib.
		NewLayout().
		PositionRect(rect).
		Direction(direction).
		Padding(&padding).
		Gap(gap)

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
		// Find the smaller dimension
		minDimension := minF(drawRect.Width, drawRect.Height)

		// Calculate roundness based on pixel radius
		// roundness = (radius * 2) / minDimension
		roundness := (radiusPixels * 2) / minDimension

		// Clamp roundness to the valid range [0.0, 1.0]
		// If requested radiusPixels * 2 > minDimension, it means the radius
		// is too large, so we cap at full roundness (1.0).
		roundness = maxF(0.0, minF(roundness, 1.0))

		// Call the original raylib function with the calculated roundness
		rl.DrawRectangleRounded(drawRect, roundness, 0, color)
		layout.Draw()
	}

	boxInstance.Rect = drawRect
	return boxInstance
}
