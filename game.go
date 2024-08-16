package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/rs/zerolog/log"
	"image/color"
)

type Game struct {
	config Config
	player *Player
}

func NewGame(width, height int) *Game {
	log.Info().Msg("Initializing game")

	startX := float32(width / 2)
	startY := float32(height / 2)

	log.Info().Msgf("Screen Width: %d", width)
	log.Info().Msgf("Screen Height: %d", height)

	return &Game{
		config: NewConfig(width, height),
		player: NewPlayer(startX, startY),
	}
}

func (g *Game) Update() error {
	g.player.Update(g.config.screenWidth, g.config.screenHeight)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	vector.DrawFilledRect(screen, g.player.x, g.player.y, float32(g.player.size), float32(g.player.size), color.White, false)
}

func (g *Game) Layout(width, height int) (int, int) {
	return g.config.screenWidth, g.config.screenHeight
}
