package main

import (
	"fmt"
	"sort"
)

func inputIntSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		fmt.Scan(&slice[i])
	}
	return slice
}

func include(arr []int, x int) bool {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == x {
			return true
		}
		if arr[mid] > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	var na, nb int
	fmt.Scan(&na, &nb)
	a := inputIntSlice(na)
	b := inputIntSlice(nb)
	sort.Sort(sort.IntSlice(b))

	cnt := 0
	for i := range a {
		if include(b, a[i]) {
			cnt++
		}
	}

	fmt.Println(float64(cnt) / float64(na+nb-cnt))
}
