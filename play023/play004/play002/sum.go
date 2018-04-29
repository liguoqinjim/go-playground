package sum

func Sum(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}

func SumAll(nums ...[]int) []int {
	var sums []int

	for _, v := range nums {
		sums = append(sums, Sum(v))
	}

	return sums
}
