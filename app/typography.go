package app

import (
	"path/filepath"
	"strconv"

	ds "github.com/dudubtw/figma/design-system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const FONT_NAME = "IBMPlexMono"
const FONT_FORMAT = ".ttf"

type mapKey string
type TypographyMap = map[mapKey]rl.Font

func newKey(size ds.DsFontSize, weight ds.DsFontWeight) mapKey {
	return mapKey(strconv.Itoa(int(size)) + string(weight))
}

func InitTypography() TypographyMap {
	sizes := []ds.DsFontSize{ds.FONT_SIZE_SM, ds.FONT_SIZE, ds.FONT_SIZE_LG}
	weights := []ds.DsFontWeight{ds.FONT_WEIGHT_REGULAR, ds.FONT_WEIGHT_MEDIUM, ds.FONT_WEIGHT_BOLD}

	typographyMap := make(TypographyMap)

	for _, size := range sizes {
		for _, weight := range weights {
			key := newKey(size, weight)

			fileName := filepath.Join("./assets", "fonts", FONT_NAME+"-"+string(weight)+FONT_FORMAT)
			typographyMap[key] = rl.LoadFontEx(fileName, int32(size), nil, 0)
		}
	}

	return typographyMap
}
