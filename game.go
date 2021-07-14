package main

import (
	"image/color"
	"math"
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

	RUN_FRAME_DELAY     = 4
	RUN_MOVEMENT_SPEED  = 500
	JUMP_FRAME_DELAY    = 2
	JUMP_FRAMES_MAX     = 10
	JUMP_MOVEMENT_SPEED = 2200
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
	Enemies     []Enemy
}

func NewGame() *Game {
	engine := newEngine()
	platforms := []platform{
		{Rect: pixel.R(0, 0, WINDOW_WIDTH*3, 50), Color: colornames.Brown},                         // bottom
		{Rect: pixel.R(0, WINDOW_HEIGHT-50, WINDOW_WIDTH, WINDOW_HEIGHT), Color: colornames.Brown}, // top
		{Rect: pixel.R(0, 50, 150, 150), Color: colornames.Brown},                                  // left
		// {Rect: pixel.R(WINDOW_WIDTH-50, 0, WINDOW_WIDTH, WINDOW_HEIGHT), Color: colornames.Brown},  // right
	}
	enemies := []Enemy{
		NewEnemy(),
	}
	return &Game{WindowColor: colornames.Gray, Engine: engine, Platforms: platforms, Enemies: enemies}
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
	player.Direction = RIGHT                // starts facing right
	lastPosition := player.Position

	camPos := lastPosition // center camera

	elapsedFramesRun := 0
	elapsedFramesJump := 0
	totalFramesJump := 0

	// game loop
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
			player.Position.X -= RUN_MOVEMENT_SPEED * timeDelta
			player.Direction = LEFT
			if elapsedFramesRun == RUN_FRAME_DELAY {
				player.FrameCount = (player.FrameCount + 1) % 8
				player.ActiveFrame = 11 + player.FrameCount
				elapsedFramesRun = 0
			}
			elapsedFramesRun++
		} else if win.Pressed(pixelgl.KeyD) {
			player.Position.X += RUN_MOVEMENT_SPEED * timeDelta
			player.Direction = RIGHT
			if elapsedFramesRun == RUN_FRAME_DELAY {
				player.FrameCount = (player.FrameCount + 1) % 8
				player.ActiveFrame = 2 + player.FrameCount
				elapsedFramesRun = 0
			}
			elapsedFramesRun++
		} else if win.Pressed(pixelgl.KeyS) {
			// crouch
		} else {
			if player.Direction == RIGHT {
				player.ActiveFrame = 1
				player.FrameCount = 0
				elapsedFramesRun = 0
			}
			if player.Direction == LEFT {
				player.ActiveFrame = 10
				player.FrameCount = 0
				elapsedFramesRun = 0
			}
		}
		if win.Pressed(pixelgl.KeyJ) {
			// jump
			if player.OnGround {
				player.Jumping = true
				player.OnGround = false
				GRAVITY = true

				switch player.Direction {
				case LEFT:
					player.ActiveFrame = 11
				case RIGHT:
					player.ActiveFrame = 2
				}
				player.Position.Y += JUMP_MOVEMENT_SPEED * timeDelta
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
			player.OnGround = true
			player.Jumping = false
			totalFramesJump = 0
			elapsedFramesJump = 0
		}

		if player.Jumping {
			if totalFramesJump >= JUMP_FRAMES_MAX {
				player.Position.Y = math.Max(player.Position.Y-JUMP_MOVEMENT_SPEED*timeDelta, 0)
			}
			if elapsedFramesJump >= JUMP_FRAME_DELAY {
				player.Position.Y += JUMP_MOVEMENT_SPEED * timeDelta
			}
			totalFramesJump++
			elapsedFramesJump++
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

		for i := 0; i < len(g.Enemies); i++ {
			e := g.Enemies[i]
			if e.CanSee(player.Position) {
				g.Enemies[i] = e.MoveTowards(player.Position)
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

		for _, e := range g.Enemies {
			e.Draw(imd)
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
