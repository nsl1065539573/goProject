package sum

func SumAll(items ...[]int) []int {
	var res []int
	for _, item := range items {
		res = append(res, Sum(item))
	}
	return res
}

func Sum(numbers []int) int {
	var res int
	for _, item := range numbers {
		res += item
	}
	return res
}
