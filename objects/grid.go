package objects

import (
	"errors"
	"io"
	"math/rand"
)

// A simple two dimensional grid
type Grid struct {
	Cells  []bool
	Width  int
	Height int
}

// Creates and returns a Grid of specified size
func MakeGrid(width int, height int) Grid {
	g := Grid{
		Cells:  make([]bool, width*height),
		Width:  width,
		Height: height,
	}
	return g
}

// Randomizes a grid's cells to alive/dead by weight
func (g *Grid) Randomize(weight float32) (err error) {
	if weight < 0 || weight > 1 {
		err = errors.New("Weight must be between 0 and 1")
	}
	for i := range g.Cells {
		g.Cells[i] = rand.Float32() < weight
	}
	return err
}

// Checks if cell at specified position is alive
func (g *Grid) Alive(x int, y int) bool {
	return g.Cells[g.Index(x, y)]
}

// Returns the total number of live cells in the grid
func (g *Grid) CountLiveCells() int {
	var ret int
	for _, c := range g.Cells {
		if c {
			ret += 1
		}
	}
	return ret
}

// Gets the life status of a cell
func (g *Grid) Get(x int, y int) bool {
	return g.Cells[g.Index(x, y)]
}

// Sets the life status of a cell
func (g *Grid) Set(x int, y int, state bool) {
	g.Cells[g.Index(x, y)] = state
}

// Projects the status of the grid to a new grid, based on Conways GoL
func (in *Grid) Project(out *Grid) {
}

// Returns the base array index of the given coords
func (g *Grid) Index(x int, y int) int {
	return y*g.Width + x
}

// Renders the Grid to stdout with X meaning alive and empty string meaing dead
func (g *Grid) Render(out io.Writer) {
	var state string
	for _, c := range g.Cells {
		if c {
			state = "X"
		} else {
			state = " "
		}
		out.Write([]byte(state))
	}
}

// Traverses the 8 squares around the coordinate and returns
// how many are living int thouse squares
func (g *Grid) AliveNeighbours(x int, y int) int {
	return 0
}
