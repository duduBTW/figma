package components

import (
	"fmt"

	"github.com/dudubtw/figma/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type InputInstance struct {
	Draw         func()
	Value        string
	Rect         rl.Rectangle
	IsBluring    bool
	IsFocusing   bool
	HasSubmitted bool
	State        InteractableState
}

type InputProps struct {
	Id          string
	X           float32
	Y           float32
	Width       float32
	Placeholder string
	Value       string
	Ui          *lib.UIStruct
	MousePoint  rl.Vector2
	// LeftIndicator rune
}

const INPUT_HEIGHT float32 = 24
const INDICATOR_FONT_SIZE int32 = 14

func (c *Components) Input(props InputProps) InputInstance {
	c.ui.TabOrder = append(c.ui.TabOrder, props.Id)
	inputInstance := InputInstance{}
	// leftIndicator := string(props.LeftIndicator)
	rect := rl.NewRectangle(props.X, props.Y, props.Width, INPUT_HEIGHT)

	rectInt32 := rect.ToInt32()
	InputEvent(rect, props)

	state := InputState(props)

	// on blur
	inputInstance.IsBluring = state == STATE_INITIAL && c.inputStates[props.Id] == STATE_ACTIVE

	// on focus
	inputInstance.IsFocusing = state == STATE_ACTIVE && c.inputStates[props.Id] == STATE_INITIAL

	// on submit
	inputInstance.HasSubmitted = state == STATE_ACTIVE && rl.IsKeyPressed(rl.KeyEnter)

	// update global state
	c.inputStates[props.Id] = state

	isEmpty := props.Value == ""
	var fontSize int32 = 14
	var textY int32 = rectInt32.Y + (int32(rect.Height)-fontSize)/2 + 1
	var textX int32 = rectInt32.X + 8
	// originalTextX := textX

	// if leftIndicator != "" {
	// 	var indicatorPadding int32 = 12
	// 	textX += rl.MeasureText(leftIndicator, INDICATOR_FONT_SIZE) + indicatorPadding
	// }

	newValue := InputValueChange(props, state)

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rect) && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		UpdateClickedCursorPosition(newValue, textX, fontSize, props)
	}

	inputInstance.Draw = func() {
		DrawRectangleRoundedPixels(rect, 4, rl.NewColor(41, 41, 41, 255))

		if isEmpty && state != STATE_ACTIVE {
			rl.DrawText(props.Placeholder, textX, textY, fontSize, rl.Fade(rl.White, 0.42))
		} else if !isEmpty {
			rl.DrawText(props.Value, textX, textY, fontSize, rl.White)
		}

		// if leftIndicator != "" {
		// 	rl.DrawText(leftIndicator, originalTextX, textY, fontSize, rl.White)
		// }

		if state == STATE_ACTIVE {
			DrawRectangleRoundedLinePixels(rect, 4, 1, rl.Fade(rl.White, 0.4))
		}

		if state == STATE_ACTIVE && props.Ui.InputCursorStart == props.Ui.InputCursorEnd {
			DrawCusor(props.Ui.InputCursorStart, newValue, textX, textY, fontSize, props.Ui)
		}
	}

	inputInstance.Rect = rect
	inputInstance.Value = newValue
	inputInstance.State = state
	return inputInstance
}

func UpdateClickedCursorPosition(value string, textX, fontSize int32, props InputProps) {
	mousePoint := props.MousePoint
	ui := props.Ui
	totalTextWidth := rl.MeasureText(value, fontSize)
	if mousePoint.X >= float32(textX+totalTextWidth) {
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
		charSize := rl.MeasureText(char, fontSize)

		if mousePoint.X >= float32(lastPos) && mousePoint.X <= float32(lastPos+charSize) {
			if mousePoint.X > float32(lastPos+(charSize/2)) {
				ui.SetCursors(index + 1)
			} else {
				ui.SetCursors(index)
			}
			return
		}
		//                    letter spacing
		lastPos += charSize + 1
		index++
	}

	fmt.Println("uh oh")
}

func DrawCusor(position int, value string, textX, textY, fontSize int32, ui *lib.UIStruct) {
	color := rl.Fade(rl.White, 0.72)
	if ShouldBlink(ui) {
		color = rl.White
	}

	x := textX + rl.MeasureText(value[:position], fontSize) + 1
	y := textY + 7
	cursorHeight := fontSize - 5
	rl.DrawLine(x, y-cursorHeight, x, y+cursorHeight, color)
}

const blinkInterval = 0.6
const blinkTotal = 0.5

func ShouldBlink(ui *lib.UIStruct) bool {
	// is blinking
	if ui.BlinkingTimer > 0 {
		return ShouldStayBlinked(ui)
	}

	ui.BlinkTimer += rl.GetFrameTime()
	if ui.BlinkTimer > blinkInterval {
		ui.BlinkTimer = 0
		ui.BlinkingTimer += 0.001
		return true
	}
	return false
}
func ShouldStayBlinked(ui *lib.UIStruct) bool {
	ui.BlinkingTimer += rl.GetFrameTime()
	if ui.BlinkingTimer > blinkInterval {
		ui.BlinkingTimer = 0
		return false
	}
	return true
}

func InputValueChange(props InputProps, state InteractableState) string {
	value := props.Value
	ui := props.Ui
	if state != STATE_ACTIVE {
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
	ui := props.Ui
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
		rl.SetMouseCursor(rl.MouseCursorIBeam)
	}

	return false
}

func InputState(props InputProps) InteractableState {
	ui := props.Ui
	switch props.Id {
	case ui.FocusedId:
		return STATE_ACTIVE
	case ui.HotId:
		return STATE_HOT
	default:
		return STATE_INITIAL
	}
}

func (input *InputInstance) Blur(ui *lib.UIStruct) {
	ui.FocusedId = ""
}
