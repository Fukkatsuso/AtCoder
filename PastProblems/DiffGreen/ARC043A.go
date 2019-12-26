package main

import (
	"fmt"
)

func inputIntSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		fmt.Scan(&slice[i])
	}
	return slice
}

func sum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	s := inputIntSlice(n)

	max, min, sum := s[0], s[0], int64(s[0])
	for i := 1; i < n; i++ {
		if max < s[i] {
			max = s[i]
		}
		if min > s[i] {
			min = s[i]
		}
		sum += int64(s[i])
	}

	if max-min == 0 && b != 0 {
		fmt.Println(-1)
		return
	}

	p := float64(b) / float64(max-min)
	q := float64(a) - p*(float64(sum)/float64(n))
	fmt.Println(p, q)
}
