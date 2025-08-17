package app

import rl "github.com/gen2brain/raylib-go/raylib"

func CubicBezierPoint(p1, c2, c3, p4 rl.Vector2, t float32) rl.Vector2 {
	u := 1 - t
	tt := t * t
	uu := u * u
	uuu := uu * u
	ttt := tt * t

	x := uuu*p1.X + 3*uu*t*c2.X + 3*u*tt*c3.X + ttt*p4.X
	y := uuu*p1.Y + 3*uu*t*c2.Y + 3*u*tt*c3.Y + ttt*p4.Y

	return rl.Vector2{X: x, Y: y}
}
