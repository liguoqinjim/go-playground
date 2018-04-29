package sum

func Sum(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}

func SumAll(nums ...[]int) []int {
	sums := make([]int, len(nums))

	for n, v := range nums {
		sums[n] = Sum(v)
	}

	return sums
}
