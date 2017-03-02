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

	cx := w / 2
	cy := h / 2

	if cx%2 == 0 {
		cx += 1
	}

	if cy%2 == 0 {
		cy += 1
	}

	var remaining []coords
	remaining = append(remaining, coords{cx, cy})

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

	entrance := rand.Intn(w-2) + 1
	for !data[1][entrance] {
		entrance = rand.Intn(w-2) + 1
	}
	data[0][entrance] = true

	exit := rand.Intn(w-2) + 1
	for !data[h-2][exit] {
		exit = rand.Intn(w-2) + 1
	}
	data[h-1][exit] = true

	return data
}

func deleteCoord(n int, c []coords) []coords {
	return append(c[:n], c[n+1:]...)
}
