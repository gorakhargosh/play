package partition

// quickFind is a quick-find slow-union partitioner.
type quickFind struct {
}

// NewQuickFind creates a new partition.
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
