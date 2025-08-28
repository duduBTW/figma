package app

import (
	"math"

	"github.com/dudubtw/figma/fmath"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	position rl.Vector2
	columns  int
	width    float32
	gap      float32

	currentIndex int

	drawStack []func()

	columnWidth float32
	rowHeight   []float32

	// computedPositions []float32
}

// return draw, height
type GridComponent func(x, y, width float32) (func(), float32)

func NewGrid(columns int, width, gap float32, position rl.Vector2) *Grid {
	gapColumnWidth := (float32(columns) - 1) * gap
	columnWidth := (width - gapColumnWidth) / float32(columns)

	return &Grid{
		columns:      columns,
		width:        width,
		gap:          gap,
		columnWidth:  columnWidth,
		currentIndex: 0,
		position:     position,
		drawStack:    []func(){},
		rowHeight:    []float32{},
	}
}

// 0 index
func (grid *Grid) CurrentColumn() int {
	return (grid.currentIndex % grid.columns)
}

// 0 index
func (grid *Grid) CurrentRow() int {
	// return (grid.columns - (grid.CurrentColumn())) / grid.columns
	return int(math.Ceil(float64(grid.currentIndex+1)/float64(grid.columns))) - 1
}

func (grid *Grid) X() float32 {
	gap := grid.gap * float32(grid.CurrentColumn())
	return grid.position.X + (grid.columnWidth * float32(grid.CurrentColumn())) + gap
}

func (grid *Grid) previusRowMaxHeight() float32 {
	currentRow := grid.CurrentRow()
	if currentRow == 0 {
		return 0
	}

	var height float32 = 0
	for i := 0; i < currentRow; i++ {
		height += grid.rowHeight[i]
	}

	return height
}

func (grid *Grid) Y() float32 {
	gap := grid.gap * float32(grid.CurrentRow())
	return grid.position.Y + grid.previusRowMaxHeight() + gap
}

func (grid *Grid) UpdateCurrentRowHeight(height float32) {
	currentRow := grid.CurrentRow()
	if len(grid.rowHeight) <= currentRow {
		grid.rowHeight = append(grid.rowHeight, 0)
	}

	grid.rowHeight[currentRow] = fmath.Max(grid.rowHeight[currentRow], height)
}

func (grid *Grid) Add(components ...GridComponent) *Grid {
	for _, component := range components {
		draw, height := component(grid.X(), grid.Y(), grid.columnWidth)
		grid.UpdateCurrentRowHeight(height)
		grid.drawStack = append(grid.drawStack, draw)
		grid.currentIndex++
	}

	return grid
}

func (grid *Grid) Draw() {
	for _, draw := range grid.drawStack {
		draw()
	}
}

func (grid *Grid) Height() float32 {
	return grid.rowHeight[len(grid.rowHeight)-1]
}
