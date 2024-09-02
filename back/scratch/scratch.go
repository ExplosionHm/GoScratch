package scratch

import (
	"image"
	_ "image/png"
	"os"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

var (
	IM = pixel.IM
	Runner = pixelgl.Run
)

func CreateWindow() *pixelgl.Window {
	w, err := pixelgl.NewWindow(pixelgl.WindowConfig{Title: "GoScratch Window", Bounds: pixel.R(0, 0, 1024, 768), Resizable: true, TransparentFramebuffer: true})
	if err != nil {
		panic(err)
	}
	return w
}

func CreateSprite(path string) *pixel.Sprite {
	img, err := LoadImage(path)
	if err != nil {
		panic(err)
	}
	return pixel.NewSprite(img, img.Bounds())
}

func LoadImage(path string) (pixel.Picture, error) {
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