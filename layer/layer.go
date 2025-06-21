package layer

import (
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedProp struct {
	Base      float32
	Keyframes [][2]float32
}

type AnimatedVector2 struct {
	X AnimatedProp
	Y AnimatedProp
}

func NewAnimatedVector2(x float32, y float32) AnimatedVector2 {
	return AnimatedVector2{
		X: AnimatedProp{Base: x, Keyframes: [][2]float32{}},
		Y: AnimatedProp{Base: y, Keyframes: [][2]float32{}},
	}
}

type Layer interface {
	GetName() string
	GetElement() *Element
	DrawComponent(ui *lib.UIStruct, mousePoint rl.Vector2) bool
	DrawControls(ui *lib.UIStruct, rect rl.Rectangle, comp components.Components)
	State() components.InteractableState
	DrawHighlight()
}

type Element struct {
	Id       string
	Name     string
	Position AnimatedVector2
}

type Circle struct {
	Element
}

func (c *Circle) GetElement() *Element {
	return &c.Element
}
func (c *Circle) DrawComponent() {
}

type Text struct {
	Element
}

func (t *Text) GetElement() *Element {
	return &t.Element
}

func (c *Text) DrawComponent() {
}
