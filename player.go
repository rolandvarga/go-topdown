package main

import (
	"github.com/faiface/pixel"
)

const (
	playerPixelWidth  = 20
	playerPixelHeight = 27
)

type Player struct {
	Position    pixel.Vec
	Collider    CollisionObject
	Speed       float64
	Direction   int
	ActiveFrame int // determines which sprite should be rendered
	FrameCount  int

	Bullets []bullet

	Sheet     pixel.Picture
	SpriteMap map[int]*pixel.Sprite
	Matrix    pixel.Matrix
}

func NewPlayer(playerSheet pixel.Picture) Player {
	spriteMap := make(map[int]*pixel.Sprite)
	for i := 0; i < 19; i++ {
		spriteMap[i] = pixel.NewSprite(playerSheet, pixel.R(
			float64(playerPixelWidth*(i-1)+2), 0, // Rect Min
			float64(playerPixelWidth*i+2), playerPixelHeight, // Rect Max
		))
	}
	activeFrame := 1 // start facing right

	return Player{
		Sheet:       playerSheet,
		SpriteMap:   spriteMap,
		ActiveFrame: activeFrame,
		Speed:       500.0,
	}
}

func (p *Player) setCollisionBody() {
	p.Collider = NewCollisionObject(pixel.R(
		p.Position.X-(p.SpriteMap[p.ActiveFrame].Frame().W()),
		p.Position.Y-(p.SpriteMap[p.ActiveFrame].Frame().H()+10),
		p.Position.X+(p.SpriteMap[p.ActiveFrame].Frame().W()-18),
		p.Position.Y+(p.SpriteMap[p.ActiveFrame].Frame().H()+10),
	))
}

func (p *Player) Shoot(direction int) {
	bullet := NewBullet(p.Position, direction)
	p.Bullets = append(p.Bullets, bullet)
}
