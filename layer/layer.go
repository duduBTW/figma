package layer

import (
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	Position rl.Vector2
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
