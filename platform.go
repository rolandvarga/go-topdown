package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type platform struct {
	Rect  pixel.Rect
	Color color.RGBA
}

func (p *platform) Draw(imd *imdraw.IMDraw) {
	imd.Color = p.Color
	imd.Push(p.Rect.Min, p.Rect.Max)
	imd.Rectangle(0)
}

func (p *platform) getRect() pixel.Rect {
	return p.Rect
}
