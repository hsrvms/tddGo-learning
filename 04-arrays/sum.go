package arrays

func Sum(numbers []int) int {
	var total int
	for _, n := range numbers {
		total += n
	}

	return total
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums[i] = 0
		} else {
			tail := numbers[1:]
			sums[i] = Sum(tail)
		}
	}
	return sums
}
