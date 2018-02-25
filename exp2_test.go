package main

import (
	"testing"
	"math"
)

func uint64exp2(p uint64) (uint64) {
	if p > 63 {
		panic("power cannot be higher than 63 for integer exp2")
	}

	return uint64(1) << p;
}

func float64exp2(p uint64) (uint64) {

	return uint64(math.Exp2(float64(p)));
}

// test math
func TestExp(t *testing.T) {
	for i := 0; i < 1000; i ++ {
		p := uint64(i % 64)

		if f, i := math.Exp2(float64(p)) , float64(uint64exp2(p)); f != i{
			t.Fatalf("expected %v got %v", f, i)
		}
	}
}

func BenchmarkFloat0(b *testing.B) {
	benchmarkFloat(0, b)
}

func BenchmarkFloat4(b *testing.B) {
	benchmarkFloat(4, b)
}

func BenchmarkFloat16(b *testing.B) {
	benchmarkFloat(16, b)
}

func BenchmarkFloat32(b *testing.B) {
	benchmarkFloat(32, b)
}

func BenchmarkFloat63(b *testing.B) {
	benchmarkFloat(63, b)
}

func BenchmarkUint0(b *testing.B) {
	benchmarkUint(0, b)
}

func BenchmarkUint4(b *testing.B) {
	benchmarkUint(4, b)
}

func BenchmarkUint16(b *testing.B) {
	benchmarkUint(16, b)
}

func BenchmarkUint32(b *testing.B) {
	benchmarkUint(32, b)
}

func BenchmarkUint63(b *testing.B) {
	benchmarkUint(63, b)
}

// actual calls
func benchmarkFloat(i uint64, b*testing.B) {
	for n := 0; n < b.N; n++ {
		float64exp2(i)
	}
}

func benchmarkUint(i uint64, b*testing.B) {
	for n := 0; n < b.N; n++ {
		uint64exp2(i)
	}
}

