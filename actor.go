package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type Actor struct {
	posX, posY float64
	spdX, spdY float64
	sprite     *ebiten.Image
}

func (s *Actor) init(posX, posY, spdX, spdY float64, sprite *ebiten.Image) {
	s.posX = posX
	s.posY = posY
	s.spdX = spdX
	s.spdY = spdY
	s.sprite = sprite
}

func (s *Actor) drawSprite(screen *ebiten.Image, sprite *ebiten.Image, x float64, y float64) {
	op := &colorm.DrawImageOptions{}
	var c colorm.ColorM
	const size float64 = 1.0 / 16.0

	op.GeoM.Translate(x, y)
	op.GeoM.Scale(size, size)

	colorm.DrawImage(screen, sprite, c, op)
}
