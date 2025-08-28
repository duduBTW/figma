package newWorkplace

import (
	"fmt"
	"strconv"

	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/components"
	ds "github.com/dudubtw/figma/design-system"
	"github.com/dudubtw/figma/fmath"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ncruces/zenity"
)

func Form() app.Component {
	if rl.IsKeyPressed(rl.KeyEnter) && app.Apk.FocusedId != "" {
		submit()
	}

	return func(rect rl.Rectangle) (func(), float32, float32) {
		var width float32 = 400
		layout := app.
			NewLayout().
			Direction(app.DIRECTION_COLUMN).
			PositionRect(rect).
			Gap(ds.SPACING_4).
			Width(width).
			Add(components.Typography("NEW FILE", ds.FONT_SIZE_LG, ds.FONT_WEIGHT_MEDIUM, ds.T2_COLOR_CONTENT)).
			Add(MiniatureRectangle()).
			Add(TitleRow()).
			Add(FPSRow()).
			Add(DurationRow()).
			Add(ResolutionRow()).
			Add(Footer())

		return layout.Draw, width, layout.Size.Height
	}
}

func MiniatureRectangle() app.Component {
	return func(positionRect rl.Rectangle) (func(), float32, float32) {
		var height float32 = fmath.WidthTo16x9Height(positionRect.Width)
		rect := rl.NewRectangle(positionRect.X, positionRect.Y, positionRect.Width, height)
		return func() {
			miniatureTexture := app.Apk.CreateWorkplace.MiniatureTexture

			interactable := app.NewInteractable("select-miniature")
			if interactable.Event(rl.GetMousePosition(), rect) {
				pickNewMiniature()
			}

			if miniatureTexture != nil {
				components.DrawImageCoverRounded(miniatureTexture, rect, ds.RADII)
			} else {
				var colors = map[app.InteractableState]rl.Color{
					app.STATE_INITIAL: ds.T2_COLOR_SURFACE,
					app.STATE_ACTIVE:  ds.T2_COLOR_SURFACE_LIGHT,
					app.STATE_HOT:     ds.T2_COLOR_SURFACE_DARK,
				}
				components.DrawRectangleRoundedPixels(rect, ds.RADII, colors[interactable.State()])

				app.
					NewLayout().
					Row().
					Gap(ds.SPACING_2).
					PositionRect(rect).
					Width(rect.Width, app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: -1}).
					Height(rect.Height, app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: -1}).
					HorizontalAlignment(app.ALIGNMENT_CENTER).
					VerticalAlignment(app.ALIGNMENT_CENTER).
					Add(components.Icon(app.ICON_IMAGE)).
					Add(components.Typography("Select miniature", ds.FONT_SIZE_LG, ds.FONT_WEIGHT_REGULAR, ds.T2_COLOR_CONTENT_ACCENT)).
					Draw()
			}
		}, rect.Width, height
	}
}

func ResolutionRow() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := newFormRowLayout(rect).
			Add(components.Typography("Resolution", ds.FONT_SIZE, ds.FONT_WEIGHT_MEDIUM, ds.T2_COLOR_CONTENT)).
			Add(ResolutionInputs())

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}

func ResolutionInputs() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.NewLayout().
			Row().
			PositionRect(rect).
			Gap(ds.SPACING_2).
			Width(rect.Width,
				app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 0.5},
				app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 0.5}).
			Add(ResolutionWidthInput()).
			Add(ResolutionHeightInput())

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}

func ResolutionWidthInput() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		input := components.Input(components.InputProps{
			X:          rect.X,
			Y:          rect.Y,
			Id:         "resolution-w",
			Width:      rect.Width,
			Value:      strconv.Itoa(app.Apk.CreateWorkplace.FormData.ResolutionWidth),
			MousePoint: rl.GetMousePosition(),
		})

		resWidthInputValue, err := strconv.Atoi(input.Value)
		if err == nil {
			app.Apk.CreateWorkplace.FormData.ResolutionHeight = resWidthInputValue
		}

		return input.Draw, input.Rect.Width, input.Rect.Height
	}
}

func ResolutionHeightInput() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		input := components.Input(components.InputProps{
			X:          rect.X,
			Y:          rect.Y,
			Id:         "resolution-h",
			Width:      rect.Width,
			Value:      strconv.Itoa(app.Apk.CreateWorkplace.FormData.ResolutionHeight),
			MousePoint: rl.GetMousePosition(),
		})

		resHeightInputValue, err := strconv.Atoi(input.Value)
		if err == nil {
			app.Apk.CreateWorkplace.FormData.ResolutionHeight = resHeightInputValue
		}

		return input.Draw, input.Rect.Width, input.Rect.Height
	}
}

func TitleRow() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := newFormRowLayout(rect).
			Add(components.Typography("Title", ds.FONT_SIZE, ds.FONT_WEIGHT_MEDIUM, ds.T2_COLOR_CONTENT)).
			Add(TitleInput())

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}

func TitleInput() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		input := components.Input(components.InputProps{
			X:          rect.X,
			Y:          rect.Y,
			Id:         "title",
			Width:      rect.Width,
			Value:      app.Apk.CreateWorkplace.FormData.Title,
			MousePoint: rl.GetMousePosition(),
		})

		app.Apk.CreateWorkplace.FormData.Title = input.Value

		return input.Draw, input.Rect.Width, input.Rect.Height
	}
}

func FPSRow() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := newFormRowLayout(rect).
			Add(components.Typography("FPS", ds.FONT_SIZE, ds.FONT_WEIGHT_MEDIUM, ds.T2_COLOR_CONTENT)).
			Add(FPSInput())

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}

func FPSInput() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		input := components.Input(components.InputProps{
			X:          rect.X,
			Y:          rect.Y,
			Id:         "fps",
			Width:      rect.Width,
			Value:      strconv.Itoa(app.Apk.CreateWorkplace.FormData.Framerate),
			MousePoint: rl.GetMousePosition(),
		})

		frameRateUserInput, err := strconv.Atoi(input.Value)
		if err == nil {
			app.Apk.CreateWorkplace.FormData.Framerate = frameRateUserInput
		}

		return input.Draw, input.Rect.Width, input.Rect.Height
	}
}

func DurationRow() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := newFormRowLayout(rect).
			Add(components.Typography("Duration", ds.FONT_SIZE, ds.FONT_WEIGHT_MEDIUM, ds.T2_COLOR_CONTENT)).
			Add(DurationInput())

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}

func DurationInput() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		input := components.Input(components.InputProps{
			X:          rect.X,
			Y:          rect.Y,
			Id:         "duration",
			Width:      rect.Width,
			Value:      app.Apk.CreateWorkplace.FormData.Duration,
			MousePoint: rl.GetMousePosition(),
		})

		app.Apk.CreateWorkplace.FormData.Duration = input.Value

		return input.Draw, input.Rect.Width, input.Rect.Height
	}
}

func Footer() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.
			NewLayout().
			Row().
			PositionRect(rect).
			Width(rect.Width, app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: -1}).
			Gap(ds.SPACING_2).
			HorizontalAlignment(app.ALIGNMENT_END).
			Add(CancelButton()).
			Add(CreateButton())

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}

func CreateButton() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := components.Button(
			"create",
			components.BUTTON_VARIANT_PRIMARY,
			rl.NewVector2(rect.X, rect.Y),
			[]app.Component{components.ButtonText("Create")},
		)

		if button.Clicked {
			app.Apk.CreateWorkplace.Submit()
		}

		return button.Draw, button.Rect.Width, button.Rect.Height
	}
}

func CancelButton() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		button := components.Button(
			"cancel",
			components.BUTTON_VARIANT_OUTLINED,
			rl.NewVector2(rect.X, rect.Y),
			[]app.Component{components.ButtonText("Cancel")},
		)

		if button.Clicked {
			app.Apk.Navigate(app.PAGE_HOME)
		}

		return button.Draw, button.Rect.Width, button.Rect.Height
	}
}

// -----------
// UTILS
// -----------

func pickNewMiniature() {
	miniatureTexture := app.Apk.CreateWorkplace.MiniatureTexture
	dir, err := zenity.SelectFile(
		zenity.Title("Select the miniature"),
		zenity.FileFilter{
			Patterns: []string{"*.png", "*.jpg", ".jpeg"},
		},
	)

	if err != nil {
		return
	}

	if miniatureTexture != nil {
		rl.UnloadTexture(*miniatureTexture)
	}

	newMiniatureTexture := rl.LoadTexture(dir)
	app.Apk.CreateWorkplace.MiniatureTexture = &newMiniatureTexture
	app.Apk.CreateWorkplace.FormData.PreviewImagePath = dir
}

func newFormRowLayout(rect rl.Rectangle) *app.Layout {
	return app.NewLayout().
		PositionRect(rect).
		Row().
		Gap(ds.SPACING_5).
		Width(rect.Width,
			app.ChildSize{SizeType: app.SIZE_ABSOLUTE, Value: 80},
			app.ChildSize{SizeType: app.SIZE_WEIGHT, Value: 1})
}

func submit() {
	err := app.Apk.CreateWorkplace.Submit()
	if err != nil {
		fmt.Println(err)
	}

	app.Apk.Navigate(app.PAGE_HOME)
}
