package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Collider interface {
	collidesWith(objects []Collider, delete bool) bool
}

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

// collidesWith provides collision detection between player & platforms. When delete set to true
// it removes player pixels.
func (co *CollisionObject) collidesWith(platforms []platform, delete bool) bool {
	for _, p := range platforms {
		if p.getRect().Intersects(co.Rect) {
			return true
		}
	}
	return false
}
