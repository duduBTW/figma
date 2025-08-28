package home

import (
	"github.com/dudubtw/figma/app"
	ds "github.com/dudubtw/figma/design-system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Page() {
	app.NewLayout().
		Position(rl.NewVector2(0, 0)).
		Padding(app.NewPadding().All(ds.SPACING_4)).
		Gap(ds.SPACING_4).
		Column().
		Width(float32(rl.GetScreenWidth())).
		Height(float32(rl.GetScreenHeight())).
		Add(UpperPart()).
		Add(FileGrid()).
		Draw()
}
