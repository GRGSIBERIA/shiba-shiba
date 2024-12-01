package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

const (
	screenWidth  = 320
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

	s := shibaImage.Bounds().Size()
	op := &colorm.DrawImageOptions{}
	var c colorm.ColorM
	op.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
	op.GeoM.Scale(1.0/16.0, 1.0/16.0)
	op.GeoM.Translate(float64(screenWidth)/2, float64(screenHeight)/2)

	colorm.DrawImage(screen, shibaImage, c, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func init() {
	png, err := os.Open("./img/shiba.png")
	if err != nil {
		log.Fatal(err)
	}
	defer png.Close()

	img, _, err := image.Decode(png)
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
