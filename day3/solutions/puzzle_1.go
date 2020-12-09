package solutions

import "fmt"

const xStep = 3
const yStep = 1

func Puzzle1(puzzleInputFile string) {
	input := readInputFile(puzzleInputFile)

	biome := NewBiome(input, xStep, yStep)

	treeCount := 0
	for {
		biome.step()
		if biome.isTree() {
			treeCount++
		}
		if biome.bottomReached() {
			break
		}
	}
	fmt.Printf("puzzle 1 -- tree count: %d\n", treeCount)
}

