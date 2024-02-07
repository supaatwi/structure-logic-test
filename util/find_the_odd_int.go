package util

import "sort"

func FindTheOddInt(nums []int) int {
	temp := make(map[int]int)
	for _, num := range nums {
		temp[num]++
	}

	mx := make([][]int, len(temp))
	i := 0
	for n, c := range temp {
		mx[i] = []int{n, c}
		i++
	}
	sort.Slice(mx, func(i, j int) bool {
		return mx[i][1] < mx[j][1]
	})

	for _, m := range mx {
		if m[1]%2 != 0 {
			return m[0]
		}
	}
	return -1
}
