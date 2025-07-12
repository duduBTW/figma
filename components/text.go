package components

import (
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Text(text string, fontSize int32) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		width := rl.MeasureText(text, fontSize)
		return func() {
			rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y, fontSize, rl.White)
		}, float32(width), float32(fontSize)
	}
}
