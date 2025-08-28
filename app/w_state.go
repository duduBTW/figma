package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type HomeWorkplaceFile struct {
	Title            string
	Framerate        int
	Duration         string
	ResolutionWidth  int
	ResolutionHeight int
	PreviewImagePath string
}

type WorkplaceFile struct {
	Title            string
	Framerate        int
	Duration         string
	ResolutionWidth  int
	ResolutionHeight int
	PreviewImagePath string
	Layers           []Layer
}

type SelectedKeyframe struct {
	LayerId  string
	Keyframe float32
}

type DroppingTexture struct {
	Texture rl.Texture2D
	Path    string
}

type Workplace struct {
	Id          string
	CurrentFile WorkplaceFile

	DrawFrameHighlight func()

	SelectedLayer        Layer
	Layers               []Layer
	SelectedKeyframe     SelectedKeyframe
	SelectedAnimatedProp *AnimatedProp

	SelectedFrame int

	IsPlaying bool

	SelectedTool Tool

	FrameWidth            float32
	TimelineScroll        float32
	TimelineCurveSelected bool

	VisibleFrames [2]int

	DroppingImg *DroppingTexture

	FramesRect rl.Rectangle

	imageTextures map[string]rl.Texture2D
}

func NewWorkplace() Workplace {
	return Workplace{
		SelectedTool:          ToolSelection,
		Layers:                []Layer{},
		VisibleFrames:         [2]int{0, 240},
		TimelineCurveSelected: false,
		imageTextures:         map[string]rl.Texture2D{},
	}
}

func (workplace *Workplace) TogglePlay() {
	workplace.IsPlaying = !workplace.IsPlaying
}

func (workplace *Workplace) GetXTimelineFrame(timelineRect rl.Rectangle, frame float32) float32 {
	return timelineRect.X + workplace.FrameWidth*frame + 1
}
func (workplace *Workplace) ScrollTimeline() {
	var scrollSpeed float32 = 10
	workplace.TimelineScroll -= rl.GetMouseWheelMove() * scrollSpeed
}

func (workplace *Workplace) SetSelectedLayer(newLayer Layer) {
	workplace.SelectedLayer = newLayer
	workplace.SelectedAnimatedProp = &newLayer.GetElement().Position.X
}

func (workplace *Workplace) AppendLayer(newLayer Layer) {
	workplace.Layers = append(workplace.Layers, newLayer)
	workplace.SetSelectedLayer(newLayer)
	workplace.SelectedTool = ToolSelection
}

func (workplace *Workplace) NewLayerId() string {
	id := 0
	if len(workplace.Layers) > 0 {
		newId, _ := strconv.Atoi(workplace.Layers[len(workplace.Layers)-1].GetElement().Id)
		id = newId
	}
	return strconv.Itoa(id + 1)
}

func (workplace *Workplace) SetDroppingTexture(source string) {
	workplace.DroppingImg = &DroppingTexture{
		Texture: rl.LoadTexture(source),
		Path:    source,
	}
}

func (workplace *Workplace) ToggleCurveSelected() {
	workplace.TimelineCurveSelected = !workplace.TimelineCurveSelected
}

func (workplace *Workplace) Frame() {
	if !workplace.IsPlaying {
		return
	}

	// Loop
	if workplace.VisibleFrames[1] == workplace.SelectedFrame {
		workplace.SelectedFrame = workplace.VisibleFrames[0]
		return
	}

	// Next
	workplace.SelectedFrame++
}

func (workplace *Workplace) Save() {
	go func() {
		Apk.CurrentFile.Layers = Apk.Layers

		fileName := Apk.CurrentFile.Title + ".json"
		filePath := "projects/" + fileName
		jsonFileData, err := json.Marshal(Apk.CurrentFile)
		if err != nil {
			fmt.Println("Failed saving file", err)
			panic(1)
		}

		os.WriteFile(filePath, jsonFileData, 0644)
	}()
}

func (workplace *Workplace) LoadImagePath(path string) rl.Texture2D {
	texture := rl.LoadTexture(path)
	workplace.imageTextures[path] = texture
	return texture
}

func (workplace *Workplace) GetImagePath(path string) rl.Texture2D {
	return workplace.imageTextures[path]
}

func (workplace *Workplace) Unload() {
	for _, texture := range workplace.imageTextures {
		rl.UnloadTexture(texture)
	}
	*workplace = NewWorkplace()
}
