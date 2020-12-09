package solutions

import "fmt"

func Puzzle2(puzzleInputFile string) {
	input := readInputFile(puzzleInputFile)
	treeCount11 := treeCountForSlope(input, 1, 1)
	treeCount31 := treeCountForSlope(input, 3, 1)
	treeCount51 := treeCountForSlope(input, 5, 1)
	treeCount71 := treeCountForSlope(input, 7, 1)
	treeCount12 := treeCountForSlope(input, 1, 2)

	fmt.Printf("puzzle 2 -- slope Right 1; down 1; tree count: %d\n", treeCount11)
	fmt.Printf("puzzle 2 -- slope Right 3; down 1; tree count: %d\n", treeCount31)
	fmt.Printf("puzzle 2 -- slope Right 5; down 1; tree count: %d\n", treeCount51)
	fmt.Printf("puzzle 2 -- slope Right 7; down 1; tree count: %d\n", treeCount71)
	fmt.Printf("puzzle 2 -- slope Right 1; down 2; tree count: %d\n", treeCount12)
	fmt.Printf("puzzle 2 -- result: %d\n", treeCount11*treeCount31*treeCount51*treeCount71*treeCount12)
}

func treeCountForSlope(internalMap [][]bool, x, y int) int {
	biome := NewBiome(internalMap, x, y)

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
	return treeCount
}
