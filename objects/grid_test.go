package objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMake(t *testing.T) {
	g := MakeGrid(5, 5)
	assert.Equal(t, 25, len(g.Cells), "There should be 5*5=25 grid cells")
}

func TestRandomize(t *testing.T) {
	g := MakeGrid(5, 5)
	g.Randomize(0.5)
	assert.NotEqual(t, 0, g.CountLiveCells(), "There should be some live cells")
	g.Randomize(1)
	assert.Equal(t, 25, g.CountLiveCells(), "All cells should be alive")
}

func TestSet(t *testing.T) {
	g := MakeGrid(5, 5)
	g.Set(0, 0, true)
	assert.Equal(t, true, g.Get(0, 0), "Cell 0x0 should be alive")
}

func TestGet(t *testing.T) {
	g := MakeGrid(5, 5)
	g.Cells[0] = true
	assert.Equal(t, true, g.Get(0, 0), "Cell 0x0 should be alive")
}

func TestProject(t *testing.T) {
	g := MakeGrid(5, 5)
	g.Set(0, 0, true)
	g.Set(1, 0, true)
	g.Set(2, 0, true)
	b := MakeGrid(5, 5)
	g.Project(&b)
	assert.Equal(t, true, b.Get(1, 1), "Cell 1x1 should be alive")
}

type OutStream struct{}

var bytesWritten int

func (outWrite OutStream) Write(p []byte) (n int, err error) {
	n = len(p)
	bytesWritten += n
	return len(p), nil
}

func TestRender(t *testing.T) {
	g := MakeGrid(5, 5)
	bytesWritten = 0
	out := OutStream{}
	g.Render(out)
	assert.Equal(t, 30, bytesWritten, "Incorrect write length")
}

func TestLiveCells(t *testing.T) {
	g := MakeGrid(5, 5)
	assert.Equal(t, 0, g.CountLiveCells(),
		"Grids should start with no live cells")
	g.Set(0, 0, true)
	assert.Equal(t, 1, g.CountLiveCells(),
		"There should be only one live cell")
}

func TestIndex(t *testing.T) {
	g := MakeGrid(5, 5)
	assert.Equal(t, 0, g.Index(0, 0), "Index of 0x0 should be 0")
	assert.Equal(t, 4, g.Index(4, 0), "Index of 4x0 should be 4")
	assert.Equal(t, 12, g.Index(2, 2), "Index of 2x2 should be 12")
}

func TestCoords(t *testing.T) {
	g := MakeGrid(5, 5)
	x, y := g.GetCoords(13)
	assert.Equal(t, 3, x, "Index 13 should have x coords 3")
	assert.Equal(t, 2, y, "Index 13 should have y coords 2")
	x, y = g.GetCoords(21)
	assert.Equal(t, 1, x, "Index 21 should have x coords 1")
	assert.Equal(t, 4, y, "Index 21 should have y coords 4")
}

func TestLiveNeighbours(t *testing.T) {
	g := MakeGrid(5, 5)
	g.Set(0, 0, true)
	g.Set(1, 0, true)
	g.Set(2, 0, true)
	assert.Equal(t, 1, g.AliveNeighbours(0, 0),
		"0x0 should have one living neighbour")
	assert.Equal(t, 2, g.AliveNeighbours(1, 0),
		"1x0 should have two living neighbours")
	assert.Equal(t, 1, g.AliveNeighbours(3, 0),
		"2x0 should have one living neighbour")
	assert.Equal(t, 3, g.AliveNeighbours(1, 1),
		"1x1 should have three living neighbours")
}
