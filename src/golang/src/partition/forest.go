package partition

// forest is a path-compressed weighted quick-union disjoint set partitioner.
type forest struct {
	id     []int
	weight []int
}

// NewForestPartition generates a new partition.
func NewForestPartition(size int) Partition {
	p := &forest{
		id:     make([]int, size),
		weight: make([]int, size),
	}
	for i := 0; i < size; i++ {
		p.id[i] = i
		p.weight[i] = 1
	}
	return p
}

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

func (p *forest) FindSet(x int) int {
	return p.findRoot2(x)
}

func (p forest) Connected(x, y int) bool {
	return p.FindSet(x) == p.FindSet(y)
}
