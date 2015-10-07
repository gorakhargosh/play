package partition

// forest is a path-compressed weighted quick-union disjoint set partitioner.
type forest struct {
	// Stores the node ids where each node id initially starts out as pointing at
	// itself. As more union operations are performed on the dataset, nodes point
	// at the root as their representative element of their disjoint set.
	id []int

	// Stores the weights for node ids that have been seen.
	weight []uint

	// The presence of a node id in this map indicates that
	// the element has been previously seen in a union operation.
	seen map[int]bool

	// The capacity of the partition.
	capacity int
}

// NewForestPartition generates a new partition.
func NewForestPartition(size int) Partition {
	p := &forest{
		id:       make([]int, size),
		weight:   make([]uint, size),
		seen:     make(map[int]bool),
		capacity: size,
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
// recursive point-all-nodes-at-root path compression when unwinding the
// recursion stack.
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

	if a != b { // equivalent to !p.Connected(a, b)
		// We perform a union only if the two sets are not already connected. A
		// union of two disjoint sets without this check would result in an
		// unbalanced tree.

		// Mark these elements as seen when performing their union.
		p.seen[a] = true
		p.seen[b] = true

		if p.weight[a] < p.weight[b] {
			p.id[a] = b
			p.weight[b] += p.weight[a]
			// Reduce the number of disjoint sets by 1.
		} else {
			p.id[b] = a
			p.weight[a] += p.weight[b]
			// Reduce the number of disjoint sets by 1.
		}
	}
}

func (p *forest) Weight(x int) uint {
	return p.weight[p.FindSet(x)]
}

func (p *forest) MinWeight(countIndividuals bool) uint {
	weight := uint(0)
	minWeight := uint(0)
	for i := 0; i < p.capacity; i++ {
		if _, ok := p.seen[i]; p.id[i] == i && (ok || countIndividuals) {
			// We have a root element.
			weight = p.weight[i]
			if minWeight == 0 || weight < minWeight {
				minWeight = weight
			}
		}
	}
	return weight
}

func (p *forest) MaxWeight(countIndividuals bool) uint {
	weight := uint(0)
	maxWeight := uint(0)
	for i := 0; i < p.capacity; i++ {
		if _, ok := p.seen[i]; p.id[i] == i && (ok || countIndividuals) {
			weight = p.weight[i]
			if maxWeight == 0 || weight > maxWeight {
				maxWeight = weight
			}
		}
	}
	return weight
}

func (p *forest) Capacity() int {
	return p.capacity
}

// Determines the number of disjoint sets in the partition.
func (p *forest) Count(countIndividuals bool) int {
	roots := make(map[int]int)
	for i := 0; i < p.capacity; i++ {
		if _, ok := p.seen[i]; p.id[i] == i && (countIndividuals || ok) {
			roots[i] = i
		}
	}
	return len(roots)
}
