package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	num := 1
	ans := 0
	for i := 0; i < n; i++ {
		if a[i] == num {
			num++ // 次の数字の探索に移る
		} else {
			ans++ // 砕く
		}
	}
	if num == 1 {
		ans = -1
	}
	fmt.Println(ans)
}
