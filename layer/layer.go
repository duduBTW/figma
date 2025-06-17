package layer

import (
	"fmt"
	"strconv"

	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layer interface {
	GetElement() *Element
	DrawComponent(ui *lib.UIStruct, mousePoint rl.Vector2) bool
	DrawControls(ui *lib.UIStruct, rect rl.Rectangle)
	State() components.InteractableState
	DrawHighlight()
}

type Element struct {
	Id       string
	Position rl.Vector2
}

type Rectangle struct {
	Element
	Width  int32
	Height int32
	Color  rl.Color

	interactable components.Interactable
}

func (r *Rectangle) GetElement() *Element {
	return &r.Element
}
func (r *Rectangle) State() components.InteractableState {
	return r.interactable.State()
}
func (r *Rectangle) DrawHighlight() {
	rl.DrawRectangleLinesEx(r.Rect(), 2, rl.Blue)
}
func (r *Rectangle) Rect() rl.Rectangle {
	return rl.NewRectangle(r.Position.X, r.Position.Y, float32(r.Width), float32(r.Height))
}
func (r *Rectangle) DrawComponent(ui *lib.UIStruct, mousePoint rl.Vector2) bool {
	interactable := components.NewInteractable(r.Id, ui)
	clicked := interactable.Event(mousePoint, rl.NewRectangle(r.Position.X, r.Position.Y, float32(r.Height), float32(r.Height)))
	rl.DrawRectangleRec(r.Rect(), r.Color)
	r.interactable = interactable

	return clicked
}

func (r *Rectangle) DrawControls(ui *lib.UIStruct, rect rl.Rectangle) {
	var newValue = components.Input(components.InputProps{
		X:             rect.X,
		Y:             rect.Y,
		Id:            r.Id + "pox-y",
		Width:         rect.Width,
		Value:         fmt.Sprint(r.Position.Y),
		MousePoint:    rl.GetMousePosition(),
		Ui:            ui,
		LeftIndicator: 'Y',
	})
	var newIntValue, _ = strconv.ParseFloat(newValue, 32)
	r.Position.Y = float32(newIntValue)
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
