package home

import (
	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/components"
	ds "github.com/dudubtw/figma/design-system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func UpperPart() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.
			NewLayout().
			Direction(app.DIRECTION_ROW).
			PositionRect(rect).
			Gap(ds.SPACING_2).
			Add(components.OpenTabs()).
			Add(NewFileButton())

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}

func NewFileButton() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := components.Button(
			"new-file",
			components.BUTTON_VARIANT_PRIMARY,
			rl.NewVector2(rect.X, rect.Y),
			[]app.Component{components.Icon(app.ICON_PLUS), components.ButtonText("New file")},
		)

		if button.Clicked {
			app.Apk.Navigate(app.PAGE_NEW_WORKPLACE)
		}

		return button.Draw, button.Rect.Width, button.Rect.Height
	}
}
