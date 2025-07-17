package app

type App struct {
	State
	Components
}

var Apk = App{
	State:      NewState(),
	Components: NewComponents(),
}

func (app *App) Frame() {
	app.Components.InputNames = map[string]bool{}

	if !app.State.IsPlaying {
		return
	}

	// Loop
	if app.State.VisibleFrames[1] == app.State.SelectedFrame {
		app.State.SelectedFrame = app.State.VisibleFrames[0]
		return
	}

	// Next
	app.State.SelectedFrame++
}
