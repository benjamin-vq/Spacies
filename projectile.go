package main

import "github.com/rs/zerolog/log"

type Projectile struct {
	X, Y   float32
	Speed  float32
	Size   int
	Active bool
}

func NewProjectile(x, y float32) *Projectile {
	return &Projectile{
		X:      x,
		Y:      y,
		Speed:  5,
		Size:   5,
		Active: true,
	}
}

func (p *Projectile) Update() {
	p.Y -= p.Speed

	isOffScreen := p.Y+float32(p.Size) < 0
	if isOffScreen {
		log.Debug().Msgf("Projectile off screen, marking as inactive")
		p.Active = false
	}
}
