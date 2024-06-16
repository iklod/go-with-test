package arrslice

func Sum(arr []int) int {
	var sum int
	for _, number := range arr {
		sum += number
	}
	return sum
}

func SumAll(arrToSum ...[]int) []int {
	// lengthOfNumbers := len(arrToSum)
	// sums := make([]int, lengthOfNumbers)
	var sums []int
	for _, numbers := range arrToSum {
		// sums[i] = Sum(numbers)
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(arrToSum ...[]int) []int {
	var sums []int
	for _, numbers := range arrToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}

	return sums
}
