package partition

// optimalPartition is a path-compressed weighted quick-union disjoint set
// partitioner.
type optimalPartition struct {
}

// NewPartition generates a new partition.
func NewPartition() Partition {
	return &optimalPartition{}
}

func (q *optimalPartition) Union(x, int int) {
}

func (q optimalPartition) FindSet(x int) int {
	return 0
}

func (q optimalPartition) Connected(x, y int) bool {
	return false
}
