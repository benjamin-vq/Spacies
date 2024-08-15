package main

import (
	"flag"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

var (
	screenWidth, screenHeight int
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
}

func (g *Game) Layout(width, height int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	flag.IntVar(&screenWidth, "width", 640, "Screen width")
	flag.IntVar(&screenHeight, "height", 480, "Screen height")
	flag.Parse()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Spacies")

	log.Fatal(ebiten.RunGame(&Game{}))
}
