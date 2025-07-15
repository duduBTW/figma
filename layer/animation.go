package layer

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedProp struct {
	Name            string
	Base            float32
	SortedKeyframes [][2]float32
	KeyframesMap    map[float32]float32

	inputValue string
}

func NewAnimatedProp(defaultValue float32, name string) AnimatedProp {
	return AnimatedProp{Base: defaultValue, SortedKeyframes: [][2]float32{}, KeyframesMap: map[float32]float32{}, Name: name, inputValue: empty}
}

func (prop *AnimatedProp) InsertKeyframe(key, value float32) {
	fmt.Println("Inserting keyframe!", key, value, prop)
	prop.SortedKeyframes = append(prop.SortedKeyframes, [2]float32{key, value})
	sort.Slice(prop.SortedKeyframes, func(i, j int) bool {
		return prop.SortedKeyframes[i][0] < prop.SortedKeyframes[j][0]
	})
	prop.KeyframesMap[key] = value
}

func (prop *AnimatedProp) Set(value float32, ui *lib.UIStruct) {
	if len(prop.KeyframesMap) == 0 {
		prop.Base = value
		return
	}

	key := float32(ui.SelectedFrame)
	for index, kf := range prop.SortedKeyframes {
		if kf[0] == key {
			prop.SortedKeyframes[index][1] = value
			prop.KeyframesMap[key] = value
			return
		}
	}

	prop.InsertKeyframe(key, value)
}

// MAYBE CACHE THIS SO YOU DONT HAVE TO RUN EVERY TIME IF NOTHING CHANGED
func (animatedProp AnimatedProp) KeyFramePosition(selectedFrame int) float32 {
	prop := animatedProp.Base
	keyframes := animatedProp.SortedKeyframes

	if len(keyframes) >= 2 {
		if selectedFrame <= int(keyframes[0][0]) {
			prop = keyframes[0][1]
		} else if selectedFrame >= int(keyframes[len(keyframes)-1][0]) {
			prop = keyframes[len(keyframes)-1][1]
		} else {
			var framesAround [2][2]float32
			for index, keyframe := range keyframes {
				if keyframe[0] > float32(selectedFrame) {
					framesAround = [2][2]float32{keyframes[index-1], keyframe}
					break
				}
			}
			framePos := lib.InverseLerp(framesAround[0][0], framesAround[1][0], float32(selectedFrame))
			prop = lib.Lerp(framesAround[0][1], framesAround[1][1], lib.Clamp(framePos, 0, 1))
		}

	} else if len(keyframes) == 1 {
		prop = keyframes[0][1]
	}

	return prop
}

func (animatedProp *AnimatedProp) Input(ui *lib.UIStruct, comp components.Components, layer Layer, prefix string) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := lib.
			NewLayout().
			PositionRect(rect).
			Row().
			Width(rect.Width,
				lib.ChildSize{SizeType: lib.SIZE_WEIGHT, Value: 1},
				lib.ChildSize{SizeType: lib.SIZE_ABSOLUTE, Value: 24}).
			Add(InputEditableContent(animatedProp, ui, comp, layer, prefix)).
			Add(KeyFrameButton(animatedProp, ui, comp, layer, prefix))
		return layout.Draw, 0, layout.Size.Height
	}
}

func KeyFrameButton(animatedProp *AnimatedProp, ui *lib.UIStruct, comp components.Components, layer Layer, prefix string) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := comp.Button(layer.GetName()+animatedProp.Name+prefix, rl.NewVector2(rect.X, rect.Y), []lib.Component{})
		if button.Clicked {
			animatedProp.InsertKeyframe(float32(ui.SelectedFrame), animatedProp.Base)
		}
		return button.Draw, 0, 0
	}
}

func InputEditableContent(animatedProp *AnimatedProp, ui *lib.UIStruct, comp components.Components, layer Layer, prefix string) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		inputValue := animatedProp.inputValue
		updateValue := fmt.Sprint(animatedProp.KeyFramePosition(ui.SelectedFrame))
		if inputValue == empty {
			inputValue = updateValue
		}

		input := comp.Input(components.InputProps{
			X:          rect.X,
			Y:          rect.Y,
			Id:         layer.GetName() + animatedProp.Name + prefix,
			Width:      rect.Width,
			Value:      inputValue,
			MousePoint: rl.GetMousePosition(),
			Ui:         ui,
		})

		if input.IsFocusing {
			animatedProp.inputValue = input.Value
		}

		if input.State == components.STATE_ACTIVE {
			animatedProp.inputValue = input.Value
		}

		if input.IsBluring || input.HasSubmitted {
			input.Blur(ui)

			if animatedProp.inputValue == "" || inputValue == updateValue {
				animatedProp.inputValue = empty
			} else {
				fmt.Println(inputValue, updateValue)
				var newIntValue, err = strconv.ParseFloat(animatedProp.inputValue, 32)
				if err != nil {
					animatedProp.inputValue = updateValue
				}

				animatedProp.Set(float32(newIntValue), ui)
				animatedProp.inputValue = empty
			}

		}

		return input.Draw, input.Rect.Width, input.Rect.Height
	}
}

func (animatedProp *AnimatedProp) CanDrawTimeline() bool {
	return len(animatedProp.SortedKeyframes) > 0
}

func (animatedProp *AnimatedProp) TimelineFrames(ui *lib.UIStruct) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		const height = 24
		var draw = func() {
			y := int32(rect.Y + (height / 2))
			rl.DrawLine(rect.ToInt32().X, y, rect.ToInt32().X+rect.ToInt32().Width, y, rl.Blue)

			keyframes := animatedProp.SortedKeyframes
			if len(keyframes) > 0 {
				for _, keyframe := range keyframes {
					x := ui.GetXTimelineFrame(rect, keyframe[0])
					keyframeRect := rl.NewRectangle(x, float32(y), 10, 10)
					rl.DrawRectanglePro(
						keyframeRect,        // A 10×20 rectangle
						rl.NewVector2(5, 5), // Center of the rectangle
						45,
						rl.Blue,
					)

					// if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(keyframeRect.X-5, keyframeRect.Y-5, keyframeRect.Width, keyframeRect.Height)) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					// 	// clicked
					// }
				}
			}
		}

		return draw, 0, height
	}
}
