package layout

import (
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type timelineLayout struct{}

func (timelineLayout) NewTilelineRowLayout(positionRect rl.Rectangle) *lib.Layout {
	return lib.
		NewLayout().
		PositionRect(positionRect).
		Row().
		Gap(12)
}

func (timelineLayout) Root(rect rl.Rectangle) *lib.Layout {
	return lib.
		NewLayout().
		PositionRect(rect).
		Column().
		Gap(12).
		Width(rect.Width)
}

func (timelineLayout) Row(rect rl.Rectangle) *lib.Layout {
	return lib.
		NewLayout().
		Row().
		PositionRect(rect).
		Gap(12).
		Width(rect.Width,
			lib.ChildSize{SizeType: lib.SIZE_ABSOLUTE, Value: 280},
			lib.ChildSize{SizeType: lib.SIZE_WEIGHT, Value: 1})
}
func (timelineLayout) Panel(rect rl.Rectangle) *lib.Layout {
	return lib.
		NewLayout().
		PositionRect(rect).
		Padding(lib.NewPadding().Start(32)).
		Row().
		Gap(16).
		Width(rect.Width,
			lib.ChildSize{SizeType: lib.SIZE_WEIGHT, Value: 1},
			lib.ChildSize{SizeType: lib.SIZE_ABSOLUTE, Value: 100})
}

var Timeline = timelineLayout{}
