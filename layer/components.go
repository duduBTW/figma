package layer

import (
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewControlsLayout(x float32, y float32, width float32) (lib.ContrainedLayout, rl.Rectangle) {
	contrains := rl.NewRectangle(x, y, width, 24)
	return lib.NewConstrainedLayout(lib.ContrainedLayout{
		Direction: lib.DIRECTION_ROW,
		Gap:       32,
		Contrains: contrains,
		ChildrenSize: []lib.ChildSize{
			{
				SizeType: lib.SIZE_ABSOLUTE,
				Value:    60,
			},
			{
				SizeType: lib.SIZE_WEIGHT,
				Value:    1,
			},
		},
	}), contrains
}

func NewPanelLayout(rect rl.Rectangle) *lib.Layout {
	return lib.NewLayout(lib.PublicLayouyt{
		Gap:       8,
		Direction: lib.DIRECTION_COLUMN,
	}, rl.NewVector2(rect.X, rect.Y))
}

func InputsLayout(amount int, rect rl.Rectangle) lib.ContrainedLayout {
	childrenSize := make([]lib.ChildSize, 0, amount)
	childSize := 1 / float32(amount)
	for i := 0; i < amount; i++ {
		childrenSize = append(childrenSize, lib.ChildSize{
			SizeType: lib.SIZE_WEIGHT,
			Value:    childSize,
		})
	}

	return lib.NewConstrainedLayout(lib.ContrainedLayout{
		Direction:    lib.DIRECTION_ROW,
		Gap:          8,
		Contrains:    rl.NewRectangle(rect.X, rect.Y, rect.Width, 24),
		ChildrenSize: childrenSize,
	})
}

func Label(text string) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		fontSize := 14
		rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y+4, int32(fontSize), rl.White)
	}
}
