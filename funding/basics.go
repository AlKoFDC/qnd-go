package funding

// A function can have multiple return parameters.
func swap(a, b string) (string, string) {
	return b, a
}

// Return parameters can be named.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	// Type conversion
	i = 7
	f = float64(i)
	// Type assertion
	n = interface{}("golang")
	s = n.(string)
)
