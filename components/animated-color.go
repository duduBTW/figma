package components

import (
	"fmt"

	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ncruces/zenity"
)

type animatedColorComponent struct {
	prop   *app.AnimatedColor
	layer  app.Layer
	prefix string
}

func NewAnimatedColor(animatedProp *app.AnimatedColor, layer app.Layer, prefix string) *animatedColorComponent {
	return &animatedColorComponent{
		prop:   animatedProp,
		layer:  layer,
		prefix: prefix,
	}
}

func (component *animatedColorComponent) CanDrawTimeline() bool {
	return len(component.prop.Red.SortedKeyframesTimeline()) > 0
}
func (component *animatedColorComponent) Controls() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := NewSidebarProperyLabel(rect).
			Add(SidebarProperyLabel("Color")).
			Add(component.Input())
		return row.Draw, 0, row.Size.Height
	}
}

func (component *animatedColorComponent) Input() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout :=
			app.NewLayout().
				PositionRect(rect).
				Row().
				Width(rect.Width,
					app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1},
					app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: 24},
					app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: 24}).
				Add(component.colorControlsInputs()).
				Add(component.keyFrameButton()).
				Add(component.colorPicker())
		return layout.Draw, 0, layout.Size.Height
	}
}

func (component *animatedColorComponent) colorPicker() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layer := component.layer
		prefix := component.prefix
		button := Button(layer.GetElement().Id+"color-picker"+prefix, BUTTON_VARIANT_PRIMARY, rl.NewVector2(rect.X, rect.Y), []app.Component{})
		if button.Clicked {
			color, err := zenity.SelectColor(
				zenity.Title(layer.GetName()),
			)

			if err == nil {
				r, g, b, a := color.RGBA()
				component.prop.Set(float32(r), float32(g), float32(b), float32(a))
				fmt.Println(component.prop.Red.SortedKeyframes)
			}
		}

		return button.Draw, 0, 0
	}
}

func (component *animatedColorComponent) keyFrameButton() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		animatedProp := component.prop
		layer := component.layer
		prefix := component.prefix
		button := Button(layer.GetElement().Id+"color"+prefix, BUTTON_VARIANT_PRIMARY, rl.NewVector2(rect.X, rect.Y), []app.Component{})
		if button.Clicked {
			animatedProp.InsertKeyframe()
		}
		return button.Draw, 0, 0
	}
}

func (c *animatedColorComponent) colorControlsInputs() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		tempValue := c.prop.InputValue

		// Current values from keyframes
		currentHex := app.ColorToHex(c.prop.Get(app.Apk.State.SelectedFrame))
		if tempValue == app.EMPTY {
			tempValue = currentHex
		}

		input := Input(InputProps{
			X:          rect.X,
			Y:          rect.Y,
			Id:         c.layer.GetElement().Id + "_colorhex_" + c.prefix,
			Width:      rect.Width,
			Value:      tempValue,
			MousePoint: rl.GetMousePosition(),
		})

		if input.IsFocusing || input.State == app.STATE_ACTIVE {
			c.prop.InputValue = input.Value
		}

		if input.IsBluring || input.HasSubmitted {
			input.Blur()
			if c.prop.InputValue == "" || c.prop.InputValue == currentHex {
				c.prop.InputValue = app.EMPTY
			} else {
				r, g, b, a, err := app.HexToColor(c.prop.InputValue)
				if err != nil {
					c.prop.InputValue = currentHex
				}

				c.prop.Set(r, g, b, a)
				c.prop.InputValue = app.EMPTY
			}
		}

		return input.Draw, 0, input.Rect.Height
	}
}
