package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Padding struct {
	top    float32
	bottom float32
	start  float32
	end    float32
}

func (p *Padding) Axis(horizontal, vertical float32) {
	p.top = vertical
	p.bottom = vertical
	p.start = horizontal
	p.end = horizontal
}
func (p *Padding) All(padding float32) {
	p.top = padding
	p.bottom = padding
	p.start = padding
	p.end = padding
}
func (p *Padding) Top(top float32) {
	p.top = top
}
func (p *Padding) Bottom(bottom float32) {
	p.bottom = bottom
}
func (p *Padding) Start(start float32) {
	p.start = start
}
func (p *Padding) End(end float32) {
	p.end = end
}

type Alignment string

const (
	ALIGNMENT_START  Alignment = "start"
	ALIGNMENT_CENTER Alignment = "center"
	ALIGNMENT_END    Alignment = "end"
)

type Direction string

const (
	DIRECTION_ROW    Direction = "row"
	DIRECTION_COLUMN Direction = "column"
)

type SizeType string

const (
	SIZE_ABSOLUTE SizeType = "absolute"
	SIZE_WEIGHT   SizeType = "weight"
)

type Size struct {
	Width  float32
	Height float32
}

type PublicLayouyt struct {
	Padding
	Direction Direction
	Gap       int
}

type Layout struct {
	PublicLayouyt
	index int

	VerticalAlignment   Alignment
	HorizontalAlighment Alignment

	Position rl.Vector2
	Size     Size

	drawStack []func()
}

type Component func(avaliablePosition rl.Vector2) (func(), rl.Rectangle)

// FIX-ME DIRECTION NON PASSED
func NewLayout(props PublicLayouyt, contrains rl.Vector2) *Layout {
	return &Layout{
		PublicLayouyt: props,
		Position: rl.Vector2{
			X: contrains.X + props.Padding.start,
			Y: contrains.Y + props.Padding.top,
		},
		Size: Size{
			Width:  props.Padding.start + props.Padding.end,
			Height: props.Padding.top + props.Padding.bottom,
		},
	}
}

func (layout *Layout) Draw() {
	for _, draw := range layout.drawStack {
		draw()
	}
}

func (layout *Layout) Add(component Component) {
	draw, position := component(layout.Position)
	layout.drawStack = append(layout.drawStack, draw)
	layout.next(position)
}

func (layout *Layout) next(component rl.Rectangle) {
	switch layout.Direction {
	case DIRECTION_ROW:
		// size
		layout.Size.Width = layout.Size.Width + component.Width
		// gap
		if layout.index != 0 {
			layout.Size.Width += float32(layout.Gap)
		}

		layout.Size.Height = Max(layout.Size.Height, component.Height+(layout.Padding.top+layout.Padding.bottom))

		// position
		layout.Position.X = layout.Position.X + component.Width + float32(layout.Gap)
	case DIRECTION_COLUMN:
		// size
		layout.Size.Height = layout.Size.Height + component.Height
		// gap
		if layout.index != 0 {
			layout.Size.Height += float32(layout.Gap)
		}

		layout.Size.Width = Max(layout.Size.Width, component.Width+(layout.Padding.start+layout.Padding.end))

		// positioN
		layout.Position.Y = layout.Position.Y + component.Height + float32(layout.Gap)
	}
	layout.index++
}

func NewComponent(component Component) Component {
	return component
}
