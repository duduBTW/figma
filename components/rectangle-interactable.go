package components

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectangleInteractable struct {
	id string
}

func (interactable *RectangleInteractable) State() app.InteractableState {
	switch interactable.id {
	case app.Apk.ActiveId:
		return app.STATE_ACTIVE
	case app.Apk.HotId:
		return app.STATE_HOT
	default:
		return app.STATE_INITIAL
	}
}

func (interactable *RectangleInteractable) Event(mousePoint rl.Vector2, rect rl.Rectangle) bool {
	id := interactable.id
	isActive := id == app.Apk.ActiveId
	isInside := rl.CheckCollisionPointRec(mousePoint, rect)

	activeId := app.Apk.ActiveId
	hotId := app.Apk.HotId
	if hotId == id && !isInside {
		hotId = ""
		return false
	}

	// Other element is being interacted with
	if (activeId != "" && !isActive) ||
		(hotId != "" && !isInside) {
		return false
	}

	if isActive && rl.IsMouseButtonUp(rl.MouseButtonLeft) {
		app.Apk.ActiveId = ""
		app.Apk.HotId = ""
		return isInside
	}

	if isActive {
		return false
	}

	if !isInside {
		return false
	}

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		app.Apk.ActiveId = id
		app.Apk.HotId = ""
		return false
	}

	if hotId == "" {
		app.Apk.HotId = id
	}

	return false
}

func NewRectangleInteractable(id string) RectangleInteractable {
	return RectangleInteractable{
		id: id,
	}
}
