package main

import (
	/*@INSERT imports*/
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
	/*@INSERT defintions*/
	test int
)

func Game() {
	/*@INSERT main*/
	window := scratch.CreateWindow("GoScratch", 1024, 768, true)
	window.SetSmooth(true)

	s := scratch.CreateSprite("C:/Users/sandn/Documents/Projects/goscratch/back/images/gopher.png")

	for !window.Closed() {
		/*@INSERT loop*/
		window.Clear(color.White)
		s.Draw(window, pixel.IM.Moved(window.MousePosition()))
		window.Update()
	}
}
