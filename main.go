package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	shibaImage *ebiten.Image
)

type Game struct {
	player Player
	keys   []ebiten.Key
}

const (
	screenWidth  = 320
	screenHeight = 240
)

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	g.player.keyboardMove(g)

	g.player.update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	g.player.drawSprite(screen)
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

	var g Game
	g.player.posX = 32.0
	g.player.posY = 32.0
	g.player.sprite = shibaImage

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
