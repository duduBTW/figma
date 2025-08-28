package components

import (
	"github.com/dudubtw/figma/app"
	ds "github.com/dudubtw/figma/design-system"
	"github.com/dudubtw/figma/fmath"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonVariant = string

const (
	BUTTON_VARIANT_PRIMARY  = "primary"
	BUTTON_VARIANT_GHOST    = "ghost"
	BUTTON_VARIANT_OUTLINED = "outlined"
)

type variantColor = map[app.InteractableState]rl.Color

var variantColors = map[ButtonVariant]variantColor{
	BUTTON_VARIANT_PRIMARY: {
		app.STATE_INITIAL: ds.T2_COLOR_PRIMARY,
		app.STATE_HOT:     ds.T2_COLOR_PRIMARY_HOT,
		app.STATE_ACTIVE:  ds.T2_COLOR_PRIMARY_ACTIVE,
	},
	BUTTON_VARIANT_GHOST: {
		app.STATE_INITIAL: rl.Fade(rl.White, 0),
		app.STATE_HOT:     rl.Fade(rl.White, 0.1),
		app.STATE_ACTIVE:  rl.Fade(rl.White, 0.2),
	},
	BUTTON_VARIANT_OUTLINED: {
		app.STATE_INITIAL: rl.Fade(rl.White, 0),
		app.STATE_HOT:     rl.Fade(rl.White, 0.1),
		app.STATE_ACTIVE:  rl.Fade(rl.White, 0.2),
	},
}

type ButtonInstance struct {
	Draw    func()
	Clicked bool
	Rect    rl.Rectangle
}

func Button(id string, variant ButtonVariant, position rl.Vector2, children []app.Component) ButtonInstance {
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

	var containerBackgroundColor = variantColors[variant][interactable.State()]
	buttonInstance.Draw = func() {
		if variant == BUTTON_VARIANT_OUTLINED {
			DrawRectangleRoundedLinePixels(containerRect, ds.RADII_SM, 1, ds.T2_BORDER)
		}

		DrawRectangleRoundedPixels(containerRect, ds.RADII_SM, containerBackgroundColor)

		layout.Draw()
	}

	buttonInstance.Clicked = clicked
	buttonInstance.Rect = containerRect

	return buttonInstance
}

func ButtonText(text string) app.Component {
	return Typography(text, ds.FONT_SIZE_LG, ds.FONT_WEIGHT_MEDIUM, ds.T2_COLOR_CONTENT)
}
