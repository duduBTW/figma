package components

import (
	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectangleInteractable struct {
	id string
	ui *lib.UIStruct
}

func (interactable *RectangleInteractable) State() InteractableState {
	ui := interactable.ui
	switch interactable.id {
	case ui.ActiveId:
		return STATE_ACTIVE
	case ui.HotId:
		return STATE_HOT
	default:
		return STATE_INITIAL
	}
}

func (interactable *RectangleInteractable) Event(mousePoint rl.Vector2, rect rl.Rectangle) bool {
	ui := interactable.ui
	id := interactable.id
	isActive := id == ui.ActiveId
	isInside := rl.CheckCollisionPointRec(mousePoint, rect)

	if ui.HotId == id && !isInside {
		ui.HotId = ""
		return false
	}

	// Other element is being interacted with
	if (ui.ActiveId != "" && !isActive) ||
		(ui.HotId != "" && !isInside) {
		return false
	}

	if isActive && rl.IsMouseButtonUp(rl.MouseButtonLeft) {
		ui.ActiveId = ""
		ui.HotId = ""
		return isInside
	}

	if isActive {
		return false
	}

	if !isInside {
		return false
	}

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		ui.ActiveId = id
		ui.HotId = ""
		return false
	}

	if ui.HotId == "" {
		ui.HotId = id
	}

	return false
}

func NewRectangleInteractable(id string, ui *lib.UIStruct) RectangleInteractable {
	return RectangleInteractable{
		id: id,
		ui: ui,
	}
}
