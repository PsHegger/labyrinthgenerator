package labyrinth

import (
	"bytes"
	"labyrinth_generator/labyrinth/generator"
	"image"
	"image/color"
)

type Labyrinth struct {
	w, h int
	data [][]bool
}

func (l Labyrinth) String() string {
	var buffer bytes.Buffer

	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			if l.data[y][x] {
				buffer.WriteString(" ")
			} else {
				buffer.WriteString("#")
			}
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

func (l Labyrinth) Image() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, l.w, l.h))

	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			if !l.data[y][x] {
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			} else {
				img.Set(x, y, color.RGBA{255, 255, 255, 255})
			}
		}
	}

	return img
}

func NewLabyrinth(w, h int, generator generator.LabyrinthGenerator) *Labyrinth {
	l := new(Labyrinth)
	l.w = w
	l.h = h

	if generator != nil {
		l.data = generator.Generate(w, h)
	} else {
		l.data = make([][]bool, h)
		for i := 0; i < h; i++ {
			l.data[i] = make([]bool, w)
		}
	}

	return l
}
