package lib

import (
	"errors"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layout struct {
	padding   Padding
	direction Direction
	position  rl.Vector2
	width     ContrainedSize
	height    ContrainedSize
	gap       float32
	index     int
	drawStack []func()

	Size rl.Rectangle
}

type ContrainedSize struct {
	Value     float32
	Contrains []ChildSize
	computed  []float32

	// where the user is right now
	currentIndex int
}

func NewContrainedSize(value float32, constrains ...ChildSize) ContrainedSize {
	constrinedSize := ContrainedSize{}
	constrinedSize.Value = value
	constrinedSize.Contrains = constrains

	if len(constrains) == 0 {
		constrinedSize.Contrains = []ChildSize{{
			SizeType: SIZE_WEIGHT,
			Value:    1,
		}}
	}

	return constrinedSize
}
func (size *ContrainedSize) Exists() bool {
	return size.Value != 0 && len(size.Contrains) > 0
}
func (size *ContrainedSize) Compute(gap float32, padding float32) {
	if !size.Exists() {
		return
	}

	size.Value -= padding
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

	value := size.computed[size.currentIndex]
	return value
}
func (size *ContrainedSize) NextValue(gap float32) float32 {
	value := size.CurrentValue(gap)

	if size.currentIndex < len(size.computed)-1 {
		size.currentIndex++
	}
	return value
}

// return draw fun, width, height
// TODO
// FIND BETTER WAY OF MIXING LAYOUTS WITH EACHOTHER
type Component func(rect rl.Rectangle) (func(), float32, float32)

func NewLayout() *Layout {
	return &Layout{}
}

func (layout *Layout) Padding(padding *Padding) *Layout {
	layout.padding = *padding
	layout.Size.X = layout.position.X + padding.start
	layout.Size.Y = layout.position.Y + padding.top
	return layout
}
func (layout *Layout) Gap(gap float32) *Layout {
	layout.gap = gap
	return layout
}
func (layout *Layout) Direction(direction Direction) *Layout {
	layout.direction = direction
	return layout
}
func (layout *Layout) Row() *Layout {
	layout.direction = DIRECTION_ROW
	return layout
}
func (layout *Layout) Column() *Layout {
	layout.direction = DIRECTION_COLUMN
	return layout
}
func (layout *Layout) Position(position rl.Vector2) *Layout {
	layout.position = position
	return layout
}
func (layout *Layout) PositionRect(rect rl.Rectangle) *Layout {
	layout.position = rl.NewVector2(rect.X, rect.Y)
	layout.Size.X = rect.X
	layout.Size.Y = rect.Y
	return layout
}

func (layout *Layout) Width(value float32, constrains ...ChildSize) *Layout {
	layout.width = NewContrainedSize(value, constrains...)
	layout.width.Compute(layout.gap, layout.padding.start+layout.padding.end)
	layout.Size.Width = layout.width.Value
	return layout
}
func (layout *Layout) Height(value float32, constrains ...ChildSize) *Layout {
	layout.height = NewContrainedSize(value, constrains...)
	layout.height.Compute(layout.gap, layout.padding.top+layout.padding.bottom)
	// TODO
	// IMPROVE CURRENT RECT AND INITIAL RECT LOGIC
	layout.Size.Height = layout.height.Value
	return layout
}

func (layout *Layout) Add(component Component) *Layout {
	width := layout.width.CurrentValue(layout.gap)
	height := layout.height.CurrentValue(layout.gap)
	draw, width, height := component(rl.NewRectangle(layout.Size.X, layout.Size.Y, width, height))
	layout.drawStack = append(layout.drawStack, draw)
	layout.next(width, height)
	return layout
}

func (layout *Layout) Draw() {
	for _, draw := range layout.drawStack {
		draw()
	}
}

func (layout *Layout) next(width, height float32) {
	// TODO - FIX THIS, STOP USING ZERO
	// Width
	if width == 0 {
		width = layout.width.NextValue(layout.gap)
	}

	// Height
	if height == 0 {
		height = layout.height.NextValue(layout.gap)
	}

	switch layout.direction {
	case DIRECTION_ROW:
		isFirst := layout.position.X == layout.Size.X
		layout.Size.X += width + layout.gap

		currentHeight := layout.Size.Height
		layout.Size.Height = Max(currentHeight, height+layout.padding.top+layout.padding.bottom)
		if layout.width.Exists() {
			return
		}

		layout.Size.Width += width
		if !isFirst {
			layout.Size.Width += layout.gap
		}
	case DIRECTION_COLUMN:
		isFirst := layout.position.Y == layout.Size.Y
		layout.Size.Y += height + layout.gap

		currentWidth := layout.Size.Width
		layout.Size.Width = Max(currentWidth, width+layout.padding.start+layout.padding.end)
		if layout.height.Exists() {
			return
		}

		layout.Size.Height += height
		if !isFirst {
			layout.Size.Height += layout.gap
		}
	}
}

func ComputeChildren(childrenSize []ChildSize, value float32, gap float32) ([]float32, error) {
	var computedSizes = make([]float32, len(childrenSize))
	if len(childrenSize) == 0 {
		return computedSizes, errors.New("no children to compute")
	}

	type Index = int
	var remainingSize float32 = value - gap*float32(len(childrenSize)-1)

	var weightSum float32 = 0
	var weightSizes = make(map[Index]float32)
	for index, ChildSize := range childrenSize {
		value := ChildSize.Value
		if ChildSize.SizeType == SIZE_WEIGHT {
			weightSizes[index] = value
			weightSum += value
			continue
		}

		computedSizes[index] = value
		remainingSize -= value
	}

	if weightSum > 1 {
		return computedSizes, errors.New("weight sum not equal to 1")
	}

	for index, weight := range weightSizes {
		computedSizes[index] = float32(remainingSize) * weight
	}

	return computedSizes, nil
}
