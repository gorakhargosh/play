package partition

// quickFind is a quick-find slow-union partitioner.
type quickFind struct {
	id       []int
	capacity int

	// The presence of a node id in this map indicates that the element has been
	// previously seen in a union operation.
	seen map[int]bool
}

// NewLinearPartition creates a new partition.
func NewLinearPartition(size int) Partition {
	p := &quickFind{
		id: make([]int, size),
	}
	// Initialize the partition to start with each element representing a disjoint
	// subset.
	for i := 0; i < size; i++ {
		p.id[i] = i
	}
	return p
}

func (p *quickFind) Union(x, y int) {
	a := p.FindSet(x)
	b := p.FindSet(y)
	p.seen[a] = true
	p.seen[b] = true
	// Check saves an iteration.
	if a != b {
		for i := 0; i < len(p.id); i++ {
			if p.id[i] == a {
				p.id[i] = b
			}
		}
	}
}

func (p quickFind) FindSet(x int) int {
	return p.id[x]
}

func (p quickFind) Connected(x, y int) bool {
	return p.FindSet(x) == p.FindSet(y)
}

func (p *quickFind) Capacity() int {
	return p.capacity
}

// Determines the number of disjoint sets in the partition.
func (p *quickFind) Count(countIndividuals bool) int {
	roots := make(map[int]int)
	for i := 0; i < p.capacity; i++ {
		if _, ok := p.seen[i]; p.id[i] == i && (countIndividuals || ok) {
			roots[i] = i
		}
	}
	return len(roots)
}
