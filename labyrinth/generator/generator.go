package generator

type LabyrinthGenerator interface {
	Generate(w, h int) [][]bool
}

type coords struct {
	x, y int
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

func buildTunnel(start, end coords, d *[][]bool) {
	var dirx int
	switch {
	case end.x < start.x:
		dirx = -1
	case end.x > start.x:
		dirx = 1
	default:
		dirx = 0
	}

	var diry int
	switch {
	case end.y < start.y:
		diry = -1
	case end.y > start.y:
		diry = 1
	default:
		diry = 0
	}

	x := start.x
	y := start.y
	for !(*d)[end.y][end.x] {
		(*d)[y][x] = true
		x += dirx
		y += diry
	}
}
