package main

import "fmt"

func lowerBound(a []int, key int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := (l + r) / 2
		if a[mid] < key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func upperBound(a []int, key int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := (l + r) / 2
		if a[mid] <= key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

// 最長増加部分列の長さを返す
// strict: trueなら狭義，falseなら広義
func LISLen(a []int, strict bool) int {
	t := make([]int, 0)
	for i := range a {
		var idx int
		if strict {
			idx = lowerBound(t, a[i])
		} else {
			idx = upperBound(t, a[i])
		}
		if idx == len(t) {
			t = append(t, a[i])
		} else {
			t[idx] = a[i]
		}
	}
	return len(t)
}

func main() {
	a := []int{1, 3, 5, 2, 4, 6, 6, 7}

	fmt.Println(LISLen(a, true))
	fmt.Println(LISLen(a, false))
}
