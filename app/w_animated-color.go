package app

import (
	"fmt"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedColor struct {
	Name       string
	InputValue string

	Red   AnimatedProp
	Blue  AnimatedProp
	Green AnimatedProp
	Alpha AnimatedProp
}

const (
	red_KEY   = "red"
	blue_KEY  = "blue"
	green_KEY = "green"
	alpha_KEY = "alpha"
)

func NewAnimatedColor(red, blue, green, alpha float32, name string) AnimatedColor {
	return AnimatedColor{
		Red:        NewAnimatedProp(red, red_KEY),
		Blue:       NewAnimatedProp(blue, blue_KEY),
		Green:      NewAnimatedProp(green, green_KEY),
		Alpha:      NewAnimatedProp(alpha, alpha_KEY),
		InputValue: EMPTY,
		Name:       name,
	}
}

func (a *AnimatedColor) Get(selectedFrame int) rl.Color {
	red := uint8(a.Red.KeyFramePosition(selectedFrame))
	green := uint8(a.Green.KeyFramePosition(selectedFrame))
	blue := uint8(a.Blue.KeyFramePosition(selectedFrame))
	alpha := uint8(a.Alpha.KeyFramePosition(selectedFrame))
	return rl.NewColor(red, green, blue, alpha)
}

func (a *AnimatedColor) Set(red, green, blue, alpha float32) {
	selectedFrame := float32(Apk.Workplace.SelectedFrame)
	a.Red.InsertKeyframe(selectedFrame, red)
	a.Green.InsertKeyframe(selectedFrame, green)
	a.Blue.InsertKeyframe(selectedFrame, blue)
	a.Alpha.InsertKeyframe(selectedFrame, alpha)
}

func (a *AnimatedColor) InsertKeyframe() {
	frame := float32(Apk.Workplace.SelectedFrame)
	a.Red.InsertKeyframe(frame, a.Red.Base)
	a.Green.InsertKeyframe(frame, a.Green.Base)
	a.Blue.InsertKeyframe(frame, a.Blue.Base)
	a.Alpha.InsertKeyframe(frame, a.Alpha.Base)

}

// -----------------
// Utils
// -----------------

func ColorToHex(color rl.Color) string {
	return fmt.Sprintf("#%02X%02X%02X%02X", int(color.R), int(color.G), int(color.B), int(color.A))
}

func HexToColor(hex string) (float32, float32, float32, float32, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 6 {
		// assume fully opaque
		r, _ := strconv.ParseInt(hex[0:2], 16, 32)
		g, _ := strconv.ParseInt(hex[2:4], 16, 32)
		b, _ := strconv.ParseInt(hex[4:6], 16, 32)
		return float32(r), float32(g), float32(b), 255, nil
	} else if len(hex) == 8 {
		r, _ := strconv.ParseInt(hex[0:2], 16, 32)
		g, _ := strconv.ParseInt(hex[2:4], 16, 32)
		b, _ := strconv.ParseInt(hex[4:6], 16, 32)
		a, _ := strconv.ParseInt(hex[6:8], 16, 32)
		return float32(r), float32(g), float32(b), float32(a), nil
	}
	return 0, 0, 0, 0, fmt.Errorf("invalid hex")
}
