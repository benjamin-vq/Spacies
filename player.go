package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
)

type Player struct {
	x, y  float32
	speed float32
	size  int
}

func NewPlayer(x, y float32) *Player {
	log.Info().Msgf("Creating new player with starting position (%.0f, %.0f)", x, y)
	return &Player{
		x:     x,
		y:     y,
		speed: 4,
		size:  20,
	}
}

func (p *Player) Update(width, height int) {
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

}
