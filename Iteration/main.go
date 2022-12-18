package iteration

func Repeat(character string, vezes int) string {
	var repeated string
	for i := 0; i < vezes; i++ {
		repeated += character
	}

	return repeated
}
