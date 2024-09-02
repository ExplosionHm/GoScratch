package main

import (
	/*@INSERT imports*/
	"fmt"
	"goscratch/scratch"
	"image/color"
)

func main() {
	fmt.Println("Started")
	scratch.Runner(Game)
}

var (
	/*@INSERT defintions*/
	test int
)

func Game() {
	/*@INSERT main*/
	window := scratch.CreateWindow()
	window.SetSmooth(true)
	s := scratch.CreateSprite("C:/Users/sandn/Documents/Projects/goscratch/back/images/gopher.png")
	for !window.Closed() {
		/*@INSERT loop*/
		window.Clear(color.White)
		s.Draw(window, scratch.IM.Moved(window.MousePosition()))
		window.Update()
	}
}
