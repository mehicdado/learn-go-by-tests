package iteration

//Repeat given string 5 times as a result.
func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}

	return repeated
}
