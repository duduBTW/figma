package app

import (
	"encoding/json"
	"errors"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const roundedFrag = `
#version 330

in vec2 fragTexCoord;
in vec4 fragColor;

uniform sampler2D texture0;
uniform vec4 colDiffuse;

uniform vec2 uTexSize;
uniform float uRadius;

out vec4 finalColor;

float roundedRectAlpha(vec2 p, vec2 size, float r) {
    vec2 halfSize = size * 0.5;
    vec2 d = abs(p - halfSize) - (halfSize - vec2(r));
    float outside = length(max(d, 0.0)) - r;
    float aa = 1.0 - smoothstep(-1.0, 1.0, outside);
    return clamp(aa, 0.0, 1.0);
}

void main() {
    vec4 tex = texture(texture0, fragTexCoord) * colDiffuse * fragColor;
    vec2 p = fragTexCoord * uTexSize;
    float a = roundedRectAlpha(p, uTexSize, uRadius);
    finalColor = vec4(tex.rgb, tex.a * a);
}
`

type CreateWorkplace struct {
	MiniatureTexture   *rl.Texture2D
	RoundedImageShader rl.Shader
	FormData           WorkplaceFile
}

func NewCreateWorkplace() CreateWorkplace {
	return CreateWorkplace{
		MiniatureTexture:   nil,
		RoundedImageShader: rl.LoadShaderFromMemory("", roundedFrag),
		FormData: WorkplaceFile{
			Title:            "My Project",
			Framerate:        24,
			Duration:         "01:00",
			ResolutionWidth:  1920,
			ResolutionHeight: 1080,
			Layers:           []Layer{},
		},
	}
}

func NewWorkplaceLoad() {
	Apk.CreateWorkplace = NewCreateWorkplace()
}

func (state *CreateWorkplace) Unload() {
	rl.UnloadTexture(*state.MiniatureTexture)
}

func (newWorkplaceData *CreateWorkplace) Submit() error {
	fileName := newWorkplaceData.FormData.Title + ".json"
	filePath := "projects/" + fileName
	if _, err := os.Stat(filePath); err == nil {
		return errors.New("File already exists")
	}

	jsonFileData, err := json.Marshal(newWorkplaceData.FormData)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jsonFileData, 0644)
}
