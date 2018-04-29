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

//把数据中除了第一个数的数相加
func SumAllTails(nums ...[]int) []int {
	var sums []int

	for _, v := range nums {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tails := v[1:]
			sums = append(sums, Sum(tails))
		}
	}

	return sums
}
