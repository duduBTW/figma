package components

import (
	"github.com/dudubtw/figma/app"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Helper function to find the minimum of two float32 values
func minF(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

// Helper function to find the maximum of two float32 values
func maxF(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

type Roundness = float32

const (
	ROUNDED = 8
)

func DrawRectangleRoundedPixels(rec rl.Rectangle, radiusPixels Roundness, color rl.Color) {
	// Ensure radius is not negative
	radiusPixels = maxF(0, radiusPixels)

	// Prevent division by zero or weirdness with non-positive dimensions
	if rec.Width <= 0 || rec.Height <= 0 {
		// Optionally draw a regular rectangle or just return
		// rl.DrawRectangleRec(rec, color)
		return
	}

	// Find the smaller dimension
	minDimension := minF(rec.Width, rec.Height)

	// Calculate roundness based on pixel radius
	// roundness = (radius * 2) / minDimension
	roundness := (radiusPixels * 2) / minDimension

	// Clamp roundness to the valid range [0.0, 1.0]
	// If requested radiusPixels * 2 > minDimension, it means the radius
	// is too large, so we cap at full roundness (1.0).
	roundness = maxF(0.0, minF(roundness, 1.0))

	// Call the original raylib function with the calculated roundness
	rl.DrawRectangleRounded(rec, roundness, 16, color)
}

func DrawRectangleRoundedLinePixels(rec rl.Rectangle, radiusPixels Roundness, lineThick float32, color rl.Color) {
	// Ensure radius is not negative
	radiusPixels = maxF(0, radiusPixels)

	// Prevent division by zero or weirdness with non-positive dimensions
	if rec.Width <= 0 || rec.Height <= 0 {
		// Optionally draw a regular rectangle or just return
		// rl.DrawRectangleRec(rec, color)
		return
	}

	// Find the smaller dimension
	minDimension := minF(rec.Width, rec.Height)

	// Calculate roundness based on pixel radius
	// roundness = (radius * 2) / minDimension
	roundness := (radiusPixels * 2) / minDimension

	// Clamp roundness to the valid range [0.0, 1.0]
	// If requested radiusPixels * 2 > minDimension, it means the radius
	// is too large, so we cap at full roundness (1.0).
	roundness = maxF(0.0, minF(roundness, 1.0))

	// Call the original raylib function with the calculated roundness
	rl.DrawRectangleRoundedLinesEx(rec, roundness, 0, lineThick, color)
}

var renderTextureCache = map[[3]float32]rl.RenderTexture2D{}

func DrawImageCoverRounded(tex *rl.Texture2D, dest rl.Rectangle, radius float32) {
	key := [3]float32{float32(tex.ID), dest.Width, dest.Height}
	rt, ok := renderTextureCache[key]
	if !ok {
		// for key, oldRt := range renderTextureCache {
		// 	if key[0] != float32(tex.ID) {
		// 		continue
		// 	}

		// 	rl.UnloadRenderTexture(oldRt)
		// }

		rl.UnloadRenderTexture(rt)
		renderTextureCache[key] = rl.LoadRenderTexture(int32(dest.Width), int32(dest.Height))
		rl.SetTextureFilter(renderTextureCache[key].Texture, rl.FilterTrilinear)
	}

	sh := app.Apk.RoundedImageShader
	uTexSizeLoc := rl.GetShaderLocation(sh, "uTexSize")
	uRadiusLoc := rl.GetShaderLocation(sh, "uRadius")

	imgW := float32(tex.Width)
	imgH := float32(tex.Height)
	scale := max(dest.Width/imgW, dest.Height/imgH)
	cropW := dest.Width / scale
	cropH := dest.Height / scale
	src := rl.NewRectangle((imgW-cropW)*0.5, (imgH-cropH)*0.5, cropW, cropH)

	rl.BeginTextureMode(rt)
	rl.ClearBackground(rl.Blank)
	rl.DrawTexturePro(*tex, src, rl.NewRectangle(0, 0, float32(rt.Texture.Width), float32(rt.Texture.Height)), rl.NewVector2(0, 0), 0, rl.White)
	rl.EndTextureMode()

	rl.BeginShaderMode(sh)
	rl.SetShaderValue(sh, uTexSizeLoc, []float32{float32(rt.Texture.Width), float32(rt.Texture.Height)}, rl.ShaderUniformVec2)
	rl.SetShaderValue(sh, uRadiusLoc, []float32{radius}, rl.ShaderUniformFloat)
	rtSrc := rl.NewRectangle(0, 0, float32(rt.Texture.Width), -float32(rt.Texture.Height))
	rl.DrawTexturePro(rt.Texture, rtSrc, dest, rl.NewVector2(0, 0), 0, rl.White)
	rl.EndShaderMode()
}
