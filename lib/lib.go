package lib

type ChildSize struct {
	SizeType
	Value float32
}

type Padding struct {
	top    float32
	bottom float32
	start  float32
	end    float32
}

func NewPadding() *Padding {
	return &Padding{}
}

func (p *Padding) Axis(horizontal, vertical float32) *Padding {
	p.top = vertical
	p.bottom = vertical
	p.start = horizontal
	p.end = horizontal
	return p
}
func (p *Padding) All(padding float32) *Padding {
	p.top = padding
	p.bottom = padding
	p.start = padding
	p.end = padding
	return p
}
func (p *Padding) Top(top float32) *Padding {
	p.top = top
	return p
}
func (p *Padding) Bottom(bottom float32) *Padding {
	p.bottom = bottom
	return p
}
func (p *Padding) Start(start float32) *Padding {
	p.start = start
	return p
}
func (p *Padding) End(end float32) *Padding {
	p.end = end
	return p
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
