package app

import (
	ds "github.com/dudubtw/figma/design-system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const PROJECTS_FOLDER = "./projects"

type App struct {
	Workplace
	CreateWorkplace
	Home

	SelectedPage Page

	CanInteract bool

	icons Icons

	TabOrder []string

	FocusedId string
	ActiveId  string
	HotId     string

	InputStates      map[string]InteractableState
	InputNames       map[string]bool
	InputCursorStart int
	InputCursorEnd   int
	// Time between blinks
	BlinkTimer float32
	// Time a blink stayed active
	BlinkingTimer float32

	TypographyMap TypographyMap
}

var Apk = App{
	Workplace:    NewWorkplace(),
	CanInteract:  false,
	icons:        Icons{},
	InputStates:  map[string]InteractableState{},
	InputNames:   map[string]bool{},
	TabOrder:     []string{},
	SelectedPage: PAGE_HOME,
	// SelectedPage: PAGE_WORKPLACE,
}

func (app *App) Frame() {
	app.InputNames = map[string]bool{}

	switch app.SelectedPage {
	case PAGE_WORKPLACE:
		app.Workplace.Frame()
	}
}

func (app *App) SetCursors(pos int) {
	app.InputCursorStart = pos
	app.InputCursorEnd = pos
}
func (app *App) IncrementCursor() {
	app.InputCursorStart += 1
	app.InputCursorEnd += 1
}
func (app *App) DecrementCursor() {
	app.InputCursorStart -= 1
	app.InputCursorEnd -= 1
}

func (app *App) Icon(name IconName) rl.Texture2D {
	icon := app.icons[name]

	// already loaded
	if icon != nil {
		return *icon
	}

	loadedIcon := loadIcon("D:\\Peronal\\figma\\assets\\icons\\" + string(name) + ".svg")
	app.icons[name] = &loadedIcon
	return loadedIcon
}

func (app *App) ResetTabOrder() {
	app.TabOrder = []string{}
}

func (app *App) Navigate(to Page) {
	switch app.SelectedPage {
	case PAGE_WORKPLACE:
		Apk.Workplace.Unload()
	case PAGE_HOME:
		Apk.Home.Unload()
	case PAGE_NEW_WORKPLACE:
		Apk.CreateWorkplace.Unload()
	}

	switch to {
	case PAGE_WORKPLACE:
		// TODO - Move the load function from onclick to here
		// WorkplaceLoad()
	case PAGE_HOME:
		HomeLoad()
	case PAGE_NEW_WORKPLACE:
		NewWorkplaceLoad()
	}
	app.SelectedPage = to
}

func (app *App) GetFont(size ds.DsFontSize, weight ds.DsFontWeight) rl.Font {
	return app.TypographyMap[newKey(size, weight)]
}
