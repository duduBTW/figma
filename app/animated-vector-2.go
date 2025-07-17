package app

const (
	x_KEY = "x"
	y_KEY = "y"
	EMPTY = "_$_$_$_$_empty__._$$__"
)

type AnimatedVector2 struct {
	Id string
	X  AnimatedProp
	Y  AnimatedProp
}

func NewAnimatedVector2(id string, x float32, y float32) *AnimatedVector2 {
	return &AnimatedVector2{
		Id: id,
		X:  NewAnimatedProp(x, x_KEY),
		Y:  NewAnimatedProp(y, y_KEY),
	}
}
