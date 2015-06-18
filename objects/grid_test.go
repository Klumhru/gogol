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
	g.Cells[0] = true
	g.Cells[1] = true
	g.Cells[2] = true
	b := MakeGrid(5, 5)
	g.Project(&b)
	assert.Equal(t, true, g.Get(1, 1), "Cell 1x1 should be alive")
}
