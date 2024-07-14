package arraysandslices

func Sum(numbers [5]int) (total int) {
	for _, number := range numbers {
		total += number
	}
	return
}
