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

func (animatedProp *AnimatedProp) Input(ui *lib.UIStruct, comp components.Components) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		layout := lib.NewConstrainedLayout(lib.ContrainedLayout{
			Contrains: rect,
			Direction: lib.DIRECTION_ROW,
			ChildrenSize: []lib.ChildSize{
				{
					SizeType: lib.SIZE_WEIGHT,
					Value:    1,
				},
				{
					SizeType: lib.SIZE_ABSOLUTE,
					Value:    24,
				},
			},
		})
		layout.Add(InputEditableContent(animatedProp, ui, comp))
		layout.Add(KeyFrameButton(animatedProp, ui, comp))
		layout.Draw()
	}
}
func (animatedProp *AnimatedProp) NewInput(ui *lib.UIStruct, comp components.Components) lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := lib.NewMixLayout(lib.PublicMixLayouyt{
			Direction: lib.DIRECTION_ROW,
			InitialRect: lib.MixLayouytRect{
				Position: rl.NewVector2(rect.X, rect.Y),
				Width: lib.ContrainedSize{
					Value: rect.Width,
					Contrains: []lib.ChildSize{
						{
							SizeType: lib.SIZE_WEIGHT,
							Value:    1,
						},
						{
							SizeType: lib.SIZE_ABSOLUTE,
							Value:    24,
						},
					},
				},
			},
		})
		layout.Add(NewInputEditableContent(animatedProp, ui, comp))
		layout.Add(NewKeyFrameButton(animatedProp, ui, comp))
		return layout.Draw, 0, layout.CurrentRect.Height
	}
}

func KeyFrameButton(animatedProp *AnimatedProp, ui *lib.UIStruct, comp components.Components) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		button := comp.Button(animatedProp.Name, rl.NewVector2(rect.X, rect.Y), []lib.Component{})
		if button.Clicked {
			animatedProp.InsertKeyframe(float32(ui.SelectedFrame), animatedProp.Base)
		}
		button.Draw()
	}
}

func NewKeyFrameButton(animatedProp *AnimatedProp, ui *lib.UIStruct, comp components.Components) lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := comp.Button(animatedProp.Name, rl.NewVector2(rect.X, rect.Y), []lib.Component{})
		if button.Clicked {
			animatedProp.InsertKeyframe(float32(ui.SelectedFrame), animatedProp.Base)
		}
		return button.Draw, button.Rect.Width, button.Rect.Height
	}
}

func InputEditableContent(animatedProp *AnimatedProp, ui *lib.UIStruct, comp components.Components) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		inputValue := animatedProp.inputValue
		updateValue := animatedProp.KeyFramePosition(ui.SelectedFrame)
		if inputValue == empty {
			inputValue = strconv.Itoa(int(updateValue))
		}

		input := comp.Input(components.InputProps{
			X:          rect.X,
			Y:          rect.Y,
			Id:         animatedProp.Name,
			Width:      rect.Width,
			Value:      inputValue,
			MousePoint: rl.GetMousePosition(),
			Ui:         ui,
			// LeftIndicator: rune(key[0]),
		})

		if input.IsFocusing {
			animatedProp.inputValue = input.Value
		}

		if input.State == components.STATE_ACTIVE {
			animatedProp.inputValue = input.Value
		}

		if input.IsBluring || input.HasSubmitted {
			input.Blur(ui)

			if animatedProp.inputValue == "" {
				animatedProp.inputValue = empty
			} else {
				var newIntValue, err = strconv.ParseFloat(animatedProp.inputValue, 32)
				if err != nil {
					animatedProp.inputValue = fmt.Sprint(updateValue)
				}

				animatedProp.Set(float32(newIntValue), ui)
				animatedProp.inputValue = empty
			}

		}

		input.Draw()
	}
}

func NewInputEditableContent(animatedProp *AnimatedProp, ui *lib.UIStruct, comp components.Components) lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		inputValue := animatedProp.inputValue
		updateValue := animatedProp.KeyFramePosition(ui.SelectedFrame)
		if inputValue == empty {
			inputValue = strconv.Itoa(int(updateValue))
		}

		input := comp.Input(components.InputProps{
			X:          rect.X,
			Y:          rect.Y,
			Id:         animatedProp.Name,
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

			if animatedProp.inputValue == "" {
				animatedProp.inputValue = empty
			} else {
				var newIntValue, err = strconv.ParseFloat(animatedProp.inputValue, 32)
				if err != nil {
					animatedProp.inputValue = fmt.Sprint(updateValue)
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

func (animatedProp *AnimatedProp) TimelineFrames() lib.MixComponent {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		const height = 24
		var draw = func() {
			y := int32(rect.Y + (height / 2))
			rl.DrawLine(rect.ToInt32().X, y, rect.ToInt32().X+rect.ToInt32().Width, y, rl.NewColor(68, 68, 68, 255))
		}

		return draw, 0, height
	}
}
