package components

import (
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonInstance struct {
	Draw    func()
	Clicked bool
	Rect    rl.Rectangle
}

func (c *Components) Button(id string, position rl.Vector2, children []lib.Component) ButtonInstance {
	buttonInstance := ButtonInstance{}
	interactable := NewInteractable(id, c.ui)

	padding := lib.Padding{}
	padding.All(8)
	layout := lib.NewLayout(lib.PublicLayouyt{
		Padding:   padding,
		Direction: lib.DIRECTION_ROW,
		Gap:       8,
	}, position)

	for _, component := range children {
		layout.Add(component)
	}

	containerRect := rl.NewRectangle(position.X, position.Y, layout.Size.Width, layout.Size.Height)
	clicked := interactable.Event(rl.GetMousePosition(), containerRect)
	var containerBackgroundColor = rl.Purple
	switch interactable.State() {
	case STATE_HOT:
		containerBackgroundColor = rl.DarkPurple
	case STATE_ACTIVE:
		containerBackgroundColor = rl.Pink
	}

	buttonInstance.Draw = func() {
		DrawRectangleRoundedPixels(containerRect, 4, containerBackgroundColor)
		layout.Draw()
	}

	buttonInstance.Clicked = clicked
	buttonInstance.Rect = containerRect

	return buttonInstance
}
