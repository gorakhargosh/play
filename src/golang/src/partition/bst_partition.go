package partition

// bstPartition is a weighted quick-union disjoint set partitioner.
type bstPartition struct {
}

// NewBSTPartition generates a new partition.
func NewBSTPartition(size int) Partition {
	return &bstPartition{}
}

func (p *bstPartition) Union(x, y int) {
}

func (p bstPartition) FindSet(x int) int {
	return 0
}

func (p bstPartition) Connected(x, y int) bool {
	return false
}
