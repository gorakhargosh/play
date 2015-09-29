package partition

// quickUnion is a quick-union disjoint set partitioner
type quickUnion struct {
	id []int
}

// NewQuickUnion creates a new partition.
func NewQuickUnion(size int) Partition {
	p := &quickUnion{
		id: make([]int, size),
	}
	for i := 0; i < size; i++ {
		p.id[i] = i
	}
	return p
}

func (p *quickUnion) Union(x, y int) {
}

func (p quickUnion) FindSet(x int) int {
	return 0
}

func (p quickUnion) Connected(x, y int) bool {
	return false
}
