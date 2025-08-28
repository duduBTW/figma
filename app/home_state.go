package app

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Home struct {
	Files             []HomeWorkplaceFile
	MinuatureTextures [](*rl.Texture2D)
	SelectedFile      HomeWorkplaceFile
}

func NewHome() Home {
	files, err := ReadFiles()
	if err != nil {
		fmt.Println("Failed to read files!", err)
		panic(1)
	}

	return Home{
		Files:             files,
		MinuatureTextures: LoadMiniatures(files),
	}
}

func (state *Home) Unload() {
	for _, texture := range state.MinuatureTextures {
		if texture == nil {
			continue
		}
		rl.UnloadTexture(*texture)
	}
}

func ReadFiles() ([]HomeWorkplaceFile, error) {
	workplaceFiles := []HomeWorkplaceFile{}

	files, err := os.ReadDir(PROJECTS_FOLDER)
	if err != nil {
		return workplaceFiles, err
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		path := filepath.Join(PROJECTS_FOLDER, file.Name())

		// Read file
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading file:", path, err)
			continue
		}

		// Decode JSON
		var workplaceFile HomeWorkplaceFile
		if err := json.Unmarshal(data, &workplaceFile); err != nil {
			return workplaceFiles, err
		}

		workplaceFiles = append(workplaceFiles, workplaceFile)
	}

	return workplaceFiles, nil
}

func LoadMiniatures(workplaceFiles []HomeWorkplaceFile) []*rl.Texture2D {
	miniatureTextures := make([](*rl.Texture2D), len(workplaceFiles))
	for index, file := range workplaceFiles {
		miniaturePath := file.PreviewImagePath
		if miniaturePath == "" {
			miniatureTextures[index] = nil
			continue
		}

		texture := rl.LoadTexture(miniaturePath)
		rl.GenTextureMipmaps(&texture)
		rl.SetTextureFilter(texture, rl.FilterTrilinear)
		miniatureTextures[index] = &texture
	}

	return miniatureTextures
}

func HomeLoad() {
	Apk.Home = NewHome()
}
