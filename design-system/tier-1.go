package ds

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	COLOR_BRAND       DsColor = rl.NewColor(185, 39, 88, 255)
	COLOR_BRAND_DARK  DsColor = rl.NewColor(177, 25, 68, 255)
	COLOR_BRAND_LIGHT DsColor = rl.NewColor(204, 52, 113, 255)

	COLOR_GRAY_50  DsColor = rl.NewColor(255, 255, 255, 255)
	COLOR_GRAY_100 DsColor = rl.NewColor(209, 209, 209, 255)
	COLOR_GRAY_200 DsColor = rl.NewColor(184, 184, 184, 255)
	COLOR_GRAY_300 DsColor = rl.NewColor(158, 158, 158, 255)
	COLOR_GRAY_400 DsColor = rl.NewColor(41, 41, 41, 255)
	COLOR_GRAY_500 DsColor = rl.NewColor(34, 34, 34, 255)
	COLOR_GRAY_600 DsColor = rl.NewColor(26, 26, 26, 255)
	COLOR_GRAY_950 DsColor = rl.NewColor(0, 0, 0, 255)

	COLOR_BLUE_400 DsColor = rl.NewColor(101, 208, 240, 255)
)

const (
	SPACING_1 DsSpacing = 4
	SPACING_2 DsSpacing = 8
	SPACING_3 DsSpacing = 12
	SPACING_4 DsSpacing = 16
	SPACING_5 DsSpacing = 24
)

const (
	RADII_SM DsRadii = 4
	RADII    DsRadii = 6
)

const (
	FONT_SIZE_SM DsFontSize = 10
	FONT_SIZE    DsFontSize = 14
	FONT_SIZE_LG DsFontSize = 16

	FONT_WEIGHT_BOLD    DsFontWeight = "Bold"
	FONT_WEIGHT_REGULAR DsFontWeight = "Regular"
	FONT_WEIGHT_MEDIUM  DsFontWeight = "Medium"
)
