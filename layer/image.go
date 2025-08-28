package layer

import (
	"strconv"

	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/components"
	"github.com/dudubtw/figma/layout"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Image struct {
	app.Element
	Width  app.AnimatedProp
	Height app.AnimatedProp

	Path string
}

func NewImage(id string, position rl.Vector2, path string, index int) Image {
	texture := app.Apk.Workplace.LoadImagePath(path)
	return Image{
		Width:   app.NewAnimatedProp(float32(texture.Width), width_KEY),
		Height:  app.NewAnimatedProp(float32(texture.Height), height_KEY),
		Element: app.NewElement(id, position, "Image "+strconv.Itoa(index+1), "image"),
		Path:    path,
	}
}

func (i *Image) GetName() string {
	return i.Name
}
func (i *Image) GetElement() *app.Element {
	return &i.Element
}
func (i *Image) DrawHighlight() {
	rect := i.Rect(app.Apk.Workplace.SelectedFrame)
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
func (i *Image) Rect(selectedFrame int) rl.Rectangle {
	x := i.Position.X.KeyFramePosition(selectedFrame)
	y := i.Position.Y.KeyFramePosition(selectedFrame)
	width := i.Width.KeyFramePosition(selectedFrame)
	height := i.Height.KeyFramePosition(selectedFrame)

	return rl.NewRectangle(x, y, width, height)
}
func (i *Image) DrawComponent(mousePoint rl.Vector2, canvasRect rl.Rectangle) bool {
	i.Interactable = app.NewInteractable(i.Id)
	rect := i.Rect(app.Apk.Workplace.SelectedFrame)

	// Only updates the event if the mouse is inside the canvas
	// i.Interactable.State() != app.STATE_ACTIVE &&
	if rl.CheckCollisionPointRec(mousePoint, canvasRect) {
		i.Interactable.Event(mousePoint, rect)
	}

	texture := app.Apk.Workplace.GetImagePath(i.Path)
	rl.DrawTexture(texture, rect.ToInt32().X, rect.ToInt32().Y, rl.White)

	// components.DrawImageCoverRounded(&texture, rect, 0)
	return i.Interactable.State() == app.STATE_ACTIVE
}
func (i *Image) DrawControls(rect rl.Rectangle) {
	components.NewPanelLayout(rect).
		Add(components.NewAnimatedVector2(i.Position, i, "").Controls()).
		Draw()
}
func (i *Image) DrawTimeline() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := layout.Timeline.Root(rect).
			Add(components.TimelinePanelTitle(app.ICON_IMAGE, i.Name, i))

		prefix := "timeline"
		layout.Add(components.NewAnimatedVector2(i.Position, i, prefix).Timeline()...)

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}
