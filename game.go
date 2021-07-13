package main

import (
	"image/color"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	WINDOW_WIDTH  = 1024
	WINDOW_HEIGHT = 768

	DEBUG_MODE = false

	BULLET_MAX_AMOUNT = 5
	BULLET_MAX_FRAMES = 10

	FRAME_DELAY = 4
)

const (
	LEFT = iota
	DOWN
	RIGHT
	UP
)

var (
	GRAVITY = true
)

type Game struct {
	WindowColor color.RGBA
	Engine      Engine
	Platforms   []platform // TODO parsed tilemap exported from Tiled, and create platform objects based on that
}

func NewGame() *Game {
	engine := newEngine()
	platforms := []platform{
		{Rect: pixel.R(0, 0, WINDOW_WIDTH, 50), Color: colornames.Brown},                           // bottom
		{Rect: pixel.R(0, WINDOW_HEIGHT-50, WINDOW_WIDTH, WINDOW_HEIGHT), Color: colornames.Brown}, // top
		{Rect: pixel.R(0, 50, 50, WINDOW_HEIGHT), Color: colornames.Brown},                         // left
		// {Rect: pixel.R(WINDOW_WIDTH-50, 0, WINDOW_WIDTH, WINDOW_HEIGHT), Color: colornames.Brown},  // right
	}
	return &Game{WindowColor: colornames.Gray, Engine: engine, Platforms: platforms}
}

func (g *Game) run() {
	// all code goes here
	cfg := pixelgl.WindowConfig{
		Title:  "$$$INSERT COIN$$$",
		Bounds: pixel.R(0, 0, WINDOW_WIDTH, WINDOW_HEIGHT),
		VSync:  true, // sync window framerate with monitor framerate
	}

	// create main window
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	// create player
	playersheet, err := g.Engine.loadPictureAt(g.Engine.Assets["soldier_movement_sprites"])
	if err != nil {
		panic(err)
	}
	player := NewPlayer(playersheet)

	player.Position = win.Bounds().Center() // starts off character middle of screen
	player.Direction = 2                    // starts facing right
	lastPosition := player.Position

	camPos := lastPosition // center camera

	framesElapsed := 0

	last := time.Now()
	for !win.Closed() {
		timeDelta := time.Since(last).Seconds()
		last = time.Now()

		lastPosition = player.Position

		// gravity
		if GRAVITY {
			player.Position.Y -= 10
		}

		if win.Pressed(pixelgl.KeyA) {
			player.Position.X -= player.Speed * timeDelta
			player.Direction = LEFT
			if framesElapsed == FRAME_DELAY {
				player.FrameCount = (player.FrameCount + 1) % 8
				player.ActiveFrame = 11 + player.FrameCount
				framesElapsed = 0
			}
			framesElapsed++
		} else if win.Pressed(pixelgl.KeyD) {
			player.Position.X += player.Speed * timeDelta
			player.Direction = RIGHT
			if framesElapsed == FRAME_DELAY {
				player.FrameCount = (player.FrameCount + 1) % 8
				player.ActiveFrame = 2 + player.FrameCount
				framesElapsed = 0
			}
			framesElapsed++
		} else if win.Pressed(pixelgl.KeyS) {
			// crouch
		} else if win.Pressed(pixelgl.KeyW) {
			// jump
			player.Position.Y += player.Speed * timeDelta * 1.5
			GRAVITY = true
		} else {
			if player.Direction == RIGHT {
				player.ActiveFrame = 1
				player.FrameCount = 0
				framesElapsed = 0
			}
			if player.Direction == LEFT {
				player.ActiveFrame = 10
				player.FrameCount = 0
				framesElapsed = 0
			}
		}
		if win.Pressed(pixelgl.KeySpace) {
			if len(player.Bullets) < BULLET_MAX_AMOUNT {
				player.Shoot(player.Direction)
			}
		}

		player.setCollisionBody()
		if DEBUG_MODE {
			player.Collider.Draw(imd)
		}

		if player.Collider.collidesWith(g.Platforms) {
			player.Position = lastPosition
			GRAVITY = false
		}

		for i := 0; i < len(player.Bullets); i++ {
			b := player.Bullets[i]
			player.Bullets[i] = b.update()

			// if the bullet hits any of the platforms or has reached the maximum
			// number of allowed frames, delete it
			if b.collidesWith(g.Platforms) || b.Frames > BULLET_MAX_FRAMES {
				player.Bullets = append(player.Bullets[:i], player.Bullets[i+1:]...)
				continue
			}
		}

		player.Matrix = pixel.IM.Scaled(pixel.ZV, 4).Moved(player.Position)
		camPos = player.Position                                 // make camera follow player
		cam := pixel.IM.Moved(win.Bounds().Center().Sub(camPos)) // pin to center of screen
		win.SetMatrix(cam)

		win.Clear(g.WindowColor) // changes window color & also clears window

		// draw new sprites here
		player.SpriteMap[player.ActiveFrame].Draw(win, player.Matrix)

		for _, p := range g.Platforms {
			p.Draw(imd)
		}

		for _, b := range player.Bullets {
			b.draw(imd)
		}

		imd.Draw(win)

		win.Update()

		// clear imd after everything else has been updated in current frame
		imd.Reset()
		imd.Clear()
	}
}

func main() {
	game := NewGame()
	pixelgl.Run(game.run)
}
