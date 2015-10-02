package partition

// forest is a path-compressed weighted quick-union disjoint set partitioner.
type forest struct {
	// Stores the node ids where each node id initially starts out as pointing at
	// itself. As more union operations are performed on the dataset, nodes point
	// at the root as their representative element of their disjoint set.
	id []int

	// Stores the weights for node ids that have been seen.
	weight []uint64

	// The presence of a node id in this map indicates that
	// the element has been previously seen in a union operation.
	seen map[int]bool
}

// NewForestPartition generates a new partition.
func NewForestPartition(size int) Partition {
	p := &forest{
		id:     make([]int, size),
		weight: make([]uint64, size),
		seen:   make(map[int]bool),
	}
	for i := 0; i < size; i++ {
		p.id[i] = i
		p.weight[i] = 1
	}
	return p
}

// findRoot determines the root of the set while performing a lazy, one-pass
// point-at-grandparent path compression.
func (p *forest) findRoot(x int) int {
	for x != p.id[x] {
		p.id[x] = p.id[p.id[x]]
		x = p.id[x]
	}
	return x
}

// findRoot2 determines the root of the set while performing an eager, two-pass
// iterative point-all-nodes-at-root path compression.
func (p *forest) findRoot2(x int) int {
	// Two pass variant that sets root for all traversed elements.
	i := x
	for x != p.id[x] {
		x = p.id[x]
	}
	// x is now the root.
	for i != p.id[i] {
		i = p.id[i]
		p.id[i] = x
	}
	return x
}

// findRoot2R determines the root of the set while performing an eager, two-pass
// recursive point-all-nodes-at-root path compression.
func (p *forest) findRoot2R(x int) int {
	if x != p.id[x] {
		p.id[x] = p.findRoot2R(p.id[x])
	}
	return p.id[x]
}

// Finds the representative element of the set to which x belongs.
func (p *forest) FindSet(x int) int {
	return p.findRoot2(x)
}

// Connected determines whether two elements are in the same set.
func (p forest) Connected(x, y int) bool {
	return p.FindSet(x) == p.FindSet(y)
}

// Union combines the sets represented by x and y.
func (p *forest) Union(x, y int) {
	a := p.FindSet(x)
	b := p.FindSet(y)
	if p.weight[a] < p.weight[b] {
		p.id[a] = b
		p.weight[b] += p.weight[a]
	} else {
		p.id[b] = a
		p.weight[a] += p.weight[b]
	}
}
