package main

import (
	"fmt"
	"goscratch/scratch"
	"image/color"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

func main() {
	fmt.Println("Started")
	pixelgl.Run(Game)
}

var (
	test = ""
)

func Game() {
	/*@INSERT start*/
	windows := scratch.CreateWindow("GoScratch", 1024, 768, true)
	windows.SetSmooth(true)

	s := scratch.CreateSprite("C:/Users/sandn/Documents/Projects/goscratch/back/images/gopher.png")

	for !windows.Closed() {
		/*@INSERT loop*/
		windows.Clear(color.White)
		s.Draw(windows, pixel.IM.Moved(windows.MousePosition()))
		windows.Update()
	}
}
