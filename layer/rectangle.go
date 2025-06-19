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
			Position: rl.NewVector2(rect.X, rect.Y),
			Name:     "Rectangle " + strconv.Itoa(index+1),
		},
		InputValues: map[string]string{
			"y": strconv.Itoa(int(rect.Y)),
			"x": strconv.Itoa(int(rect.X)),
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
func (r *Rectangle) DrawHighlight() {
	rl.DrawRectangleLinesEx(r.Rect(), 2, rl.Blue)
}
func (r *Rectangle) Rect() rl.Rectangle {
	return rl.NewRectangle(r.Position.X, r.Position.Y, float32(r.Width), float32(r.Height))
}
func (r *Rectangle) DrawComponent(ui *lib.UIStruct, mousePoint rl.Vector2) bool {
	interactable := components.NewInteractable(r.Id, ui)
	clicked := interactable.Event(mousePoint, rl.NewRectangle(r.Position.X, r.Position.Y, float32(r.Width), float32(r.Height)))
	rl.DrawRectangleRec(r.Rect(), r.Color)
	r.interactable = interactable

	return clicked
}

func (r *Rectangle) DrawControls(ui *lib.UIStruct, rect rl.Rectangle, comp components.Components) {
	// layout := lib.NewLayout(lib.PublicLayouyt{
	// 	Gap:       8,
	// 	Direction: lib.DIRECTION_COLUMN,
	// }, rl.NewVector2(rect.X, rect.Y))

	row := lib.NewConstrainedLayout(lib.ContrainedLayout{
		Direction: lib.DIRECTION_ROW,
		Gap:       32,
		Contrains: rl.NewRectangle(rect.X, rect.Y, rect.Width, 24),
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

	row.Render(Label())
	row.Render(Inputs(r, ui, comp))

	// layout.Draw()
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

		row.Render(PanelInput(r, ui, comp, "y", &r.Position.Y))
		row.Render(PanelInput(r, ui, comp, "x", &r.Position.X))
	}
}

func Label() lib.ContrainedComponent {
	return func(rect rl.Rectangle) {
		text := "Position"
		fontSize := 14
		rl.DrawText(text, rect.ToInt32().X, rect.ToInt32().Y+4, int32(fontSize), rl.White)
	}
}

func PanelInput(r *Rectangle, ui *lib.UIStruct, comp components.Components, key string, value *float32) lib.ContrainedComponent {
	return func(avaliablePosition rl.Rectangle) {
		input := comp.Input(components.InputProps{
			X:             avaliablePosition.X,
			Y:             avaliablePosition.Y,
			Id:            r.Id + key,
			Width:         avaliablePosition.Width,
			Value:         r.InputValues[key],
			MousePoint:    rl.GetMousePosition(),
			Ui:            ui,
			LeftIndicator: rune(key[0]),
		})

		r.InputValues[key] = input.Value
		if input.IsBluring || input.HasSubmitted {
			input.Blur(ui)

			var newIntValue, err = strconv.ParseFloat(r.InputValues[key], 32)
			if err == nil {
				*value = float32(newIntValue)
			} else {
				r.InputValues[key] = fmt.Sprint(value)
			}
		}

		input.Draw()
		// return input.Draw, input.Rect
	}
}
