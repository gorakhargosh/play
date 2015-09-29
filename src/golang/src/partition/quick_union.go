package partition

// quickUnion is a quick-union disjoint set partitioner
type quickUnion struct {
}

// NewQuickUnion creates a new partition.
func NewQuickUnion() Partition {
	return &quickUnion{}
}

func (q *quickUnion) Union(x, int int) {
}

func (q quickUnion) FindSet(x int) int {
	return 0
}

func (q quickUnion) Connected(x, y int) bool {
	return false
}
