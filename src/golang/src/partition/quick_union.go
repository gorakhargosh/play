package partition

type quickUnion struct {
}

// NewQuickUnion generates a quick-find slow union partitioner.
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
