package file

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/rnd00/kochengs/rcd"
)

// Generate will generate filename
func Generate(l int) string {
	// use time(unix) as seed
	var sr *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	// generate string using cs as base/seed for the name
	const cs = "aiueo12345"
	b := make([]byte, l)
	for i := range b {
		b[i] = cs[sr.Intn(len(cs))]
	}

	// compile into a string
	n := fmt.Sprintf("kc[%s].jpg", string(b))
	return string(n)
}

// Check will check if filename existed already or not
func Check(x string) error {
	a, _ := os.Stat(x)
	if a != nil {
		return errors.New(rcd.Stats(rcd.FilenameExist))
	}
	return nil
}
