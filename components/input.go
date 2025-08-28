package components

import (
	"fmt"

	"github.com/dudubtw/figma/app"
	ds "github.com/dudubtw/figma/design-system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type InputInstance struct {
	Draw         func()
	Value        string
	Rect         rl.Rectangle
	IsBluring    bool
	IsFocusing   bool
	HasSubmitted bool
	State        app.InteractableState
}

type InputProps struct {
	Id          string
	X           float32
	Y           float32
	Width       float32
	Placeholder string
	Value       string
	MousePoint  rl.Vector2
	// LeftIndicator rune
}

const INPUT_HEIGHT float32 = 24
const INDICATOR_FONT_SIZE int32 = 14

func Input(props InputProps) InputInstance {
	if app.Apk.InputNames[props.Id] {
		// fmt.Println("Input with the same id declared: ", props.Id)
		// panic(1)
	}

	app.Apk.InputNames[props.Id] = true
	app.Apk.TabOrder = append(app.Apk.TabOrder, props.Id)
	inputInstance := InputInstance{}
	// leftIndicator := string(props.LeftIndicator)
	rect := rl.NewRectangle(props.X, props.Y, props.Width, INPUT_HEIGHT)

	rectInt32 := rect.ToInt32()
	InputEvent(rect, props)

	state := InputState(props)

	// on blur
	inputInstance.IsBluring = state == app.STATE_INITIAL && app.Apk.InputStates[props.Id] == app.STATE_ACTIVE

	// on focus
	inputInstance.IsFocusing = state == app.STATE_ACTIVE && app.Apk.InputStates[props.Id] == app.STATE_INITIAL

	// on submit
	inputInstance.HasSubmitted = state == app.STATE_ACTIVE && rl.IsKeyPressed(rl.KeyEnter)

	// update global state
	app.Apk.InputStates[props.Id] = state

	isEmpty := props.Value == ""
	var fontSize float32 = float32(ds.FONT_SIZE)
	font := app.Apk.GetFont(ds.FONT_SIZE, ds.FONT_WEIGHT_MEDIUM)
	var textY int32 = rectInt32.Y + (int32(rect.Height)-int32(fontSize))/2 + 1
	var textX int32 = rectInt32.X + 8
	// originalTextX := textX

	// if leftIndicator != "" {
	// 	var indicatorPadding int32 = 12
	// 	textX += rl.MeasureText(leftIndicator, INDICATOR_FONT_SIZE) + indicatorPadding
	// }

	newValue := InputValueChange(props, state)

	if app.Apk.CanInteract && rl.CheckCollisionPointRec(rl.GetMousePosition(), rect) && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		UpdateClickedCursorPosition(newValue, textX, fontSize, font, props)
	}

	inputInstance.Draw = func() {
		DrawRectangleRoundedPixels(rect, 4, rl.NewColor(41, 41, 41, 255))
		textVec := rl.NewVector2(float32(textX), float32(textY))

		if isEmpty && state != app.STATE_ACTIVE {
			rl.DrawTextEx(font, props.Placeholder, textVec, float32(ds.FONT_SIZE), 1, ds.T2_COLOR_CONTENT_ACCENT)
		} else if !isEmpty {
			rl.DrawTextEx(font, props.Value, textVec, float32(ds.FONT_SIZE), 1, ds.T2_COLOR_CONTENT)
		}

		// if leftIndicator != "" {
		// 	rl.DrawText(leftIndicator, originalTextX, textY, fontSize, rl.White)
		// }

		if state == app.STATE_ACTIVE {
			DrawRectangleRoundedLinePixels(rect, 5, 2, ds.T2_COLOR_HIGHLIGHT)
		}

		if state == app.STATE_ACTIVE && app.Apk.InputCursorStart == app.Apk.InputCursorEnd {
			DrawCusor(app.Apk.InputCursorStart, newValue, textX, textY, fontSize, font)
		}
	}

	inputInstance.Rect = rect
	inputInstance.Value = newValue
	inputInstance.State = state
	return inputInstance
}

func UpdateClickedCursorPosition(value string, textX int32, fontSize float32, font rl.Font, props InputProps) {
	mousePoint := props.MousePoint
	ui := &app.Apk
	totalTextSize := rl.MeasureTextEx(font, value, fontSize, 1)
	if mousePoint.X >= float32(textX)+totalTextSize.X {
		ui.SetCursors(len(value))
		return
	}

	if mousePoint.X <= float32(textX) {
		ui.SetCursors(0)
		return
	}

	var lastPos int32 = textX
	index := 0
	for index <= len(value)-1 {
		char := value[index : index+1]
		charSize := rl.MeasureTextEx(font, char, fontSize, 1)

		if mousePoint.X >= float32(lastPos) && mousePoint.X <= float32(lastPos+int32(charSize.X)) {
			if mousePoint.X > float32(lastPos)+(charSize.X/2) {
				ui.SetCursors(index + 1)
			} else {
				ui.SetCursors(index)
			}
			return
		}
		//                    letter spacing
		lastPos += int32(charSize.X) + 1
		index++
	}

	fmt.Println("uh oh")
}

func DrawCusor(position int, value string, textX, textY int32, fontSize float32, font rl.Font) {
	color := rl.Fade(rl.White, 0.72)
	if ShouldBlink() {
		color = rl.White
	}

	x := textX + int32(rl.MeasureTextEx(font, value[:position], fontSize, 1).X) + 1
	y := textY + 7
	cursorHeight := int32(fontSize) - 5
	rl.DrawLine(x, y-cursorHeight, x, y+cursorHeight, color)
}

const blinkInterval = 0.6
const blinkTotal = 0.5

func ShouldBlink() bool {
	// is blinking
	if app.Apk.BlinkingTimer > 0 {
		return ShouldStayBlinked()
	}

	app.Apk.BlinkTimer += rl.GetFrameTime()
	if app.Apk.BlinkTimer > blinkInterval {
		app.Apk.BlinkTimer = 0
		app.Apk.BlinkingTimer += 0.001
		return true
	}
	return false
}
func ShouldStayBlinked() bool {
	app.Apk.BlinkingTimer += rl.GetFrameTime()
	if app.Apk.BlinkingTimer > blinkInterval {
		app.Apk.BlinkingTimer = 0
		return false
	}
	return true
}

func InputValueChange(props InputProps, state app.InteractableState) string {
	value := props.Value
	ui := &app.Apk
	if state != app.STATE_ACTIVE {
		return value
	}

	key := rl.GetCharPressed()
	for key > 0 {
		if len(value) == 0 {
			value += string(key)
		} else {
			value = string(value[:ui.InputCursorStart]) + string(key) + string(value[ui.InputCursorStart:])
		}
		key = rl.GetCharPressed()
		ui.IncrementCursor()
	}

	if !app.Apk.CanInteract {
		return value
	}

	if (rl.IsKeyPressed(rl.KeyBackspace) || rl.IsKeyPressedRepeat(rl.KeyBackspace)) && ui.InputCursorStart > 0 {
		value = string(value[:ui.InputCursorStart-1]) + string(value[ui.InputCursorStart:])
		ui.DecrementCursor()
	}

	if (rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft)) && ui.InputCursorStart > 0 {
		if rl.IsKeyDown(rl.KeyLeftControl) {
			ui.SetCursors(0)
		} else {
			ui.DecrementCursor()
		}
	}

	if (rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressedRepeat(rl.KeyRight)) && ui.InputCursorStart < len(props.Value) {
		if rl.IsKeyDown(rl.KeyLeftControl) {
			ui.SetCursors(len(props.Value))
		} else {
			ui.IncrementCursor()
		}
	}

	return value
}

func InputEvent(rect rl.Rectangle, props InputProps) bool {
	ui := &app.Apk
	id := props.Id
	mousePoint := props.MousePoint
	isFocused := id == ui.FocusedId
	isInside := rl.CheckCollisionPointRec(mousePoint, rect)

	if isFocused && rl.IsMouseButtonDown(rl.MouseButtonLeft) && !isInside {
		ui.FocusedId = ""
		return false
	}

	if isFocused {
		if isInside {
		}

		return false
	}

	if ui.HotId == id && !isInside {
		ui.HotId = ""
		return false
	}

	// Other element is being interacted with
	if (ui.ActiveId != "" && !isFocused) ||
		(ui.HotId != "" && !isInside) {
		return false
	}

	// clicked
	if isInside && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		ui.FocusedId = id
		ui.HotId = ""
		return true
	}

	if !isFocused && isInside {
		ui.HotId = id
	}

	return false
}

func InputState(props InputProps) app.InteractableState {
	ui := &app.Apk
	switch props.Id {
	case ui.FocusedId:
		return app.STATE_ACTIVE
	case ui.HotId:
		return app.STATE_HOT
	default:
		return app.STATE_INITIAL
	}
}

func (input *InputInstance) Blur() {
	app.Apk.FocusedId = ""
}
