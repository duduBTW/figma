package components

import (
	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/fmath"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonInstance struct {
	Draw    func()
	Clicked bool
	Rect    rl.Rectangle
}

func Button(id string, position rl.Vector2, children []app.Component) ButtonInstance {
	buttonInstance := ButtonInstance{}
	interactable := app.NewInteractable(id)

	layout := app.
		NewLayout().
		Position(position).
		Row().
		Padding(app.NewPadding().All(8)).
		Gap(8)

	for _, component := range children {
		layout.Add(component)
	}

	containerRect := rl.NewRectangle(position.X, position.Y, fmath.Max(layout.Size.Width, 24), fmath.Max(layout.Size.Height, 24))
	clicked := interactable.Event(rl.GetMousePosition(), containerRect)
	var containerBackgroundColor = rl.Purple
	switch interactable.State() {
	case app.STATE_HOT:
		containerBackgroundColor = rl.DarkPurple
	case app.STATE_ACTIVE:
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
