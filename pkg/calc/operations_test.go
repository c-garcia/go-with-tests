package calc

import (
	"testing"
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
