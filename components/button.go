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

	layout := lib.
		NewLayout().
		Position(position).
		Row().
		Padding(lib.NewPadding().All(8)).
		Gap(8)

	for _, component := range children {
		layout.Add(component)
	}

	containerRect := rl.NewRectangle(position.X, position.Y, lib.Max(layout.Size.Width, 24), lib.Max(layout.Size.Height, 24))
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
