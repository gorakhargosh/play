package partition

import "testing"

const (
	N = 10
)

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

func TestLinearConnectivity(t *testing.T) {
	checkConnectivity(NewLinearPartition(N), t)
}

func TestForestPartitionConnectivity(t *testing.T) {
	checkConnectivity(NewForestPartition(N), t)
}

func benchmarkConnectivity(p Partition, b *testing.B) {
	p.Union(4, 3)
	p.Union(3, 8)
	p.Union(6, 5)
	p.Union(9, 4)
	p.Union(2, 1)
	p.Union(5, 0)
	p.Union(7, 2)
	p.Union(6, 1)
	p.Union(1, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Connected(0, 7)
	}
}

func BenchmarkLinearPartition(b *testing.B) {
	benchmarkConnectivity(NewLinearPartition(N), b)
}

func BenchmarkForestPartition(b *testing.B) {
	benchmarkConnectivity(NewForestPartition(N), b)
}
