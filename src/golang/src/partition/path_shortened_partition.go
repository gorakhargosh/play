package partition

// pathShortenedPartition is a path-compressed weighted quick-union disjoint set
// partitioner.
type pathShortenedPartition struct {
	id     []int
	weight []int
}

// NewPathShortenedPartition generates a new partition.
func NewPathShortenedPartition(size int) Partition {
	p := &pathShortenedPartition{
		id:     make([]int, size),
		weight: make([]int, size),
	}
	for i := 0; i < size; i++ {
		p.id[i] = i
		p.weight[i] = 1
	}
	return p
}

func (p *pathShortenedPartition) Union(x, y int) {
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

func (p *pathShortenedPartition) FindSet(x int) int {
	// Two pass variant that sets root for all traversed elements.
	i := x
	for x != p.id[x] {
		x = p.id[x]
	}
	// x is now the root.
	for i != p.id[i] {
		i = p.id[i]
		p.id[i] = x
	}
	return x
}

func (p pathShortenedPartition) Connected(x, y int) bool {
	return p.FindSet(x) == p.FindSet(y)
}
