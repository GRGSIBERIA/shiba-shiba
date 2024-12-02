package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type Actor struct {
	posX, posY float64
	spdX, spdY float64
	accX, accY float64
	sprite     *ebiten.Image
}

type Player struct {
	Actor
}

func (s *Actor) init(posX, posY, spdX, spdY float64, sprite *ebiten.Image) {
	s.posX = posX
	s.posY = posY
	s.spdX = spdX
	s.spdY = spdY
	s.accX = 0.0
	s.accY = 0.0
	s.sprite = sprite
}

func (s *Actor) update() {
	const dt float64 = 0.0166666666666667
	s.spdX += s.accX * dt
	s.spdY += s.accY * dt
	s.posX += s.spdX * dt
	s.posY += s.spdY * dt
	fmt.Println(s.posX, s.posY)
}

func (s *Actor) drawSprite(screen *ebiten.Image) {
	op := &colorm.DrawImageOptions{}
	var c colorm.ColorM
	const size float64 = 1.0 / 16.0

	op.GeoM.Translate(s.posX, s.posY)
	op.GeoM.Scale(size, size)

	colorm.DrawImage(screen, s.sprite, c, op)
}

func (s *Player) keyboardMove(g *Game) {
	const acc float64 = 1.0
	for _, k := range g.keys {
		switch k.String() {
		case "W":
			s.spdY += -acc
		case "A":
			s.spdX += -acc
		case "S":
			s.spdY += acc
		case "D":
			s.spdX += acc
		}
	}

	//var norm float64 = 1.0 / math.Sqrt(s.accX*s.accX+s.accY*s.accY)
	//s.accX *= norm
	//s.accY *= norm
}
