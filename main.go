package main

import (
	"io"
	"log"
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

const (
	screenWidth = 320
	screenHeight = 240
)

var (
	shibaImage *ebiten.Image
)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func init() {
	buffer := 
	img, _, err := image.Decode(bytes.NewReader(images.shiba_png))
	if err != nil {
		log.Fatal(err)
	}
	shibaImage = ebiten.NewImageFromImage(img)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}