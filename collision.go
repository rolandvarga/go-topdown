package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Collider interface {
	collidesWith(coll Collider) bool
	getRect() pixel.Rect
}

type PlayerCollider struct {
	Rect pixel.Rect
}

func NewPlayerCollider(rect pixel.Rect) PlayerCollider {
	return PlayerCollider{Rect: rect}
}

func (co PlayerCollider) getRect() pixel.Rect {
	return co.Rect
}

func (co PlayerCollider) Draw(imd *imdraw.IMDraw) {
	imd.Color = colornames.Greenyellow
	imd.Push(co.Rect.Min, co.Rect.Max)
	imd.Rectangle(1)
}

func (co PlayerCollider) collidesWith(coll Collider) bool {
	return coll.getRect().Intersects(co.Rect)
}

type EnemyCollider struct {
	Rect pixel.Rect
}

func NewEnemyCollider(rect pixel.Rect) EnemyCollider {
	return EnemyCollider{Rect: rect}
}

func (co EnemyCollider) getRect() pixel.Rect {
	return co.Rect
}

func (co EnemyCollider) Draw(imd *imdraw.IMDraw) {
	imd.Color = colornames.Greenyellow
	imd.Push(co.Rect.Min, co.Rect.Max)
	imd.Rectangle(1)
}

func (co EnemyCollider) collidesWith(coll Collider) bool {
	return coll.getRect().Intersects(co.Rect)
}

type PlatformCollider struct {
	Rect pixel.Rect
}

func NewPlatformCollider(rect pixel.Rect) PlatformCollider {
	return PlatformCollider{Rect: rect}
}

func (co PlatformCollider) getRect() pixel.Rect {
	return co.Rect
}

func (co PlatformCollider) Draw(imd *imdraw.IMDraw) {
	imd.Color = colornames.Greenyellow
	imd.Push(co.Rect.Min, co.Rect.Max)
	imd.Rectangle(1)
}

func (co PlatformCollider) collidesWith(coll Collider) bool {
	return coll.getRect().Intersects(co.Rect)
}
