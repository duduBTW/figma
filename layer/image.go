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

	Texture rl.Texture2D
}

func NewImage(id string, position rl.Vector2, texture rl.Texture2D, index int) Image {
	return Image{
		Element: app.NewElement(id, position, "Image "+strconv.Itoa(index+1)),
		Texture: texture,
	}
}

func (i *Image) GetName() string {
	return i.Name
}
func (i *Image) GetElement() *app.Element {
	return &i.Element
}
func (i *Image) DrawHighlight() {
	rect := i.Rect(app.Apk.State.SelectedFrame)
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
	return rl.NewRectangle(x, y, float32(i.Texture.Width), float32(i.Texture.Height))
}
func (i *Image) DrawComponent(mousePoint rl.Vector2) bool {
	i.Interactable = app.NewInteractable(i.Id)
	rect := i.Rect(app.Apk.SelectedFrame)
	i.Interactable.Event(mousePoint, rect)
	rl.DrawTexture(i.Texture, rect.ToInt32().X, rect.ToInt32().Y, rl.White)
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
			Add(components.TimelinePanelTitle(i.Name, i))

		prefix := "timeline"
		layout.Add(components.NewAnimatedVector2(i.Position, i, prefix).Timeline()...)

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}
