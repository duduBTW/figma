package main

import (
	"C"
	"image"
	"image/draw"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

const (
	BOTTOM_PANEL_HEIGHT = 360
	TOOL_DOCK_HEIGHT    = 48
	SIDE_PANEL_WIDTH    = 480
	PANEL_GAP           = 8
	PANEL_ROUNDNESS     = 6
)

var IconTexture rl.Texture2D

func DrawRectangleRoundedPixels(rec rl.Rectangle, radiusPixels float32, color rl.Color) {
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
	rl.DrawRectangleRounded(rec, roundness, 0, color)
}

func imageToRaylibImage(img *image.RGBA) rl.Image {
	size := len(img.Pix)
	data := C.malloc(C.size_t(size))
	copy((*[1 << 30]byte)(data)[0:size], img.Pix) // Copy pixel data into C heap

	return rl.Image{
		Data:    data,
		Width:   int32(img.Rect.Dx()),
		Height:  int32(img.Rect.Dy()),
		Mipmaps: 1,
		Format:  rl.UncompressedR8g8b8a8,
	}
}

func LoadSVGAsTexture(filePath string, width, height int) rl.Texture2D {
	// Load and parse the SVG
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	icon, err := oksvg.ReadIconStream(file)
	if err != nil {
		panic(err)
	}

	icon.SetTarget(0, 0, float64(width), float64(height))

	// Create RGBA image buffer
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.Transparent, image.Point{}, draw.Src)

	// Rasterize SVG to image buffer
	scanner := rasterx.NewScannerGV(width, height, img, img.Bounds())
	raster := rasterx.NewDasher(width, height, scanner)
	icon.Draw(raster, 1.0)

	rayImg := imageToRaylibImage(img)

	// Load texture from image
	texture := rl.LoadTextureFromImage(&rayImg)

	// Cleanup
	rl.UnloadImage(&rayImg)

	return texture
}
