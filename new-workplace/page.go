package newWorkplace

import (
	"github.com/dudubtw/figma/app"
	ds "github.com/dudubtw/figma/design-system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const PAGE_WIDTH = 400

func Page() {
	app.NewLayout().
		Position(rl.NewVector2(0, 0)).
		Padding(app.NewPadding().All(ds.T2_BODY_PADDING)).
		Row().
		Width(float32(rl.GetScreenWidth()), app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: -1}).
		Height(float32(rl.GetScreenHeight()), app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: -1}).
		VerticalAlignment(app.ALIGNMENT_CENTER).
		HorizontalAlignment(app.ALIGNMENT_CENTER).
		Add(Form()).
		Draw()
}
