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
	DrawTimeline(ui *lib.UIStruct, comp components.Components) lib.Component
	State() components.InteractableState
	DrawHighlight(lib.UIStruct, components.Components)
	Rect(int) rl.Rectangle
}

type Element struct {
	Id       string
	Name     string
	Position AnimatedVector2

	interactable components.Interactable
}

func NewElement(id string, rect rl.Vector2, name string) Element {
	return Element{
		Id:       id,
		Position: NewAnimatedVector2(id, rect.X, rect.Y),
		Name:     name,
	}
}

func (r *Element) State() components.InteractableState {
	return r.interactable.State()
}
