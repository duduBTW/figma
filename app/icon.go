package app

import (
	"C"

	"image"
	"image/draw"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

type IconName string
type Icons = map[IconName]*rl.Texture2D

const (
	ICON_MOUSE_POINTER IconName = "mouse-pointer-2"
	ICON_TYPE          IconName = "type"
	ICON_SQUARE        IconName = "square"
	ICON_CHEVRON_DOWN  IconName = "chevron-down"
	ICON_PEN           IconName = "pen-tool"
	ICON_IMAGE         IconName = "image"
	ICON_PLAY          IconName = "play"
	ICON_PAUSE         IconName = "pause"

	ICON_WIDTH  = 16
	ICON_HEIGHT = 16
)

func loadIcon(filePath string) rl.Texture2D {
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

	icon.SetTarget(0, 0, float64(ICON_WIDTH), float64(ICON_HEIGHT))

	// Create RGBA image buffer
	img := image.NewRGBA(image.Rect(0, 0, ICON_WIDTH, ICON_HEIGHT))
	draw.Draw(img, img.Bounds(), image.Transparent, image.Point{}, draw.Src)

	// Rasterize SVG to image buffer
	scanner := rasterx.NewScannerGV(ICON_WIDTH, ICON_HEIGHT, img, img.Bounds())
	raster := rasterx.NewDasher(ICON_WIDTH, ICON_HEIGHT, scanner)
	icon.Draw(raster, 1.0)

	rayImg := imageToRaylibImage(img)

	// Load texture from image
	texture := rl.LoadTextureFromImage(&rayImg)

	// Cleanup
	rl.UnloadImage(&rayImg)

	return texture
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
