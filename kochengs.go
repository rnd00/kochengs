package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/rnd00/kochengs/pull"
	"github.com/rs/zerolog/log"
)

func main() {
	amt := flag.Int("amt", 1, "Amount of cats will to generate")
	flag.Parse()

	log.Info().Msg("Kocheng started!")
	log.Info().Msgf("Will run for %d times!", *amt)
	start := time.Now()
	c := make(chan string)
	for i := 0; i < *amt; i++ {
		dur := ((time.Millisecond * 500) * time.Duration(i+1))
		time.Sleep(dur)
		go pull.ThisCatDoesNotExist(c)
	}
	for i := 0; i < *amt; i++ {
		log.Info().Msg(<-c)
	}
	msg := fmt.Sprintf("Finished in %s", time.Since(start))
	log.Info().Msg(msg)
}
