package main

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func UpperPart() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.
			NewLayout().
			PositionRect(rect).
			Row().
			Gap(PANEL_GAP).
			Width(rect.Width,
				app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1},
				app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: SIDE_PANEL_WIDTH}).
			Height(rect.Height).
			Add(LeftPart()).
			Add(RightPart())
		return layout.Draw, 0, 0
	}
}
