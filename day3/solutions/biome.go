package solutions

type Biome struct {
	internalMap       [][]bool
	positionX         int
	positionY         int
	widthInternalMap  int
	heightInternalMap int
	xStep             int
	yStep             int
}

func NewBiome(internalMap [][]bool, xStep, yStep int) *Biome {
	return &Biome{
		internalMap:       internalMap,
		positionX:         0,
		positionY:         0,
		widthInternalMap:  len(internalMap[0]),
		heightInternalMap: len(internalMap) - 1,
		xStep:             xStep,
		yStep:             yStep,
	}
}

func (b *Biome) step() {
	b.positionX += b.xStep
	b.positionY += b.yStep
}

func (b *Biome) isTree() bool {
	x := b.positionX % b.widthInternalMap
	return b.internalMap[b.positionY][x]
}

func (b *Biome) bottomReached() bool {
	return b.positionY == b.heightInternalMap
}
