package arraysandslices

func Sum(numbers []int) (total int) {
	for _, number := range numbers { // The `_` represents a blank identifier.
		total += number
	}
	return
}
