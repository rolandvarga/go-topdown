package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const (
	ENEMY_VISION = 400
)

type Enemy struct {
	Collider CollisionObject

	Rect   pixel.Rect
	Color  color.RGBA
	Matrix pixel.Matrix
}

func (e *Enemy) Draw(imd *imdraw.IMDraw) {
	imd.Color = e.Color
	imd.Push(e.Rect.Min, e.Rect.Max)
	imd.Rectangle(0)
}

func (e *Enemy) CanSee(player pixel.Vec) bool {
	return e.Rect.Center().To(player).Len() <= ENEMY_VISION
}

func (e *Enemy) MoveTowards(player pixel.Vec) pixel.Rect {
	return e.Rect.Moved(pixel.V(player.X/5, player.Y/5))
}
