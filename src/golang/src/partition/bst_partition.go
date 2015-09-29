package partition

// bstPartition is a weighted quick-union disjoint set partitioner.
type bstPartition struct {
}

// NewBSTPartition generates a new partition.
func NewBSTPartition(size int) Partition {
	return &bstPartition{}
}

func (q *bstPartition) Union(x, int int) {
}

func (q bstPartition) FindSet(x int) int {
	return 0
}

func (q bstPartition) Connected(x, y int) bool {
	return false
}
