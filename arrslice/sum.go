package arrslice

func Sum(arr []int) int {
	var sum int
	for _, number := range arr {
		sum += number
	}
	return sum
}
