package layer

import (
	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/layout"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	app.Element
	Color       app.AnimatedColor
	FontSize    app.AnimatedProp
	TextContent string
}

const defaultTextName = "Text"

const (
	font_SIZE_KEY = "fontSize"
)

func NewText(id string, rect rl.Vector2) Text {
	return Text{
		Element:     app.NewElement(id, rect, defaultTextName),
		Color:       app.NewAnimatedColor(255, 255, 255, 255),
		FontSize:    app.NewAnimatedProp(20, font_SIZE_KEY),
		TextContent: "Hello world!",
	}
}

func (t *Text) GetName() string {
	if t.Element.Name == defaultTextName && t.TextContent == "" {
		return t.Element.Name
	}
	return t.TextContent
}

func (t *Text) GetElement() *app.Element {
	return &t.Element
}

func (t *Text) DrawHighlight() {
	rect := t.Rect(app.Apk.SelectedFrame)
	rl.DrawRectangleLinesEx(rect, 2, rl.Blue)
}

func (t *Text) DrawComponent(mousePoint rl.Vector2) bool {
	interactable := app.NewInteractable(t.Element.Id)
	rect := t.Rect(app.Apk.SelectedFrame)
	interactable.Event(mousePoint, rect)
	fontSize := t.FontSize.KeyFramePosition(app.Apk.SelectedFrame)
	color := t.Color.Get(app.Apk.SelectedFrame)
	rl.DrawText(t.TextContent, rect.ToInt32().X, rect.ToInt32().Y, int32(fontSize), color)
	t.Element.Interactable = interactable
	return interactable.State() == app.STATE_ACTIVE
}

func (t *Text) Rect(selectedFrame int) rl.Rectangle {
	x := t.Element.Position.X.KeyFramePosition(selectedFrame)
	y := t.Element.Position.Y.KeyFramePosition(selectedFrame)

	fontSize := t.FontSize.KeyFramePosition(selectedFrame)
	return rl.NewRectangle(x, y, float32(rl.MeasureText(t.TextContent, int32(fontSize))), fontSize)
}

// -----------
// Controls
// -----------

func (t *Text) DrawControls(rect rl.Rectangle) {
	components.NewPanelLayout(rect).
		Add(components.NewAnimatedVector2(t.Position, t, "").Controls()).
		Add(t.FontSizeControls()).
		Add(components.NewAnimatedColor(&t.Color, t, "").Controls()).
		Draw()
}

func (t *Text) FontSizeControls() app.Component {
	return func(avaliablePosition rl.Rectangle) (func(), float32, float32) {
		row := components.NewSidebarProperyLabel(avaliablePosition).
			Add(components.SidebarProperyLabel("Font size")).
			Add(components.NewAnimatedProp(&t.FontSize, t, "").Input())
		return row.Draw, 0, row.Size.Height
	}
}

// -----------
// Timeline
// -----------

func (t *Text) DrawTimeline() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := layout.Timeline.Root(rect)
		layout.Add(components.TimelinePanelTitle(t.Name, t))
		prefix := "timeline"

		layout.Add(components.NewAnimatedVector2(t.Position, t, prefix).Timeline()...)
		fontSizeComponent := components.NewAnimatedProp(&t.FontSize, t, prefix)
		if fontSizeComponent.CanDrawTimeline() {
			layout.Add(components.TimelineRow("Font size", fontSizeComponent.Input(), t.FontSize.SortedKeyframes))
		}

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}
