package calc

import (
	"fmt"
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

// * Testing GCD
func TestGCD(t *testing.T) {
	if 3 != GCD(3, 9) {
		t.Fatal("GCD(3,9) should be 3")
	}
}

func ExampleGCD() {
	fmt.Println(GCD(3, 9))
	// Output:3
}

func ExampleGCD_coprime() {
	fmt.Println(CodingGCD(2, 9))
	// Output: 1
}
