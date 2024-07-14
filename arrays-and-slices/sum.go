package arraysandslices

// Given a slice of integers, add up each elements are return the total.
func Sum(numbers []int) (total int) {
	for _, number := range numbers { // The `_` represents a blank identifier.
		total += number
	}
	return
}

// Given a slice of integer slices, add up each number in the slice, and return a new slice containing the total at that index.
func SumAll(slices ...[]int) (sums []int) {
	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}
	return
}

// Given a slice of integer slices, add up each number in the slice omitting the first index, and return a new slice containing the total at that index.
func SumAllTails(slices ...[]int) (sums []int) {
	for _, numbers := range slices {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // slice[1:] will create a new slice starting from index 1, effectively removing the first element.
			sums = append(sums, Sum(tail))
		}
	}
	return
}
