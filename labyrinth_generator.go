package main

import (
	"flag"
	"fmt"
	"image/png"
	"labyrinth_generator/labyrinth"
	"labyrinth_generator/labyrinth/generator"
	"os"
)

func main() {
	width := flag.Int("width", 11, "Width of the generated labyrinth")
	height := flag.Int("height", 11, "Height of the generated labyrinth")
	mode := flag.Int("mode", 1, "Generation mode:\n1 - recursive\n2 - growing tree\n3 - eller")
	outmode := flag.Int("outmode", 1, "Output mode\n1 - ASCII\n2 - image")

	flag.Parse()

	var gen generator.LabyrinthGenerator
	switch *mode {
	case 2:
		gen = generator.GrowingTreeGenerator{}
	case 3:
		gen = generator.EllerGenerator{}
	default:
		gen = generator.RecursiveGenerator{}
	}

	if *width%2 == 0 {
		*width += 1
	}

	if *height%2 == 0 {
		*height += 1
	}

	l := labyrinth.NewLabyrinth(*width, *height, gen)

	switch *outmode {
	case 2:
		f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
		defer f.Close()
		png.Encode(f, l.Image())
	case 1:
		fmt.Print(l)
	default:
		fmt.Println("Unknown out format")
	}
}
