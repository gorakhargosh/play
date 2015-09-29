package partition

// weightedQuickUnion is a weighted quick-union disjoint set partitioner.
type weightedQuickUnion struct {
	id     []int
	weight []int
}

// NewWeightedQuickUnion generates a new partition.
func NewWeightedQuickUnion(size int) Partition {
	p := &weightedQuickUnion{
		id:     make([]int, size),
		weight: make([]int, size),
	}
	for i := 0; i < size; i++ {
		p.id[i] = i     // Each node starts partitioned.
		p.weight[i] = 1 // Each tree starts at weight 1.
	}
	return p
}

func (p *weightedQuickUnion) Union(x, y int) {
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

func (p weightedQuickUnion) FindSet(x int) int {
	for x != p.id[x] {
		x = p.id[x]
	}
	return x
}

func (p weightedQuickUnion) Connected(x, y int) bool {
	return p.FindSet(x) == p.FindSet(y)
}
