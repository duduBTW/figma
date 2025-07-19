package app

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layer interface {
	GetName() string
	GetElement() *Element
	DrawComponent(mousePoint rl.Vector2, canvasRect rl.Rectangle) bool
	DrawControls(rect rl.Rectangle)
	DrawTimeline() Component
	State() InteractableState
	DrawHighlight()
	Rect(int) rl.Rectangle
}

type Element struct {
	Id       string
	Name     string
	Position *AnimatedVector2

	Interactable Interactable
}

func NewElement(id string, rect rl.Vector2, name string) Element {
	return Element{
		Id:       id,
		Position: NewAnimatedVector2(id, rect.X, rect.Y),
		Name:     name,
	}
}

func (r *Element) State() InteractableState {
	return r.Interactable.State()
}
