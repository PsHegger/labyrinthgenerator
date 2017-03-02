package generator

import (
	"math/rand"
	"time"
)

type RecursiveGenerator struct{}

func (RecursiveGenerator) Generate(w, h int) [][]bool {
	rand.Seed(time.Now().Unix())

	data := make([][]bool, h)

	for i := 0; i < h; i++ {
		data[i] = make([]bool, w)
	}

	entrance := rand.Intn(w-2) + 1
	for entrance%2 == 0 {
		entrance = rand.Intn(w-2) + 1
	}
	data[0][entrance] = true
	data[1][entrance] = true

	recursiveGeneration(coords{entrance, 1}, w, h, &data)

	exit := rand.Intn(w-2) + 1
	for !data[h-2][exit] {
		exit = rand.Intn(w-2) + 1
	}
	data[h-1][exit] = true

	return data
}

func recursiveGeneration(c coords, w, h int, d *[][]bool) {
	possibleDests := getPossibleDestinations(c, w, h, *d)

	if len(possibleDests) == 0 {
		return
	}

	dst := possibleDests[rand.Intn(len(possibleDests))]

	buildTunnel(c, dst, d)

	recursiveGeneration(dst, w, h, d)
	recursiveGeneration(c, w, h, d)
}
