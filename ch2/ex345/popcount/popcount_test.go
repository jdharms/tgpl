package popcount

import (
	"testing"
)

// BenchmarkInitPopCount benchmarks the PopCount algorithm that uses a pre-initialized lookup table.
func BenchmarkInitPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop(uint64(i))
	}
}

func BenchmarkShiftPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Shift64(uint64(i))
	}
}

func BenchmarkShiftOffCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftOff(uint64(i))
	}
}

func TestShift(t *testing.T) {
	for i := uint64(0); i < 1000000; i++ {
		if Shift64(i) != PopCount(i) {
			t.Errorf("Shift popcount was incorrect on value: %v\n", i)
		}
	}
}

func TestShiftOff(t *testing.T) {
	for i := uint64(0); i < 1000000; i++ {
		if ShiftOff(i) != Loop(i) {
			t.Errorf("Shift popcount was incorrect on value: %v\n", i)
		}
	}
}
