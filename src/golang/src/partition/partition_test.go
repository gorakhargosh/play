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

func TestQuickFindConnectivity(t *testing.T) {
	checkConnectivity(NewQuickFind(N), t)
}

func TestQuickUnionConnectivity(t *testing.T) {
	checkConnectivity(NewQuickUnion(N), t)
}

func TestWeightedQuickUnionConnectivity(t *testing.T) {
	checkConnectivity(NewWeightedQuickUnion(N), t)
}

func TestPathCompressedPartitionConnectivity(t *testing.T) {
	checkConnectivity(NewPathCompressedPartition(N), t)
}

func TestPathShortenedPartitionConnectivity(t *testing.T) {
	checkConnectivity(NewPathShortenedPartition(N), t)
}

func TestBSTConnectivity(t *testing.T) {
	checkConnectivity(NewBSTPartition(N), t)
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

func BenchmarkQuickFind(b *testing.B) {
	benchmarkConnectivity(NewQuickFind(N), b)
}

func BenchmarkQuickUnion(b *testing.B) {
	benchmarkConnectivity(NewQuickUnion(N), b)
}

func BenchmarkWeightedQuickUnion(b *testing.B) {
	benchmarkConnectivity(NewWeightedQuickUnion(N), b)
}

func BenchmarkPathCompressedPartition(b *testing.B) {
	benchmarkConnectivity(NewPathCompressedPartition(N), b)
}

func BenchmarkPathShortenedPartition(b *testing.B) {
	benchmarkConnectivity(NewPathShortenedPartition(N), b)
}

func BenchmarkBSTPartition(b *testing.B) {
	benchmarkConnectivity(NewBSTPartition(N), b)
}
