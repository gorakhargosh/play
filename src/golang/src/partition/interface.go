package partition

// Partition is an equivalence relation over disjoint sets.
type Partition interface {
	// Union performs a union of the sets that x and y belong to.
	Union(x, y int)
	// FindSet finds the representative element of the set that x belongs to.
	FindSet(x int) int
	// Connected determines whether two elements, x and y, are connected.
	Connected(x, y int) bool

	// Determines the weight of the disjoint set represented by x.
	Weight(x int) uint

	// Calculates the weight of the smallest disjoint set within the partition.
	MinWeight() uint

	// Calculates the weight of the largest disjoint set within the partition.
	MaxWeight() uint

	// Determines the capacity of the partition.
	Capacity() int

	// Determines the number of disjoint sets that have been seen in the
	// partition. If you want to determine the number of proper disjoint sets,
	// you can use Capacity() - Count().
	CountSeen() int
}
