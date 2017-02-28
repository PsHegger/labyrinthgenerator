package generator

import (
	"math/rand"
	"time"
)

type RecursiveGenerator struct{}
type coords struct {
	x, y int
}

func (RecursiveGenerator) Generate(w, h int) [][]bool {
	rand.Seed(time.Now().Unix())

	data := make([][]bool, h)

	for i := 0; i < h; i++ {
		data[i] = make([]bool, w)
	}

	entrance := rand.Intn(w-2) + 1
	for entrance % 2 == 0 {
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

	var dirx int
	switch {
	case dst.x < c.x:
		dirx = -1
	case dst.x > c.x:
		dirx = 1
	default:
		dirx = 0
	}

	var diry int
	switch {
	case dst.y < c.y:
		diry = -1
	case dst.y > c.y:
		diry = 1
	default:
		diry = 0
	}

	x := c.x
	y := c.y
	for !(*d)[dst.y][dst.x] {
		(*d)[y][x] = true
		x += dirx
		y += diry
	}

	recursiveGeneration(dst, w, h, d)
	recursiveGeneration(c, w, h, d)
}

func getPossibleDestinations(c coords, w, h int, d [][]bool) []coords {
	var possibleCoords []coords

	if c.x-2 > 0 && !d[c.y][c.x-2] {
		possibleCoords = append(possibleCoords, coords{c.x - 2, c.y})
	}

	if c.x+2 < w-1 && !d[c.y][c.x+2] {
		possibleCoords = append(possibleCoords, coords{c.x + 2, c.y})
	}

	if c.y-2 > 0 && !d[c.y-2][c.x] {
		possibleCoords = append(possibleCoords, coords{c.x, c.y - 2})
	}

	if c.y+2 < h-1 && !d[c.y+2][c.x] {
		possibleCoords = append(possibleCoords, coords{c.x, c.y + 2})
	}

	return possibleCoords
}
