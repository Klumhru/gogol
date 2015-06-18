package main

import (
	"fmt"
	"github.com/Klumhru/gogol/objects"
)

func main() {
	grid := objects.MakeGrid(5, 5)
	fmt.Print(grid)
}
