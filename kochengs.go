package main

import (
	"flag"

	"github.com/rnd00/kochengs/pull"
	"github.com/rs/zerolog/log"
)

func main() {
	amt := flag.Int("amt", 1, "Amount of cats will to generate")
	flag.Parse()

	log.Info().Msg("Kocheng started!")
	log.Info().Msgf("Will run for %d times!", *amt)
	for i := 0; i < *amt; i++ {
		pull.ThisCatDoesNotExist()
	}
}
