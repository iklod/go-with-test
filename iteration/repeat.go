package iteration

// Repeat a value n times
func Repeat(value string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += value
	}
	return repeated
}
