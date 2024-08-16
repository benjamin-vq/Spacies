package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/rs/zerolog/log"
	"image/color"
)

type Game struct {
	config      Config
	player      *Player
	projectiles []*Projectile
}

func NewGame(width, height int) *Game {
	log.Info().Msg("Initializing game")

	startX := float32(width / 2)
	startY := float32(height / 2)

	log.Info().Msgf("Screen Width: %d", width)
	log.Info().Msgf("Screen Height: %d", height)

	return &Game{
		config:      NewConfig(width, height),
		player:      NewPlayer(startX, startY),
		projectiles: make([]*Projectile, 0, 10),
	}
}

func (g *Game) Update() error {
	projectile := g.player.Update(g.config.screenWidth, g.config.screenHeight)
	if projectile != nil {
		g.projectiles = append(g.projectiles, projectile)
	}

	// Clean inactive projectiles
	activeProjectiles := g.projectiles[:0]
	for _, proj := range g.projectiles {
		proj.Update()
		if proj.Active {
			activeProjectiles = append(activeProjectiles, proj)
		}
	}
	g.projectiles = activeProjectiles

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(g.player.x), float64(g.player.y))
	screen.DrawImage(g.player.image, opts)
	for _, proj := range g.projectiles {
		vector.DrawFilledRect(screen, proj.X, proj.Y, float32(proj.Size), float32(proj.Size), color.White, false)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return g.config.screenWidth, g.config.screenHeight
}
