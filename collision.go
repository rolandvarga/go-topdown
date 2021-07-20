package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type CollisionObject struct {
	Rect pixel.Rect
}

func NewCollisionObject(rect pixel.Rect) CollisionObject {
	return CollisionObject{Rect: rect}
}

func (co *CollisionObject) Draw(imd *imdraw.IMDraw) {
	imd.Color = colornames.Greenyellow
	imd.Push(co.Rect.Min, co.Rect.Max)
	imd.Rectangle(1)
}
