package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// * Testing double
func TestDouble(t *testing.T) {
	if 4 != Double(2) {
		t.Fatal("4 should be = Double(2)")
	}
}

// * Testing sum
func TestSum(t *testing.T) {
	if 2 != Sum(0, 2) {
		t.Fatal("2 + 0 should be 2")
	}
}

// * Testing Diff
func TestDiff(t *testing.T) {
	if 3 != Diff(8, 5) {
		t.Fatal("8 - 5 should be 3 ")
	}
}

// * Testing GCD
func TestGCD(t *testing.T) {
	if 4 != GCD(8, 12) {
		t.Fatal("GCD(8,12) should be 4")
	}
}

// * Benchmark for GCD
func BenchmarkGCD(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GCD(39916801, 479001599)
	}
}

// * Testing for GCDIt
func TestGCDIt(t *testing.T) {
	if 4 != GCDIt(8, 12) {
		t.Fatal("GCDIt(8,12) should be 4")
	}
}

// * Benchmark for GCD
func BenchmarkGCDIt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GCDIt(39916801, 479001599)
	}
}

// * Test for LCM
func TestLCM(t *testing.T) {
	assert.Equal(t, 60, LCM(12, 15))
}
