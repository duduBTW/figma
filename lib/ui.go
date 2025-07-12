package lib

import rl "github.com/gen2brain/raylib-go/raylib"

type Tool uint8

const (
	ToolSelection Tool = 0
	ToolRectangle Tool = 1
	ToolText      Tool = 2
)

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

	SelectedFrame int

	IsPlaying bool

	SelectedTool Tool

	TabOrder []string

	FrameWidth float32
}

func NewUi() UIStruct {
	return UIStruct{
		SelectedTool: ToolSelection,
		TabOrder:     []string{},
	}
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

func (ui *UIStruct) ResetTabOrder() {
	ui.TabOrder = []string{}
}

func (ui *UIStruct) TogglePlay() {
	ui.IsPlaying = !ui.IsPlaying
}

func (ui *UIStruct) GetXTimelineFrame(timelineRect rl.Rectangle, frame float32) float32 {
	if frame == 0 {
		frame = float32(ui.SelectedFrame)
	}

	return timelineRect.X + ui.FrameWidth*frame + 1
}
