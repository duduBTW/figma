package home

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dudubtw/figma/app"
	"github.com/dudubtw/figma/components"
	ds "github.com/dudubtw/figma/design-system"
	"github.com/dudubtw/figma/fmath"
	"github.com/dudubtw/figma/layer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func FileGrid() app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		if len(app.Apk.Home.Files) == 0 {
			return func() {}, 0, 0
		}

		grid := app.
			NewGrid(4, rect.Width, ds.SPACING_3, rl.NewVector2(rect.X, rect.Y))

		for index, file := range app.Apk.Home.Files {
			grid.Add(Card(file, app.Apk.Home.MinuatureTextures[index]))
		}

		return grid.Draw, 0, grid.Height()
	}
}

func Card(file app.HomeWorkplaceFile, miniatureTexture *rl.Texture2D) app.GridComponent {
	return func(x, y, width float32) (func(), float32) {
		layout := app.NewLayout().
			Column().
			Gap(ds.SPACING_2).
			Position(rl.NewVector2(x, y)).
			Width(width).
			Add(Miniature(miniatureTexture)).
			Add(components.Typography(file.Title, ds.FONT_SIZE_LG, ds.FONT_WEIGHT_BOLD, ds.T2_COLOR_CONTENT))

		isSelected := app.Apk.Home.SelectedFile.Title == file.Title

		interactable := app.NewInteractable(file.Title + "-grid-file")
		interactableRect := rl.NewRectangle(x-4, y-4, width+8, layout.Size.Height+8)
		if interactable.Event(rl.GetMousePosition(), interactableRect) {
			if isSelected {
				app.Apk.Workplace.Id = file.Title
				app.Apk.Navigate(app.PAGE_WORKPLACE)
				workplaceLoad()
			}

			app.Apk.Home.SelectedFile = file
		}

		state := interactable.State()

		return func() {
			switch state {
			case app.STATE_HOT:
				components.DrawRectangleRoundedPixels(interactableRect, ds.RADII+4, ds.T2_COLOR_SURFACE_DARK)
			case app.STATE_ACTIVE:
				components.DrawRectangleRoundedPixels(interactableRect, ds.RADII+4, ds.T2_COLOR_SURFACE_LIGHT)
			}

			layout.Draw()

			if isSelected {
				components.DrawRectangleRoundedLinePixels(interactableRect, ds.RADII+4, 2, ds.T2_COLOR_HIGHLIGHT)
			}
		}, layout.Size.Height
	}
}

func Miniature(miniatureTexture *rl.Texture2D) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		rect.Height = fmath.WidthTo16x9Height(rect.Width)

		return func() {
			if miniatureTexture != nil {
				components.DrawImageCoverRounded(miniatureTexture, rect, ds.RADII)
			} else {
				components.DrawRectangleRoundedPixels(rect, ds.RADII, ds.T2_COLOR_SURFACE)
			}
		}, rect.Width, rect.Height
	}
}

// ----------
// Helpers
// ----------

func workplaceLoad() {
	path := filepath.Join(app.PROJECTS_FOLDER, app.Apk.Id+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", path, err)
		panic(1)
	}

	var workplaceFile app.WorkplaceFile
	err = layer.UnmarshalJSON(data, &workplaceFile)
	if err != nil {
		fmt.Println("Error Unmarshal JSON", path, err)
		panic(1)
	}

	app.Apk.CurrentFile = workplaceFile
	app.Apk.Layers = workplaceFile.Layers

	for _, l := range workplaceFile.Layers {
		image, isImage := l.(*layer.Image)
		if !isImage {
			continue
		}

		app.Apk.Workplace.LoadImagePath(image.Path)
	}

	if len(workplaceFile.Layers) > 0 {
		app.Apk.SelectedLayer = workplaceFile.Layers[0]
	}
}
