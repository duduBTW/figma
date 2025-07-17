package app

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Interactable struct {
	id string
}

type InteractableState = int8

const (
	STATE_INITIAL InteractableState = 0
	STATE_HOT     InteractableState = 1
	STATE_ACTIVE  InteractableState = 2
)

func (interactable *Interactable) State() InteractableState {
	switch interactable.id {
	case Apk.ActiveId:
		return STATE_ACTIVE
	case Apk.HotId:
		return STATE_HOT
	default:
		return STATE_INITIAL
	}
}

func (interactable *Interactable) Event(mousePoint rl.Vector2, rect rl.Rectangle) bool {
	id := interactable.id
	isActive := id == Apk.ActiveId
	isInside := rl.CheckCollisionPointRec(mousePoint, rect)

	if Apk.HotId == id && !isInside {
		Apk.HotId = ""
		return false
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && isInside {
		Apk.ActiveId = id
		Apk.HotId = ""
		return false
	}

	// Other element is being interacted with
	if (Apk.ActiveId != "" && !isActive) ||
		(Apk.HotId != "" && !isInside) {
		return false
	}

	if isActive && rl.IsMouseButtonUp(rl.MouseButtonLeft) {
		Apk.ActiveId = ""
		Apk.HotId = ""
		return isInside
	}

	if isActive || !isInside {
		return false
	}

	if Apk.HotId == "" {
		Apk.HotId = id
	}

	return false
}

func NewInteractable(id string) Interactable {
	return Interactable{
		id: id,
	}
}
