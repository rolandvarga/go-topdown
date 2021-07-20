package main

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

const (
	ENEMY_VISION_RANGE   = 400
	ENEMY_MOVEMENT_SPEED = 5
)

type Enemy struct {
	Collider EnemyCollider

	Health int

	Position     pixel.Vec
	LastPosition pixel.Vec
	Rect         pixel.Rect
	Size         pixel.Vec
	Color        color.RGBA
	Matrix       pixel.Matrix
}

func NewEnemy(minX, minY, maxX, maxY float64, color color.RGBA) Enemy {
	rect := pixel.R(minX, minY, maxX, maxY)
	position := rect.Min
	size := rect.Size()

	enemy := Enemy{
		Position: position,
		Rect:     rect,
		Size:     size,
		Color:    color,
		Matrix:   pixel.IM,

		Health: 5,
	}
	enemy.Collider = enemy.updateCollisionBody()
	return enemy
}

func (e Enemy) Draw(imd *imdraw.IMDraw) {
	imd.Color = e.Color
	imd.Push(e.Rect.Min, e.Rect.Max)
	imd.Rectangle(0)
}

func (e Enemy) CanSee(player pixel.Vec) bool {
	return e.Rect.Center().To(player).Len() <= ENEMY_VISION_RANGE
}

func (e Enemy) MoveTowards(player pixel.Vec) Enemy {
	e.LastPosition = e.Position
	direction := pixel.V(
		player.X-e.Position.X,
		player.Y-e.Position.Y,
	)

	hyp := math.Sqrt(direction.X*direction.X + direction.Y*direction.Y)
	direction.X = direction.X / hyp
	direction.Y = direction.Y / hyp

	// update position AND rect
	e.Position.X += direction.X * ENEMY_MOVEMENT_SPEED
	e.Position.Y += direction.Y * ENEMY_MOVEMENT_SPEED

	e.Rect.Min = e.Position
	e.Rect.Max = e.Position.Add(e.Size)

	return e
}

func (e *Enemy) updateCollisionBody() EnemyCollider {
	return NewEnemyCollider(pixel.R(
		e.Position.X,
		e.Position.Y,
		e.Position.X+e.Size.X,
		e.Position.Y+e.Size.Y,
	))
}
