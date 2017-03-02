package generator

import (
	"math/rand"
	"time"
)

type GrowingTreeGenerator struct{}

func (GrowingTreeGenerator) Generate(w, h int) [][]bool {
	rand.Seed(time.Now().Unix())

	data := make([][]bool, h)

	for i := 0; i < h; i++ {
		data[i] = make([]bool, w)
	}

	var remaining []coords
	remaining = append(remaining, coords{w / 2, h / 2})

	for len(remaining) > 0 {
		n := rand.Intn(len(remaining))

		c := remaining[n]

		dests := getPossibleDestinations(c, w, h, data)

		ld := len(dests)

		if ld > 0 {
			dst := dests[rand.Intn(ld)]
			buildTunnel(c, dst, &data)

			remaining = append(remaining, dst)
		} else {
			remaining = deleteCoord(n, remaining)
		}
	}

	return data
}

func deleteCoord(n int, c []coords) []coords {
	return append(c[:n], c[n+1:]...)
}
