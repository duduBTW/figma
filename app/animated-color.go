package app

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedColor struct {
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

func NewAnimatedColor(red, blue, green, alpha float32) AnimatedColor {
	return AnimatedColor{
		Red:   NewAnimatedProp(red, red_KEY),
		Blue:  NewAnimatedProp(blue, blue_KEY),
		Green: NewAnimatedProp(green, green_KEY),
		Alpha: NewAnimatedProp(alpha, alpha_KEY),
	}
}

func (a *AnimatedColor) Get(selectedFrame int) rl.Color {
	red := uint8(a.Red.KeyFramePosition(selectedFrame))
	green := uint8(a.Green.KeyFramePosition(selectedFrame))
	blue := uint8(a.Blue.KeyFramePosition(selectedFrame))
	alpha := uint8(a.Alpha.KeyFramePosition(selectedFrame))
	return rl.NewColor(red, green, blue, alpha)
}
