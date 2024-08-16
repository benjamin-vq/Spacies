package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rs/zerolog/log"
	_ "image/png"
)

type Player struct {
	x, y  float32
	speed float32
	// TODO: Scale the image instead of having a size
	size          int
	image         *ebiten.Image
	shootCooldown int
}

func NewPlayer(x, y float32) *Player {
	img, _, err := ebitenutil.NewImageFromFile("zer.png")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not load player image")
	}

	log.Info().Msgf("Creating new player with starting position (%.0f, %.0f)", x, y)
	return &Player{
		x:             x,
		y:             y,
		speed:         4,
		size:          15,
		image:         img,
		shootCooldown: 0,
	}
}

func (p *Player) Update(width, height int) *Projectile {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.x = max(0, p.x-p.speed)
		log.Debug().Msgf("[KEY LEFT] Updated position to (%.0f, %.0f)", p.x, p.y)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.x = min(float32(width)-float32(p.size), p.x+p.speed)
		log.Debug().Msgf("[KEY RIGHT] Updated position to (%.0f, %.0f)", p.x, p.y)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.y = max(0, p.y-p.speed)
		log.Debug().Msgf("[KEY UP] Updated position to (%.0f, %.0f)", p.x, p.y)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.y = min(float32(height)-float32(p.size), p.y+p.speed)
		log.Debug().Msgf("[KEY DOWN] Updated position to (%.0f, %.0f)", p.x, p.y)
	}

	if p.shootCooldown > 0 {
		p.shootCooldown--
	}

	var projectile *Projectile
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.shootCooldown == 0 {
		// TODO: Should probably re-use projectiles instead of allocating new ones everytime
		projectile = NewProjectile(p.x+float32(p.image.Bounds().Dx())/2, p.y)
		p.resetShootCooldown()
		log.Debug().Msgf("[SPACE] Projectile created and shot from (%.0f, %.0f)",
			projectile.X, projectile.Y)
	}

	return projectile
}

func (p *Player) resetShootCooldown() {
	p.shootCooldown = 15
}
