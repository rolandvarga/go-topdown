package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Platform struct {
	Collider Collider

	Rect     pixel.Rect
	Color    color.RGBA
	Position pixel.Vec
	Size     pixel.Vec
}

func NewPlatform(minX, minY, maxX, maxY float64, color color.RGBA) Platform {
	rect := pixel.R(minX, minY, maxX, maxY)
	position := rect.Min
	size := rect.Size()

	platform := Platform{
		Position: position,
		Rect:     rect,
		Size:     size,
		Color:    color,
		Collider: NewPlatformCollider(rect),
	}
	return platform
}

func (p Platform) Draw(imd *imdraw.IMDraw) {
	imd.Color = p.Color
	imd.Push(p.Rect.Min, p.Rect.Max)
	imd.Rectangle(0)
}

func (p Platform) getRect() pixel.Rect {
	return p.Rect
}
