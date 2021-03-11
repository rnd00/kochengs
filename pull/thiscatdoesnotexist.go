package pull

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/rnd00/kochengs/file"
)

func check(e error) {
	if e != nil {
		log.Fatal().Err(e)
	}
}

// ThisCatDoesNotExist -- pull image from that site
func ThisCatDoesNotExist() {
	r, e := http.Get("https://thiscatdoesnotexist.com")
	check(e)
	defer r.Body.Close()

	n := file.Generate(10)
	wd, e := os.Getwd()
	check(e)
	fp := fmt.Sprintf("%s/%s", wd, n)
	fne := file.Check(fp)
	check(fne)

	o, e := os.Create(fp)
	check(e)
	defer o.Close()
	io.Copy(o, r.Body)

	log.Info().Msgf("%s created", n)
}
