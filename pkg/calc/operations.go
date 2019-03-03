package calc

// Double returns the double of an int
func Double(x int) int {
	return 2 * x
}

// Sum returns the sum of x and y
func Sum(x, y int) int {
	return x + y
}

// GCD returns the great common divisor of a and b.
// Using the Euclides algorithm
//     GCD(3, 9) == 3
func GCD(x, y int) int {

	for x != y {
		if x > y {
			x = x - y
		} else {
			y = y - x
		}
	}
	return x
}
