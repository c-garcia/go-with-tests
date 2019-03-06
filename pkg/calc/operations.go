package calc

// Double returns the double of an int
func Double(x int) int {
	return 2 * x
}

// Sum returns the sum of x and y
func Sum(x, y int) int {
	return x + y
}

// Diff returns the difference between x and y
func Diff(x, y int) int {
	return x - y
}

// GCD returns the Greatest Common Divisor of x and y
func GCD(x, y int) int {
	if x == y {
		return x
	} else if x > y {
		return GCD(x-y, y)
	} else {
		return GCD(x, y-x)
	}
}

// GCD returns the Greatest Common Divisor of x and y (More efficient?)
func GCDIt(x, y int) int {
	for x != y {
		if x > y {
			x = x - y
		} else {
			y = y - x
		}
	}
	return x
}

// LCM returns the Least Common Multiple of x and y
func LCM(x, y int) int {
	return x * y / GCD(x, y)
}
