package popcount

import "testing"

func TestPopCount(t *testing.T) {
	count := PopCount(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCount(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCount(0x123456789ABCDE)
	if count != 28 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func TestPopCountLoop(t *testing.T) {
	count := PopCountLoop(3)
	if count != 2 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountLoop(16)
	if count != 1 {
		t.Errorf("popcount number is differnt %d", count)
	}
	count = PopCountLoop(0x123456789ABCDE)
	if count != 28 {
		t.Errorf("popcount number is differnt %d", count)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(uint64(i))
	}
}
