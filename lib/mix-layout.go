package lib

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MixLayouytRect struct {
	// Rect
	Position rl.Vector2
	Width    ContrainedSize
	Height   ContrainedSize
}

type PublicMixLayouyt struct {
	Padding
	Direction   Direction
	InitialRect MixLayouytRect
	Gap         float32
}

type ContrainedSize struct {
	Value     float32
	Contrains []ChildSize
	computed  []float32

	// where the user is right now
	CurrentIndex int
}

func (size *ContrainedSize) Exists() bool {
	return size.Value != 0 && len(size.Contrains) > 0
}
func (size *ContrainedSize) Compute(gap float32) {
	if !size.Exists() {
		return
	}

	computedValue, err := ComputeChildren(size.Contrains, size.Value, gap)
	if err != nil {
		fmt.Println(err)
	}
	size.computed = computedValue
}

func (size *ContrainedSize) CurrentValue(gap float32) float32 {
	if !size.Exists() {
		return 0
	}

	value := size.computed[size.CurrentIndex]
	return value
}
func (size *ContrainedSize) NextValue(gap float32) float32 {
	value := size.CurrentValue(gap)

	if size.CurrentIndex < len(size.computed)-1 {
		size.CurrentIndex++
	} else {
		fmt.Println("not", size.CurrentIndex, len(size.computed)-1)
	}
	return value
}

type MixLayout struct {
	PublicMixLayouyt

	CurrentRect rl.Rectangle

	index     int
	drawStack []func()
}

// return draw fun, width, height
type MixComponent func(rect rl.Rectangle) (func(), float32, float32)

func NewMixLayout(props PublicMixLayouyt) *MixLayout {
	layout := MixLayout{
		PublicMixLayouyt: props,
	}
	layout.SetRect()
	layout.InitialRect.Width.Compute(layout.Gap)
	layout.InitialRect.Height.Compute(layout.Gap)
	return &layout
}

func (layout *MixLayout) SetRect() {
	rect := &layout.InitialRect
	x := rect.Position.X + layout.Padding.start
	y := rect.Position.Y + layout.Padding.top
	width := &rect.Width.Value
	height := &rect.Height.Value

	if rect.Width.Exists() {
		*width -= (layout.Padding.start + layout.Padding.end)
	}
	if rect.Height.Exists() {
		*height -= (layout.Padding.top + layout.Padding.bottom)
	}

	layout.CurrentRect = rl.NewRectangle(x, y, *width, *height)
}

func (layout *MixLayout) Add(component MixComponent) {
	width := layout.InitialRect.Width.CurrentValue(layout.Gap)
	height := layout.InitialRect.Height.CurrentValue(layout.Gap)
	draw, width, height := component(rl.NewRectangle(layout.CurrentRect.X, layout.CurrentRect.Y, width, height))
	layout.drawStack = append(layout.drawStack, draw)
	layout.next(width, height)
}

func (layout *MixLayout) Draw() {
	for _, draw := range layout.drawStack {
		draw()
	}
}

func (layout *MixLayout) next(width, height float32) {
	// Width
	if width == 0 {
		width = layout.InitialRect.Width.NextValue(layout.Gap)
	}

	// Height
	if height == 0 {
		height = layout.InitialRect.Height.NextValue(layout.Gap)
	}

	switch layout.Direction {
	case DIRECTION_ROW:
		layout.CurrentRect.X += width + layout.Gap

		currentHeight := layout.CurrentRect.Height
		layout.CurrentRect.Height = Max(currentHeight, height+layout.top+layout.bottom)
		if layout.InitialRect.Width.Exists() {
			return
		}

		layout.CurrentRect.Width += width
	case DIRECTION_COLUMN:
		layout.CurrentRect.Y += height + layout.Gap

		currentWidth := layout.CurrentRect.Width
		layout.CurrentRect.Width = Max(currentWidth, width+layout.start+layout.end)
		if layout.InitialRect.Height.Exists() {
			return
		}

		layout.CurrentRect.Height += height
	}
}
