package partition

// Partition is an equivalence relation over disjoint sets.
type Partition interface {
	Union(x, y int)
	FindSet(x int) int
	Connected(x, y int) bool
}
