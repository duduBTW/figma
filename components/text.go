package components

import (
	"github.com/dudubtw/figma/app"
	ds "github.com/dudubtw/figma/design-system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Typography(text string, fontSize ds.DsFontSize, weight ds.DsFontWeight, color ds.DsColor) app.Component {
	font := app.Apk.GetFont(fontSize, weight)
	var spacing float32 = 0

	return func(rect rl.Rectangle) (func(), float32, float32) {
		size := rl.MeasureTextEx(font, text, float32(fontSize), spacing)
		return func() {
			rl.DrawTextEx(font, text, rl.NewVector2(rect.X, rect.Y), float32(fontSize), spacing, color)
		}, size.X, size.Y
	}
}
