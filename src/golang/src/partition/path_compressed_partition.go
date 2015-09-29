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
	a := p.FindSet(x)
	b := p.FindSet(y)
	if p.weight[a] < p.weight[b] {
		p.id[a] = b
		p.weight[b] += p.weight[a]
	} else {
		p.id[b] = a
		p.weight[a] += p.weight[b]
	}
}

func (p *pathCompressedPartition) FindSet(x int) int {
	for x != p.id[x] {
		// Make each node point to its grand parent.
		// Path compressed variant.
		p.id[x] = p.id[p.id[x]]
		x = p.id[x]
	}
	return x
}

func (p pathCompressedPartition) Connected(x, y int) bool {
	return p.FindSet(x) == p.FindSet(y)
}
