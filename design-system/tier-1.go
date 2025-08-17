package designSystem

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	DS_COLOR_BRAND DsColor = rl.NewColor(185, 39, 88, 255)

	DS_COLOR_GRAY_50  DsColor = rl.NewColor(255, 255, 255, 255)
	DS_COLOR_GRAY_100 DsColor = rl.NewColor(209, 209, 209, 255)
	DS_COLOR_GRAY_200 DsColor = rl.NewColor(184, 184, 184, 255)
	DS_COLOR_GRAY_300 DsColor = rl.NewColor(158, 158, 158, 255)
	DS_COLOR_GRAY_400 DsColor = rl.NewColor(41, 41, 41, 255)
	DS_COLOR_GRAY_500 DsColor = rl.NewColor(34, 34, 34, 255)
	DS_COLOR_GRAY_600 DsColor = rl.NewColor(26, 26, 26, 255)
	DS_COLOR_GRAY_950 DsColor = rl.NewColor(0, 0, 0, 255)

	DS_COLOR_BLUE_400 DsColor = rl.NewColor(101, 208, 240, 255)
)

const (
	DS_SPACING_1 DsSpacing = 4
	DS_SPACING_2 DsSpacing = 8
	DS_SPACING_3 DsSpacing = 12
	DS_SPACING_4 DsSpacing = 16
	DS_SPACING_5 DsSpacing = 24
)

const (
	DS_RADII_SM DsRadii = 4
	DS_RADII    DsRadii = 6
)

const (
	DS_FONT_SIZE_SM DsFontSize = 10
	DS_FONT_SIZE    DsFontSize = 14
	DS_FONT_SIZE_LG DsFontSize = 16
)
