package gogol

import (
	"github.com/Klumhru/gogol/objects"
	"os"
)

func main() {
	a := objects.MakeGrid(5, 5)
	b := objects.MakeGrid(5, 5)
	grids := map[bool]*objects.Grid{
		false: &a,
		true:  &b,
	}
	grids[false].Randomize(0.3)
	next := false

	for {
		grids[next].Render(os.Stdout)
		grids[next].Project(grids[!next])
		next = !next
	}
}
