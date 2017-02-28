package generator

type LabyrinthGenerator interface {
	Generate(w, h int) [][]bool
}