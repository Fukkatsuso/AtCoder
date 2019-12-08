package main

import (
	"fmt"
	"sort"
)

type myInt64 []int64

func (a myInt64) Len() int {
	return len(a)
}

func (a myInt64) Less(i, j int) bool {
	return a[i] > a[j]
}

func (a myInt64) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var x, y, z, k int
	fmt.Scan(&x, &y, &z, &k)
	a := make(myInt64, x)
	b := make(myInt64, y)
	c := make(myInt64, z)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := range b {
		fmt.Scan(&b[i])
	}
	for i := range c {
		fmt.Scan(&c[i])
	}

	sum := make(myInt64, x*y)
	// a, bの組み合わせを考える
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			sum[i*y+j] = a[i] + b[j]
		}
	}
	sort.Sort(sum)
	sort.Sort(c)

	// sum[0:k], c[0:z]の組み合わせを考える
	ans := make(myInt64, k*z)
	iMax := min(k, x*y)
	for i := 0; i < iMax; i++ {
		for j := 0; j < z; j++ {
			ans[i*z+j] = sum[i] + c[j]
		}
	}
	sort.Sort(ans)

	for _, v := range ans[:k] {
		fmt.Println(v)
	}
}
