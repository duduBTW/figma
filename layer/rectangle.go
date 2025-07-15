package layer

import (
	"strconv"

	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/layout"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width_KEY  = "width"
	height_KEY = "height"
)

type Rectangle struct {
	Element
	Width  AnimatedProp
	Height AnimatedProp
	Color  AnimatedColor

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
		Width:       NewAnimatedProp(float32(width), width_KEY),
		Height:      NewAnimatedProp(float32(height), height_KEY),
		Color:       NewAnimatedColor(217, 217, 217, 255),
		Element:     NewElement(id, rl.NewVector2(rect.X, rect.Y), "Rectangle "+strconv.Itoa(index+1)),
		InputValues: map[string]string{},
	}
}

func (r *Rectangle) GetName() string {
	return r.Name
}
func (r *Rectangle) GetElement() *Element {
	return &r.Element
}

func (r *Rectangle) DrawHighlight(ui lib.UIStruct, comp components.Components) {
	rect := r.Rect(ui.SelectedFrame)
	rl.DrawRectangleLinesEx(rect, 2, rl.Blue)

	padding := lib.Padding{}
	padding.Axis(4, 2)
	box := comp.Box(components.BoxProps{
		Padding:      padding,
		Rect:         rl.NewRectangle(rect.X, rect.Y+rect.Height+4, 0, 0),
		Direction:    lib.DIRECTION_ROW,
		Children:     []lib.Component{RectangleDimensionsText(rect)},
		Color:        rl.Blue,
		BorderRadius: 2,
	})
	box.Draw()
}

func RectangleDimensionsText(rect rl.Rectangle) lib.Component {
	return func(avaliablePosition rl.Rectangle) (func(), float32, float32) {
		textContent := strconv.Itoa(int(rect.Width)) + " x " + strconv.Itoa(int(rect.Height))
		fontSize := 10

		return func() {
			rl.DrawText(textContent, int32(avaliablePosition.X), int32(avaliablePosition.Y), int32(fontSize), rl.White)
		}, float32(rl.MeasureText(textContent, int32(fontSize))), float32(fontSize)
	}
}

func (r *Rectangle) Rect(selectedFrame int) rl.Rectangle {
	x := r.Position.X.KeyFramePosition(selectedFrame)
	y := r.Position.Y.KeyFramePosition(selectedFrame)
	return rl.NewRectangle(x, y, r.Width.KeyFramePosition(selectedFrame), r.Height.KeyFramePosition(selectedFrame))
}
func (r *Rectangle) DrawComponent(ui *lib.UIStruct, mousePoint rl.Vector2) bool {
	interactable := components.NewInteractable(r.Id, ui)
	rect := r.Rect(ui.SelectedFrame)
	interactable.Event(mousePoint, rect)
	rl.DrawRectangleRec(rect, r.Color.Get(ui.SelectedFrame))
	r.interactable = interactable
	return interactable.State() == components.STATE_ACTIVE
}

// -----------
// Controls
// -----------

func (r *Rectangle) DrawControls(ui *lib.UIStruct, rect rl.Rectangle, comp components.Components) {
	NewPanelLayout(rect).
		Add(r.Position.Controls(ui, comp, r)).
		Add(r.SizeControls(ui, comp)).
		Add(r.Color.Controls(ui, comp, r)).
		Draw()
}

func (r *Rectangle) SizeControls(ui *lib.UIStruct, comp components.Components) lib.Component {
	return func(avaliablePosition rl.Rectangle) (func(), float32, float32) {
		row := NewControlsLayout(avaliablePosition).
			Add(Label("Size")).
			Add(r.SizeControlsInputs(ui, comp))
		return row.Draw, 0, row.Size.Height
	}
}

func (r *Rectangle) SizeControlsInputs(ui *lib.UIStruct, comp components.Components) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := InputsLayout(2, rect).
			Add(r.Width.Input(ui, comp, r, "")).
			Add(r.Height.Input(ui, comp, r, ""))
		return row.Draw, 0, row.Size.Height
	}
}

// -----------
// Timeline
// -----------

func (r *Rectangle) DrawTimeline(ui *lib.UIStruct, comp components.Components) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := layout.Timeline.Root(rect).
			Add(TimelinePanelTitle(r.Name, r, ui))

		prefix := "timeline"
		if r.Position.CanDrawTimeline() {
			r.Position.Timeline(layout, ui, comp, r, prefix)
		}

		if r.Width.CanDrawTimeline() {
			layout.Add(comp.TimelineRow("Width", r.Width.Input(ui, comp, r, prefix), r.Width.SortedKeyframes))
		}
		if r.Height.CanDrawTimeline() {
			layout.Add(comp.TimelineRow("Height", r.Height.Input(ui, comp, r, prefix), r.Height.SortedKeyframes))
		}

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}
