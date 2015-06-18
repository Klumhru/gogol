package objects

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
func (g *Grid) Randomize(weight float32) error {
	return nil
}

// Checks if cell at specified position is alive
func (g *Grid) Alive(x int, y int) bool {
	return true
}

// Returns the total number of live cells in the grid
func (g *Grid) CountLiveCells() int {
	return 0
}

// Gets the life status of a cell
func (g *Grid) Get(x int, y int) bool {
	return false
}

// Sets the life status of a cell
func (g *Grid) Set(x int, y int, status bool) {
}

// Projects the status of the grid to a new grid, based on Conways GoL
func (in *Grid) Project(out *Grid) {

}
