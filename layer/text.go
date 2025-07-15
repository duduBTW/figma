package layer

import (
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/layout"
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	Element
	Color       AnimatedColor
	FontSize    AnimatedProp
	TextContent string

	InputValues map[string]string
}

const defaultTextName = "Text"

const (
	font_SIZE_KEY = "fontSize"
)

func NewText(id string, rect rl.Vector2) Text {
	return Text{
		Element:     NewElement(id, rect, defaultTextName),
		InputValues: map[string]string{},
		Color:       NewAnimatedColor(255, 255, 255, 255),
		FontSize:    NewAnimatedProp(20, font_SIZE_KEY),
		TextContent: "Hello world!",
	}
}

func (t *Text) GetName() string {
	if t.Name == defaultTextName && t.TextContent == "" {
		return t.Name
	}
	return t.TextContent
}

func (t *Text) GetElement() *Element {
	return &t.Element
}

func (t *Text) DrawHighlight(ui lib.UIStruct, comp components.Components) {
	rect := t.Rect(ui.SelectedFrame)
	rl.DrawRectangleLinesEx(rect, 2, rl.Blue)
}

func (t *Text) DrawComponent(ui *lib.UIStruct, mousePoint rl.Vector2) bool {
	interactable := components.NewInteractable(t.Id, ui)
	rect := t.Rect(ui.SelectedFrame)
	interactable.Event(mousePoint, rect)
	fontSize := t.FontSize.KeyFramePosition(ui.SelectedFrame)
	color := t.Color.Get(ui.SelectedFrame)
	rl.DrawText(t.TextContent, rect.ToInt32().X, rect.ToInt32().Y, int32(fontSize), color)
	t.interactable = interactable
	return interactable.State() == components.STATE_ACTIVE
}

func (t *Text) Rect(selectedFrame int) rl.Rectangle {
	x := t.Position.X.KeyFramePosition(selectedFrame)
	y := t.Position.Y.KeyFramePosition(selectedFrame)

	fontSize := t.FontSize.KeyFramePosition(selectedFrame)
	return rl.NewRectangle(x, y, float32(rl.MeasureText(t.TextContent, int32(fontSize))), fontSize)
}

// -----------
// Controls
// -----------

func (t *Text) DrawControls(ui *lib.UIStruct, rect rl.Rectangle, comp components.Components) {
	NewPanelLayout(rect).
		Add(t.Position.Controls(ui, comp, t)).
		Add(t.FontSizeControls(ui, comp)).
		Add(t.Color.Controls(ui, comp, t)).
		Draw()
}

func (t *Text) FontSizeControls(ui *lib.UIStruct, comp components.Components) lib.Component {
	return func(avaliablePosition rl.Rectangle) (func(), float32, float32) {
		row := NewControlsLayout(avaliablePosition).
			Add(Label("Font size")).
			Add(t.FontSize.Input(ui, comp, t, ""))
		return row.Draw, 0, row.Size.Height
	}
}

// -----------
// Timeline
// -----------

func (t *Text) DrawTimeline(ui *lib.UIStruct, comp components.Components) lib.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := layout.Timeline.Root(rect)
		layout.Add(TimelinePanelTitle(t.Name, t, ui))
		prefix := "timeline"
		if t.Position.CanDrawTimeline() {
			t.Position.Timeline(layout, ui, comp, t, prefix)
		}

		if t.FontSize.CanDrawTimeline() {
			layout.Add(comp.TimelineRow("Font size", t.FontSize.Input(ui, comp, t, prefix), t.FontSize.SortedKeyframes))
		}

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}
