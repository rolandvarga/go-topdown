package main

import (
	"github.com/faiface/pixel"
)

const (
	PLAYER_PIXEL_WIDTH  = 20
	PLAYER_PIXEL_HEIGHT = 27
)

type Player struct {
	Position    pixel.Vec
	Collider    PlayerCollider
	Direction   int
	ActiveFrame int // determines which sprite should be rendered
	FrameCount  int
	OnGround    bool
	Jumping     bool

	Health  int
	Bullets []Bullet

	PlayerSheet pixel.Picture
	BulletSheet pixel.Picture
	SpriteMap   map[int]*pixel.Sprite
	Matrix      pixel.Matrix
}

func NewPlayer(playerSheet, bulletSheet pixel.Picture) Player {
	spriteMap := make(map[int]*pixel.Sprite)
	for i := 0; i < 19; i++ {
		spriteMap[i] = pixel.NewSprite(playerSheet, pixel.R(
			float64(PLAYER_PIXEL_WIDTH*(i-1)+2), 0, // Rect Min
			float64(PLAYER_PIXEL_WIDTH*i+2), PLAYER_PIXEL_HEIGHT, // Rect Max
		))
	}
	activeFrame := 1 // start facing right

	return Player{
		PlayerSheet: playerSheet,
		BulletSheet: bulletSheet,
		SpriteMap:   spriteMap,
		ActiveFrame: activeFrame,
	}
}

func (p *Player) updateCollisionBody() PlayerCollider {
	return NewPlayerCollider(pixel.R(
		p.Position.X-(p.SpriteMap[p.ActiveFrame].Frame().W()),
		p.Position.Y-(p.SpriteMap[p.ActiveFrame].Frame().H()+10),
		p.Position.X+(p.SpriteMap[p.ActiveFrame].Frame().W()-18),
		p.Position.Y+(p.SpriteMap[p.ActiveFrame].Frame().H()+10),
	))
}

func (p *Player) Shoot(direction int) Bullet {
	return NewBullet(p.Position, direction, p.BulletSheet)
}
