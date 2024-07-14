package arraysandslices

// Given a slice of integers, add up each elements are return the total.
func Sum(numbers []int) (total int) {
	for _, number := range numbers { // The `_` represents a blank identifier.
		total += number
	}
	return
}

// Given a slice of integer slices, add up each number in the slice, and return a new slice containing the totals at that index.
func SumAll(slices ...[]int) (sums []int) {
	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}
	return
}
