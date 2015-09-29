package partition

// pathCompressedPartition is a path-compressed weighted quick-union disjoint set
// partitioner.
type pathCompressedPartition struct {
	id     []int
	weight []int
}

// NewPathCompressedPartition generates a new partition.
func NewPathCompressedPartition(size int) Partition {
	p := &pathCompressedPartition{
		id:     make([]int, size),
		weight: make([]int, size),
	}
	for i := 0; i < size; i++ {
		p.id[i] = i
		p.weight[i] = 1
	}
	return p
}

func (p *pathCompressedPartition) Union(x, y int) {
}

func (p pathCompressedPartition) FindSet(x int) int {
	return 0
}

func (p pathCompressedPartition) Connected(x, y int) bool {
	return false
}
