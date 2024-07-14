package iteration

// Given a character, repeat it n times by the count argument given.
func Repeat(character string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += character
	}

	return
}
