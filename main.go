package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player Actor
}

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

	g.player.drawSprite(screen, shibaImage, 64, 64)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func openImagePNG(srcPath string) *ebiten.Image {
	png, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}
	defer png.Close()

	img, _, err := image.Decode(png)
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func init() {
	shibaImage = openImagePNG("./img/shiba.png")
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Shiba Shiba")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
