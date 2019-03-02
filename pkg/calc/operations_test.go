package calc

import "testing"

// * Testing double
func TestDouble(t *testing.T) {
	if 4 != Double(2) {
		t.Fatal("4 should be = Double(2)")
	}
}

// * Testing sum
func TestSum(t *testing.T) {
	t.Fatal("Test not implemented")
}
