// この方法だと隣接点以外の頂点に同じ色を塗ることになるのでNG

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	d := make(vArray, n) //頂点の次数
	for i := 0; i < n; i++ {
		d[i] = make([]int, 2)
		d[i][0] = i
	}
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a[i], &b[i])
		d[a[i]-1][1]++
		d[b[i]-1][1]++
	}

	// 次数の降順に並べる
	sort.Sort(d)

	col := 1
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		// 塗られていたらスキップ
		if ans[d[i][0]] > 0 {
			continue
		}
		ans[d[i][0]] = col
		// for {
		// 	if ans[] == 0 {
		// 		ans[] = col
		// 	}
		// }
		col++
	}
}

type vArray [][]int

func (va vArray) Len() int {
	return len(va)
}

func (va vArray) Swap(i, j int) {
	va[i][0], va[j][0] = va[j][0], va[i][0]
	va[i][1], va[j][1] = va[j][1], va[i][1]
}

func (va vArray) Less(i, j int) bool {
	return va[i][1] > va[j][1]
}
