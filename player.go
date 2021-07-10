package main

import (
	"github.com/faiface/pixel"
)

const (
	playerPixelWidth  = 16
	playerPixelHeight = 16
)

type Player struct {
	Position  pixel.Vec
	Collider  CollisionObject
	Speed     float64
	Direction int

	Bullets []bullet

	Sheet  pixel.Picture
	Sprite *pixel.Sprite
	Matrix pixel.Matrix
}

func NewPlayer(playerSheet pixel.Picture) Player {
	sprite := pixel.NewSprite(playerSheet, pixel.R(0, 0, playerPixelWidth, playerPixelHeight))
	return Player{Sheet: playerSheet, Sprite: sprite, Speed: 500.0}
}

func (p *Player) setCollision() {
	p.Collider = NewCollisionObject(pixel.R(
		p.Position.X-(p.Sprite.Frame().W()+1),
		p.Position.Y-(p.Sprite.Frame().H()+10),
		p.Position.X+p.Sheet.Bounds().W(),
		p.Position.Y+p.Sheet.Bounds().H(),
	))
}

func (p *Player) Shoot(direction int) {
	bullet := NewBullet(p.Position, direction)
	p.Bullets = append(p.Bullets, bullet)
}
