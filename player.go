package main

import (
	"fmt"

	"github.com/faiface/pixel"
)

const (
	playerPixelWidth  = 16
	playerPixelHeight = 16
)

type Player struct {
	Position pixel.Vec
	Collider CollisionObject
	Speed    float64

	Sheet  pixel.Picture
	Sprite *pixel.Sprite
	Matrix pixel.Matrix
}

func NewPlayer(playerSheet pixel.Picture) Player {
	sprite := pixel.NewSprite(playerSheet, pixel.R(0, 0, playerPixelWidth, playerPixelHeight))
	return Player{Sheet: playerSheet, Sprite: sprite, Speed: 500.0}
}

// collidesWith provides collision detection between player & platforms. When delete set to true
// it removes player pixels.
func (p *Player) collidesWith(platforms []platform, delete bool) bool {
	for _, platform := range platforms {
		if platform.Rect.Intersects(p.Collider.Rect) {
			fmt.Printf("%v - %v\n", platform.Rect, p.Position)
			return true
		}
	}
	return false
}

func (p *Player) setCollision() {
	p.Collider = NewCollisionObject(pixel.R(
		p.Position.X-(p.Sprite.Frame().W()+1),
		p.Position.Y-(p.Sprite.Frame().H()+10),
		p.Position.X+p.Sheet.Bounds().W(),
		p.Position.Y+p.Sheet.Bounds().H(),
	))
}
