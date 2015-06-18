package main

import (
	"flag"
	"fmt"
	"github.com/Klumhru/gogol/objects"
	"os"
	"os/exec"
	"time"
)

var width, height, framerate int

func init() {
	flag.IntVar(&width, "width", 10, "The grid width")
	flag.IntVar(&height, "height", 10, "The grid height")
	flag.IntVar(&framerate, "framerate", 2, "The number of updates to perform per second")
	flag.Parse()
}

func clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	a := objects.MakeGrid(width, height)
	b := objects.MakeGrid(width, height)
	grids := map[bool]*objects.Grid{
		false: &a,
		true:  &b,
	}
	grids[false].Randomize(0.3)
	next := false

	for {
		clear()
		fmt.Println("==== GAME OF LIFE ====")
		grids[next].Render(os.Stdout)
		grids[next].Project(grids[!next])
		next = !next
		time.Sleep(time.Millisecond * time.Duration(1000/framerate))
	}
}
