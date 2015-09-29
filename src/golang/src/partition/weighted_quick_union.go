package partition

// weightedQuickUnion is a weighted quick-union disjoint set partitioner.
type weightedQuickUnion struct {
	id     []int
	weight []int
}

// NewWeightedQuickUnion generates a new partition.
func NewWeightedQuickUnion(size int) Partition {
	p := &weightedQuickUnion{
		id:     make([]int, size),
		weight: make([]int, size),
	}
	for i := 0; i < size; i++ {
		p.id[i] = i
		p.weight[i] = 1
	}
	return p
}

func (p *weightedQuickUnion) Union(x, y int) {
}

func (p weightedQuickUnion) FindSet(x int) int {
	return 0
}

func (p weightedQuickUnion) Connected(x, y int) bool {
	return false
}
