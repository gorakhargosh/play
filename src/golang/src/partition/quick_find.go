package partition

// quickFind is a quick-find slow-union partitioner.
type quickFind struct {
	id []int
}

// NewQuickFind creates a new partition.
func NewQuickFind(size int) Partition {
	return &quickFind{
		id: make([]int, size),
	}
}

func (q *quickFind) Union(x, int int) {
}

func (q quickFind) FindSet(x int) int {
	return 0
}

func (q quickFind) Connected(x, y int) bool {
	return false
}
