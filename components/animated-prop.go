package components

import (
	"fmt"
	"strconv"

	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type animatedPropComponent struct {
	prop   *app.AnimatedProp
	layer  app.Layer
	prefix string
}

func NewAnimatedProp(animatedProp *app.AnimatedProp, layer app.Layer, prefix string) *animatedPropComponent {
	return &animatedPropComponent{
		prop:   animatedProp,
		layer:  layer,
		prefix: prefix,
	}
}

func (component *animatedPropComponent) Input() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout :=
			app.NewLayout().
				PositionRect(rect).
				Row().
				Width(rect.Width,
					app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1},
					app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: 24}).
				Add(component.inputEditableContent()).
				Add(component.keyFrameButton())
		return layout.Draw, 0, layout.Size.Height
	}
}

func (component *animatedPropComponent) keyFrameButton() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		animatedProp := component.prop
		layer := component.layer
		prefix := component.prefix
		button := Button(layer.GetName()+animatedProp.Name+prefix, BUTTON_VARIANT_PRIMARY, rl.NewVector2(rect.X, rect.Y), []app.Component{})
		if button.Clicked {
			animatedProp.InsertKeyframe(float32(app.Apk.State.SelectedFrame), animatedProp.Base)
		}
		return button.Draw, 0, 0
	}
}

func (component *animatedPropComponent) inputEditableContent() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		animatedProp := component.prop
		tempValue := animatedProp.InputValue
		layer := component.layer
		prefix := component.prefix
		updateValue := fmt.Sprint(int(animatedProp.KeyFramePosition(app.Apk.State.SelectedFrame)))
		if tempValue == app.EMPTY {
			tempValue = updateValue
		}

		input := Input(InputProps{
			X:          rect.X,
			Y:          rect.Y,
			Id:         layer.GetElement().Id + animatedProp.Name + prefix,
			Width:      rect.Width,
			Value:      tempValue,
			MousePoint: rl.GetMousePosition(),
		})

		if input.IsFocusing || input.State == app.STATE_ACTIVE {
			animatedProp.InputValue = input.Value
		}

		if input.IsBluring || input.HasSubmitted {
			input.Blur()

			if animatedProp.InputValue == "" || tempValue == updateValue {
				animatedProp.InputValue = app.EMPTY
			} else {
				var newIntValue, err = strconv.ParseFloat(animatedProp.InputValue, 32)
				if err != nil {
					animatedProp.InputValue = updateValue
				}

				animatedProp.Set(float32(newIntValue))
				animatedProp.InputValue = app.EMPTY
			}

		}

		return input.Draw, input.Rect.Width, input.Rect.Height
	}
}

func (component *animatedPropComponent) CanDrawTimeline() bool {
	return len(component.prop.SortedKeyframes) > 0
}

func (component *animatedPropComponent) TimelineFrames() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		const height = 24
		var draw = func() {
			y := int32(rect.Y + (height / 2))
			rl.DrawLine(rect.ToInt32().X, y, rect.ToInt32().X+rect.ToInt32().Width, y, rl.Blue)

			keyframes := component.prop.SortedKeyframes
			if len(keyframes) > 0 {
				for _, keyframe := range keyframes {
					x := app.Apk.State.GetXTimelineFrame(rect, keyframe[0])
					keyframeRect := rl.NewRectangle(x, float32(y), 10, 10)
					rl.DrawRectanglePro(
						keyframeRect,        // A 10×20 rectangle
						rl.NewVector2(5, 5), // Center of the rectangle
						45,
						rl.Blue,
					)

					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(keyframeRect.X-5, keyframeRect.Y-5, keyframeRect.Width, keyframeRect.Height)) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
						// clicked
					}
				}
			}
		}

		return draw, 0, height
	}
}
