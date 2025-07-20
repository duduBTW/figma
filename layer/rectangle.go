package layer

import (
	"strconv"

	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/layout"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width_KEY  = "width"
	height_KEY = "height"
)

type Rectangle struct {
	app.Element
	Width  app.AnimatedProp
	Height app.AnimatedProp
	Color  app.AnimatedColor
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
		Width:   app.NewAnimatedProp(float32(width), width_KEY),
		Height:  app.NewAnimatedProp(float32(height), height_KEY),
		Color:   app.NewAnimatedColor(217, 217, 217, 255),
		Element: app.NewElement(id, rl.NewVector2(rect.X, rect.Y), "Rectangle "+strconv.Itoa(index+1)),
	}
}

func (r *Rectangle) GetName() string {
	return r.Name
}
func (r *Rectangle) GetElement() *app.Element {
	return &r.Element
}

func (r *Rectangle) DrawHighlight() {
	rect := r.Rect(app.Apk.State.SelectedFrame)
	rl.DrawRectangleLinesEx(rect, 2, rl.Blue)

	box := components.Box(components.BoxProps{
		Padding:      *app.NewPadding().Axis(4, 2),
		Rect:         rl.NewRectangle(rect.X, rect.Y+rect.Height+4, 0, 0),
		Direction:    app.DIRECTION_ROW,
		Children:     []app.Component{RectangleDimensionsText(rect)},
		Color:        rl.Blue,
		BorderRadius: 2,
	})
	box.Draw()
}

func RectangleDimensionsText(rect rl.Rectangle) app.Component {
	textContent := strconv.Itoa(int(rect.Width)) + " x " + strconv.Itoa(int(rect.Height))
	var fontSize int32 = 10
	return components.Text(textContent, fontSize)
}

func (r *Rectangle) Rect(selectedFrame int) rl.Rectangle {
	x := r.Position.X.KeyFramePosition(selectedFrame)
	y := r.Position.Y.KeyFramePosition(selectedFrame)
	return rl.NewRectangle(x, y, r.Width.KeyFramePosition(selectedFrame), r.Height.KeyFramePosition(selectedFrame))
}
func (r *Rectangle) DrawComponent(mousePoint rl.Vector2, canvasRect rl.Rectangle) bool {
	r.Interactable = app.NewInteractable(r.Id)
	rect := r.Rect(app.Apk.SelectedFrame)

	// Only updates the event if the mouse is inside the canvas
	if rl.CheckCollisionPointRec(mousePoint, canvasRect) {
		r.Interactable.Event(mousePoint, rect)
	}

	rl.DrawRectangleRec(rect, r.Color.Get(app.Apk.SelectedFrame))
	return r.Interactable.State() == app.STATE_ACTIVE
}

// -----------
// Controls
// -----------

func (r *Rectangle) DrawControls(rect rl.Rectangle) {
	components.NewPanelLayout(rect).
		Add(components.NewAnimatedVector2(r.Position, r, "").Controls()).
		Add(r.SizeControls()).
		Add(components.NewAnimatedColor(&r.Color, r, "").Controls()).
		Draw()
}

func (r *Rectangle) SizeControls() app.Component {
	return func(avaliablePosition rl.Rectangle) (func(), float32, float32) {
		row := components.NewSidebarProperyLabel(avaliablePosition).
			Add(components.SidebarProperyLabel("Size")).
			Add(r.SizeControlsInputs())
		return row.Draw, 0, row.Size.Height
	}
}

func (r *Rectangle) SizeControlsInputs() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		row := components.SidebrInputsLayout(2, rect).
			Add(components.NewAnimatedProp(&r.Width, r, "").Input()).
			Add(components.NewAnimatedProp(&r.Height, r, "").Input())
		return row.Draw, 0, row.Size.Height
	}
}

// -----------
// Timeline
// -----------

func (r *Rectangle) DrawTimeline() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := layout.Timeline.Root(rect).
			Add(components.TimelinePanelTitle(r.Name, r))

		prefix := "timeline"
		layout.Add(components.NewAnimatedVector2(r.Position, r, prefix).Timeline()...)

		widthComponent := components.NewAnimatedProp(&r.Width, r, prefix)
		if widthComponent.CanDrawTimeline() {
			layout.Add(components.TimelineRow("Width", widthComponent.Input(), r.Width.SortedKeyframesTimeline()))
		}

		heightComponent := components.NewAnimatedProp(&r.Height, r, prefix)
		if heightComponent.CanDrawTimeline() {
			layout.Add(components.TimelineRow("Height", heightComponent.Input(), r.Height.SortedKeyframesTimeline()))
		}

		colorComponent := components.NewAnimatedColor(&r.Color, r, prefix)
		if colorComponent.CanDrawTimeline() {
			layout.Add(components.TimelineRow("Color", colorComponent.Input(), r.Color.Red.SortedKeyframesTimeline()))
		}

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}
