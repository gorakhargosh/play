package partition

type quickFind struct {
}

// NewQuickFind generates a quick-find slow union partitioner.
func NewQuickFind() Partition {
	return &quickFind{}
}

func (q *quickFind) Union(x, int int) {
}

func (q quickFind) FindSet(x int) int {
	return 0
}

func (q quickFind) Connected(x, y int) bool {
	return false
}
