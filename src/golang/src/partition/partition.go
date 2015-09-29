package partition

// Partition is an equivalence relation over disjoint sets.
type Partition interface {
	// Union performs a union of the sets that x and y belong to.
	Union(x, y int)
	// FindSet finds the representative element of the set that x belongs to.
	FindSet(x int) int
	// Connected determines whether two elements, x and y, are connected.
	Connected(x, y int) bool
}
