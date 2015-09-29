package partition

// treePartition is a weighted quick-union disjoint set partitioner.
type treePartition struct {
}

// NewTreePartition generates a new partition.
func NewTreePartition() Partition {
	return &treePartition{}
}

func (q *treePartition) Union(x, int int) {
}

func (q treePartition) FindSet(x int) int {
	return 0
}

func (q treePartition) Connected(x, y int) bool {
	return false
}
