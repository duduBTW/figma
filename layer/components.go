package layer

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPanelLayout(rect rl.Rectangle) *app.Layout {
	return app.
		NewLayout().
		PositionRect(rect).
		Width(rect.Width).
		Column().
		Gap(8)
}
