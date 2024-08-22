package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func main() {
	fmt.Println("Started")
	pixelgl.Run(Game)
}

func Game() {
	w, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:     "Hello world",
		Bounds:    pixel.R(0, 0, 1024, 768),
		Resizable: true,
	})
	if err != nil {
		panic(err)
	}
	w.SetSmooth(true)
	img, err := loadPicture("images/gopher.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(img, img.Bounds())
	var (
		rot = 0.0
	)
	for {
		sprite.Draw(w, pixel.IM.Moved(w.MousePosition()).Rotated(w.MousePosition(), rot))
		if w.JustPressed(pixelgl.MouseButtonLeft) {
			rot += 0.5
			fmt.Println("dwjdw", rot)
		}
		w.Update()
		if w.Closed() {
			return
		}
	}
}
