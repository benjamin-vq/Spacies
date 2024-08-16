package main

import (
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

var (
	debug        = flag.Bool("debug", false, "Outputs debug information during game execution")
	screenWidth  = flag.Int("width", 640, "Screen width")
	screenHeight = flag.Int("height", 480, "Screen height")
)

func main() {
	flag.Parse()
	configureLogging()

	game := NewGame(*screenWidth, *screenHeight)

	ebiten.SetWindowSize(game.config.screenWidth, game.config.screenHeight)
	ebiten.SetWindowTitle("Spacies")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error during game execution")
	}
	log.Info().
		Msg("Exiting game")
}

func configureLogging() {
	output := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.TimeOnly
	})
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %s |", i))
	}

	log.Logger = log.Output(output)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Logging level set to DEBUG, will output extra information during execution")
	}
}
