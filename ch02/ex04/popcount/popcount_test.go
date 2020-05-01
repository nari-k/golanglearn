package popcount

import "testing"

func TestPopCount(t *testing.T) {
	count := PopCount(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCount(0x123456789ABCDE)
	if count != 28 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func TestPopCountLoop64(t *testing.T) {
	count := PopCountLoop64(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountLoop64(0x123456789ABCDE)
	if count != 28 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
func BenchmarkPopCountLoop64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop64(uint64(i))
	}
}
