package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type bullet struct {
	Rect      pixel.Rect
	Direction int
	Frames    int
	color     color.RGBA
}

func NewBullet(position pixel.Vec, direction int) bullet {
	return bullet{
		Rect: pixel.R(
			position.X,
			position.Y-10,
			position.X+25,
			position.Y+20,
		),
		Direction: direction,

		Frames: 0, // kill after N number of frames

		color: colornames.Blueviolet,
	}
}

func (b *bullet) draw(imd *imdraw.IMDraw) {
	imd.Color = b.color
	imd.Push(b.Rect.Min, b.Rect.Max)
	imd.Rectangle(0)
}

func (b *bullet) update() bullet {
	b.Frames++

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
