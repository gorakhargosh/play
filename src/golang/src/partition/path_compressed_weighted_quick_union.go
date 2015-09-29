package partition

// pathCompressedPartition is a path-compressed weighted quick-union disjoint set
// partitioner.
type pathCompressedPartition struct {
}

// NewPathCompressedPartition generates a new partition.
func NewPathCompressedPartition(size int) Partition {
	return &pathCompressedPartition{}
}

func (q *pathCompressedPartition) Union(x, int int) {
}

func (q pathCompressedPartition) FindSet(x int) int {
	return 0
}

func (q pathCompressedPartition) Connected(x, y int) bool {
	return false
}
