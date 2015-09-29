package partition

// weightedQuickUnion is a weighted quick-union disjoint set partitioner.
type weightedQuickUnion struct {
}

// NewWeightedQuickUnion generates a new partition.
func NewWeightedQuickUnion() Partition {
	return &weightedQuickUnion{}
}

func (q *weightedQuickUnion) Union(x, int int) {
}

func (q weightedQuickUnion) FindSet(x int) int {
	return 0
}

func (q weightedQuickUnion) Connected(x, y int) bool {
	return false
}
