package layout

import (
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type timelineLayout struct{}

func (timelineLayout) NewTilelineRowLayout(positionRect rl.Rectangle) lib.Layout {
	return *lib.NewLayout(lib.PublicLayouyt{
		Gap:       12,
		Direction: lib.DIRECTION_ROW,
	}, rl.NewVector2(positionRect.X, positionRect.Y))
}

func (timelineLayout) Root(rect rl.Rectangle) *lib.MixLayout {
	return lib.NewMixLayout(lib.PublicMixLayouyt{
		Gap:       12,
		Direction: lib.DIRECTION_COLUMN,
		InitialRect: lib.MixLayouytRect{
			Position: rl.NewVector2(rect.X, rect.Y),
			Width: lib.ContrainedSize{
				Value: rect.Width,
				Contrains: []lib.ChildSize{
					{
						SizeType: lib.SIZE_WEIGHT,
						Value:    1,
					},
				},
			},
		},
	})
}

func (timelineLayout) Row(rect rl.Rectangle) *lib.MixLayout {
	return lib.NewMixLayout(lib.PublicMixLayouyt{
		Direction: lib.DIRECTION_ROW,
		Gap:       12,
		InitialRect: lib.MixLayouytRect{
			Position: rl.NewVector2(rect.X, rect.Y),
			Width: lib.ContrainedSize{
				Value: rect.Width,
				Contrains: []lib.ChildSize{
					{
						SizeType: lib.SIZE_ABSOLUTE,
						Value:    280,
					},
					{
						SizeType: lib.SIZE_WEIGHT,
						Value:    1,
					},
				},
			},
		},
	})
}
func (timelineLayout) Panel(rect rl.Rectangle) *lib.MixLayout {
	padding := lib.Padding{}
	// padding.Start(32)
	return lib.NewMixLayout(lib.PublicMixLayouyt{
		Padding:   padding,
		Direction: lib.DIRECTION_ROW,
		Gap:       16,
		InitialRect: lib.MixLayouytRect{
			Position: rl.NewVector2(rect.X, rect.Y),
			Width: lib.ContrainedSize{
				Value: rect.Width,
				Contrains: []lib.ChildSize{
					{
						SizeType: lib.SIZE_WEIGHT,
						Value:    1,
					},
					{
						SizeType: lib.SIZE_ABSOLUTE,
						Value:    100,
					},
				},
			},
		},
	})
}

var Timeline = timelineLayout{}
