package partition

import "testing"

func checkConnectivity(p Partition, t *testing.T) {
	p.Union(4, 3)
	p.Union(3, 8)
	p.Union(6, 5)
	p.Union(9, 4)
	p.Union(2, 1)
	if p.Connected(0, 7) == true {
		t.Errorf("error: misconnection: %d-%d", 0, 7)
	}
	if p.Connected(8, 9) != true {
		t.Errorf("error: not connected: %d-%d", 8, 9)
	}
	p.Union(5, 0)
	p.Union(7, 2)
	p.Union(6, 1)
	p.Union(1, 0)
	// Retest previous connection now.
	if p.Connected(0, 7) != true {
		t.Errorf("error: not connected: %d-%d", 0, 7)
	}
}

func TestConnectivity(t *testing.T) {
	checkConnectivity(NewQuickFind(), t)
	checkConnectivity(NewQuickUnion(), t)
	checkConnectivity(NewWeightedQuickUnion(), t)
	checkConnectivity(NewPartition(), t)
}

func createNetwork(p Partition) {
	p.Union(4, 3)
	p.Union(3, 8)
	p.Union(6, 5)
	p.Union(9, 4)
	p.Union(2, 1)
	p.Union(5, 0)
	p.Union(7, 2)
	p.Union(6, 1)
	p.Union(1, 0)
}

func BenchmarkQuickFind(b *testing.B) {

	b.ResetTimer()
}
