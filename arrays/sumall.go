package array

func SumAll(nums ...[]int) (sums []int) {
	for _, num := range nums {
		// mySlice[10] = 1 will get a runtime error if initial capacity is 1
		// but `append` with extend the slice
		sums = append(sums, Sum(num))
	}
	return
}

func SumAllTails(nums ...[]int) (sums []int) {
	for _, num := range nums {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			tail := num[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
