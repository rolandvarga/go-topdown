package main

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

const (
	ENEMY_VISION = 400
	ENEMY_SPEED  = 5
)

type Enemy struct {
	Collider CollisionObject

	Position pixel.Vec
	Rect     pixel.Rect
	Size     pixel.Vec
	Color    color.RGBA
	Matrix   pixel.Matrix
}

func NewEnemy() Enemy {
	rect := pixel.R(WINDOW_WIDTH, 50, WINDOW_WIDTH+150, 250)
	position := rect.Min
	size := rect.Size()

	return Enemy{
		Position: position,
		Rect:     rect,
		Size:     size,
		Color:    colornames.Green,
		Matrix:   pixel.IM,
	}
}

func (e *Enemy) Draw(imd *imdraw.IMDraw) {
	imd.Color = e.Color
	imd.Push(e.Rect.Min, e.Rect.Max)
	imd.Rectangle(0)
}

func (e *Enemy) CanSee(player pixel.Vec) bool {
	return e.Rect.Center().To(player).Len() <= ENEMY_VISION
}

func (e Enemy) MoveTowards(player pixel.Vec) Enemy {
	direction := pixel.V(
		player.X-e.Position.X,
		player.Y-e.Position.Y,
	)

	hyp := math.Sqrt(direction.X*direction.X + direction.Y*direction.Y)
	direction.X = direction.X / hyp
	direction.Y = direction.Y / hyp

	// update position AND rect
	e.Position.X += direction.X * ENEMY_SPEED
	e.Position.Y += direction.Y * ENEMY_SPEED

	e.Rect.Min = e.Position
	e.Rect.Max = e.Position.Add(e.Size)

	return e
}
