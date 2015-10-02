package partition

// quickFind is a quick-find slow-union partitioner.
type quickFind struct {
	id       []int
	capacity int
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
	xid := p.FindSet(x)
	yid := p.FindSet(y)
	for i := 0; i < len(p.id); i++ {
		if p.id[i] == xid {
			p.id[i] = yid
		}
	}
}

func (p quickFind) FindSet(x int) int {
	return p.id[x]
}

func (p quickFind) Connected(x, y int) bool {
	return p.FindSet(x) == p.FindSet(y)
}
