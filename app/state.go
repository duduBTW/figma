package app

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SelectedKeyframe struct {
	LayerId  string
	Keyframe [2]float32
}

type State struct {
	icons Icons

	DrawFrameHighlight func()

	SelectedLayer    Layer
	Layers           []Layer
	SelectedKeyframe SelectedKeyframe

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

	FrameWidth     float32
	TimelineScroll float32

	VisibleFrames [2]int

	DroppingTexture *rl.Texture2D
}

func NewState() State {
	return State{
		SelectedTool:  ToolSelection,
		TabOrder:      []string{},
		Layers:        []Layer{},
		VisibleFrames: [2]int{0, 240},
		icons:         Icons{},
	}
}

func (state *State) SetCursors(pos int) {
	state.InputCursorStart = pos
	state.InputCursorEnd = pos
}
func (state *State) IncrementCursor() {
	state.InputCursorStart += 1
	state.InputCursorEnd += 1
}
func (state *State) DecrementCursor() {
	state.InputCursorStart -= 1
	state.InputCursorEnd -= 1
}

func (state *State) ResetTabOrder() {
	state.TabOrder = []string{}
}

func (state *State) TogglePlay() {
	state.IsPlaying = !state.IsPlaying
}

func (state *State) GetXTimelineFrame(timelineRect rl.Rectangle, frame float32) float32 {
	return timelineRect.X + state.FrameWidth*frame + 1
}
func (state *State) ScrollTimeline() {
	var scrollSpeed float32 = 10
	state.TimelineScroll -= rl.GetMouseWheelMove() * scrollSpeed
}

func (state *State) AppendLayer(newLayer Layer) {
	state.Layers = append(state.Layers, newLayer)
	state.SelectedLayer = newLayer
	state.SelectedTool = ToolSelection
}

func (state *State) NewLayerId() string {
	id := 0
	if len(state.Layers) > 0 {
		newId, _ := strconv.Atoi(state.Layers[len(state.Layers)-1].GetElement().Id)
		id = newId
	}
	return strconv.Itoa(id + 1)
}

func (state *State) SetDroppingTexture(source string) {
	texture := rl.LoadTexture(source)
	state.DroppingTexture = &texture
}

func (state *State) Icon(name IconName) rl.Texture2D {
	icon := state.icons[name]

	// already loaded
	if icon != nil {
		return *icon
	}

	loadedIcon := loadIcon("D:\\Peronal\\figma\\assets\\icons\\" + string(name) + ".svg")
	state.icons[name] = &loadedIcon
	return loadedIcon
}
