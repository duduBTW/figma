package lib

type UIStruct struct {
	FocusedId string
	ActiveId  string
	HotId     string

	InputCursorStart int
	InputCursorEnd   int
	// Time between blinks
	BlinkTimer float32
	// Time a blink stayed active
	BlinkingTimer float32
}

func (ui *UIStruct) SetCursors(pos int) {
	ui.InputCursorStart = pos
	ui.InputCursorEnd = pos
}
func (ui *UIStruct) IncrementCursor() {
	ui.InputCursorStart += 1
	ui.InputCursorEnd += 1
}
func (ui *UIStruct) DecrementCursor() {
	ui.InputCursorStart -= 1
	ui.InputCursorEnd -= 1
}
