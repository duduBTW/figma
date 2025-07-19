package components

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Icon(name app.IconName) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		return func() {
			rl.DrawTexture(app.Apk.Icon(name), rect.ToInt32().X, rect.ToInt32().Y, rl.Fade(rl.White, 0.8))
		}, app.ICON_WIDTH, app.ICON_HEIGHT
	}
}
