package main

import (
	"fmt"
)

func binarySearch(a []int, key int) int {
	for l, r := 0, len(a); l < r; {
		mid := (l + r) / 2
		if a[mid] == key {
			return mid
		}
		if key < a[mid] {
			r = mid
		} else if a[mid] < key {
			l = mid + 1
		}
	}
	return -1
}

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

func count(a []int, key int) int {
	return upperBound(a, key) - lowerBound(a, key)
}

func main() {
	a := []int{1, 2, 2, 2, 3, 4, 4, 5, 5, 6, 7, 8}
	fmt.Println(binarySearch(a, 3))
	fmt.Println(lowerBound(a, 2))
	fmt.Println(upperBound(a, 2))
	fmt.Println(count(a, 2))
}
