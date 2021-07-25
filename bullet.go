package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

const (
	BULLET_PIXEL_WIDTH  = 49
	BULLET_PIXEL_HEIGHT = 27
)

type Bullet struct {
	Position    pixel.Vec
	Rect        pixel.Rect
	Direction   int
	ActiveFrame int // determines which sprite should be rendered
	Frames      int
	color       color.RGBA

	Sheet     pixel.Picture
	SpriteMap map[int]*pixel.Sprite
	Matrix    pixel.Matrix
}

func NewBullet(position pixel.Vec, direction int, bulletSheet pixel.Picture) Bullet {
	spriteMap := make(map[int]*pixel.Sprite)
	for i := 0; i < 19; i++ {
		spriteMap[i] = pixel.NewSprite(bulletSheet, pixel.R(
			float64(BULLET_PIXEL_WIDTH*(i-1)+2), 0, // Rect Min
			float64(BULLET_PIXEL_WIDTH*i+2), BULLET_PIXEL_HEIGHT, // Rect Max
		))
	}
	switch direction {
	case LEFT:
		position.X -= 139
	case RIGHT:
		position.X += 125
	}
	return Bullet{
		Position: position,
		Rect: pixel.R(
			position.X,
			position.Y-10,
			position.X+25,
			position.Y+20,
		),
		Direction: direction,

		Sheet:       bulletSheet,
		SpriteMap:   spriteMap,
		ActiveFrame: 1,
		Frames:      0, // kill after N number of frames

		color: colornames.Blueviolet,
	}
}

func (b *Bullet) draw(imd *imdraw.IMDraw) {
	imd.Color = b.color
	imd.Push(b.Rect.Min, b.Rect.Max)
	imd.Rectangle(0)
}

func (b *Bullet) update() Bullet {
	b.Frames++
	b.ActiveFrame++

	if b.Direction == LEFT {
		b.Rect.Min = pixel.V(b.Rect.Min.X-20, b.Rect.Min.Y)
		b.Rect.Max = pixel.V(b.Rect.Max.X-20, b.Rect.Max.Y)
	}
	if b.Direction == DOWN {
		b.Rect.Min = pixel.V(b.Rect.Min.X, b.Rect.Min.Y-10)
		b.Rect.Max = pixel.V(b.Rect.Max.X, b.Rect.Max.Y-10)
	}
	if b.Direction == RIGHT {
		b.Rect.Min = pixel.V(b.Rect.Min.X+20, b.Rect.Min.Y)
		b.Rect.Max = pixel.V(b.Rect.Max.X+20, b.Rect.Max.Y)
	}
	if b.Direction == UP {
		b.Rect.Min = pixel.V(b.Rect.Min.X, b.Rect.Min.Y+10)
		b.Rect.Max = pixel.V(b.Rect.Max.X, b.Rect.Max.Y+10)
	}
	return *b
}

func (b Bullet) collidesWith(coll Collider) bool {
	return coll.getRect().Intersects(b.Rect)
}
