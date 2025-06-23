package layer

import (
	"fmt"
	"strconv"

	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Rectangle struct {
	Element
	Width  int32
	Height int32
	Color  rl.Color

	interactable components.Interactable

	InputValues map[string]string
}

func NewRectangle(id string, rect rl.Rectangle, index int) Rectangle {
	width := rect.ToInt32().Width
	if width <= 4 {
		width = 100
		rect.X -= 50
	}

	height := rect.ToInt32().Height
	if height <= 4 {
		height = 100
		rect.Y -= 50
	}

	return Rectangle{
		Width:  width,
		Height: height,
		Color:  rl.Fade(rl.White, 0.32),
		Element: Element{
			Id:       id,
			Position: NewAnimatedVector2(rect.X, rect.Y),
			Name:     "Rectangle " + strconv.Itoa(index+1),
		},
		InputValues: map[string]string{
			"y": "",
			"x": "",
		},
	}
}

func (r *Rectangle) GetName() string {
	return r.Name
}
func (r *Rectangle) GetElement() *Element {
	return &r.Element
}
func (r *Rectangle) State() components.InteractableState {
	return r.interactable.State()
}
func (r *Rectangle) DrawHighlight(ui lib.UIStruct) {
	rl.DrawRectangleLinesEx(r.Rect(ui.SelectedFrame), 2, rl.Blue)
}

func KeyFramePosition(animatedProp AnimatedProp, selectedFrame int) float32 {
	x := animatedProp.Base
	keyframes := animatedProp.Keyframes

	if len(keyframes) >= 2 {
		if selectedFrame <= int(keyframes[0][0]) {
			x = keyframes[0][1]
		} else if selectedFrame >= int(keyframes[len(keyframes)-1][0]) {
			x = keyframes[len(keyframes)-1][1]
		} else {
			var framesAround [2][2]float32
			for index, keyframe := range keyframes {
				if keyframe[0] > float32(selectedFrame) {
					framesAround = [2][2]float32{keyframes[index-1], keyframe}
					break
				}
			}
			framePos := lib.InverseLerp(framesAround[0][0], framesAround[1][0], float32(selectedFrame))
			x = lib.Lerp(framesAround[0][1], framesAround[1][1], lib.Clamp(framePos, 0, 1))
		}

	} else if len(keyframes) == 1 {
		x = keyframes[0][1]
	}

	return x
}
func (r *Rectangle) Rect(selectedFrame int) rl.Rectangle {
	x := KeyFramePosition(r.Position.X, selectedFrame)
	y := KeyFramePosition(r.Position.Y, selectedFrame)

	return rl.NewRectangle(x, y, float32(r.Width), float32(r.Height))
}
func (r *Rectangle) DrawComponent(ui *lib.UIStruct, mousePoint rl.Vector2) bool {
	interactable := components.NewInteractable(r.Id, ui)
	rect := r.Rect(ui.SelectedFrame)
	clicked := interactable.Event(mousePoint, rect)
	rl.DrawRectangleRec(rect, r.Color)
	r.interactable = interactable

	return clicked
}

func (r *Rectangle) DrawControls(ui *lib.UIStruct, rect rl.Rectangle, comp components.Components) {
	layout := lib.NewLayout(lib.PublicLayouyt{
		Gap:       8,
		Direction: lib.DIRECTION_COLUMN,
	}, rl.NewVector2(rect.X, rect.Y))

	layout.Add(PositionProps(r, ui, rect, comp))
	layout.Add(PositionKeyframes(r, ui, rect, comp))
	layout.Draw()
}
func (*Rectangle) InsertKeyframe(animatedProp *AnimatedProp, ui *lib.UIStruct) {
	animatedProp.InsertKeyframe(ui)
}

func PositionProps(r *Rectangle, ui *lib.UIStruct, rect rl.Rectangle, comp components.Components) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		contrains := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, rect.Width, 24)
		row := lib.NewConstrainedLayout(lib.ContrainedLayout{
			Direction: lib.DIRECTION_ROW,
			Gap:       32,
			Contrains: contrains,
			ChildrenSize: []lib.ChildSize{
				{
					SizeType: lib.SIZE_ABSOLUTE,
					Value:    60,
				},
				{
					SizeType: lib.SIZE_WEIGHT,
					Value:    1,
				},
			},
		})

		row.Add(Label("Position p"))
		row.Add(Inputs(r, ui, comp))

		return row.Draw, contrains
	}
}

func PositionKeyframes(r *Rectangle, ui *lib.UIStruct, rect rl.Rectangle, comp components.Components) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		contrains := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, rect.Width, 24)
		row := lib.NewConstrainedLayout(lib.ContrainedLayout{
			Direction: lib.DIRECTION_ROW,
			Gap:       32,
			Contrains: contrains,
			ChildrenSize: []lib.ChildSize{
				{
					SizeType: lib.SIZE_ABSOLUTE,
					Value:    60,
				},
				{
					SizeType: lib.SIZE_WEIGHT,
					Value:    1,
				},
			},
		})

		row.Add(Label("Position k"))
		row.Add(Keyframes(r, ui, comp))
		return row.Draw, contrains
	}
}

func Keyframes(r *Rectangle, ui *lib.UIStruct, comp components.Components) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		row := lib.NewConstrainedLayout(lib.ContrainedLayout{
			Direction: lib.DIRECTION_ROW,
			Gap:       8,
			Contrains: rl.NewRectangle(rect.X, rect.Y, rect.Width, 24),
			ChildrenSize: []lib.ChildSize{
				{
					SizeType: lib.SIZE_WEIGHT,
					Value:    0.5,
				},
				{
					SizeType: lib.SIZE_WEIGHT,
					Value:    0.5,
				},
			},
		})

		row.Add(KeyframeButton("x", &r.Position.X, ui, comp))
		row.Add(KeyframeButton("y", &r.Position.Y, ui, comp))
		row.Draw()
	}
}

func KeyframeButton(text string, animatedProp *AnimatedProp, ui *lib.UIStruct, comp components.Components) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		button := comp.Button("keyframe"+text, rl.NewVector2(rect.X, rect.Y), []lib.Component{KeyframeButtonContent(text)})

		if button.Clicked {
			animatedProp.InsertKeyframe(ui)
		}

		button.Draw()
	}
}

func KeyframeButtonContent(textContet string) lib.Component {
	return func(avaliablePosition rl.Vector2) (func(), rl.Rectangle) {
		fontSize := 16
		rect := rl.NewRectangle(avaliablePosition.X, avaliablePosition.Y, float32(rl.MeasureText(textContet, int32(fontSize))), float32(fontSize))
		return func() {
			rl.DrawText(textContet, int32(avaliablePosition.X), int32(avaliablePosition.Y), int32(fontSize), rl.White)
		}, rect
	}
}

func Inputs(r *Rectangle, ui *lib.UIStruct, comp components.Components) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		row := lib.NewConstrainedLayout(lib.ContrainedLayout{
			Direction: lib.DIRECTION_ROW,
			Gap:       8,
			Contrains: rl.NewRectangle(rect.X, rect.Y, rect.Width, 24),
			ChildrenSize: []lib.ChildSize{
				{
					SizeType: lib.SIZE_WEIGHT,
					Value:    0.5,
				},
				{
					SizeType: lib.SIZE_WEIGHT,
					Value:    0.5,
				},
			},
		})

		row.Add(PanelInput(r, ui, comp, "x", r.Rect(ui.SelectedFrame).X, &r.Position.X))
		row.Add(PanelInput(r, ui, comp, "y", r.Rect(ui.SelectedFrame).Y, &r.Position.Y))
		row.Draw()
	}
}

func Label(text string) lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		fontSize := 14
		rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y+4, int32(fontSize), rl.White)
	}
}

func PanelInput(r *Rectangle, ui *lib.UIStruct, comp components.Components, key string, displayValue float32, updateValue *AnimatedProp) lib.ContrainedComponent {
	return func(avaliablePosition rl.Rectangle) {
		inputValue := r.InputValues[key]
		if inputValue == "" {
			inputValue = strconv.Itoa(int(displayValue))
		}

		input := comp.Input(components.InputProps{
			X:             avaliablePosition.X,
			Y:             avaliablePosition.Y,
			Id:            r.Id + key,
			Width:         avaliablePosition.Width,
			Value:         inputValue,
			MousePoint:    rl.GetMousePosition(),
			Ui:            ui,
			LeftIndicator: rune(key[0]),
		})

		if input.IsFocusing {
			r.InputValues[key] = input.Value
		}

		if input.State == components.STATE_ACTIVE {
			r.InputValues[key] = input.Value
		}

		if input.IsBluring || input.HasSubmitted {
			input.Blur(ui)

			var newIntValue, err = strconv.ParseFloat(r.InputValues[key], 32)
			if err == nil {
				updateValue.Base = float32(newIntValue)
			} else {
				r.InputValues[key] = fmt.Sprint(updateValue)
			}

			if len(updateValue.Keyframes) >= 1 {
				updateValue.InsertKeyframe(ui)
			}

			r.InputValues[key] = ""
		}

		input.Draw()
	}
}
