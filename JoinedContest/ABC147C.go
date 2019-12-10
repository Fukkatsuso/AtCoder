// コンテスト後
// AC

package main

import (
	"fmt"
)

func honest(n int, i uint) int {
	return (n >> i) & 1
}

func honestsNum(n int) int {
	ret := 0
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			ret++
		}
		n >>= 1
	}
	return ret
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	var n uint
	fmt.Scan(&n)
	a := make([]int, n)
	x := make([][]int, n)
	y := make([][]int, n)
	for i := uint(0); i < n; i++ {
		fmt.Scan(&a[i])
		x[i] = make([]int, a[i])
		y[i] = make([]int, a[i])
		for j := 0; j < a[i]; j++ {
			fmt.Scan(&x[i][j], &y[i][j])
			x[i][j]--
		}
	}

	// bit全探索
	iMax := 1 << n
	ans := 0 // i = 0のとき
	for i := 1; i < iMax; i++ {
		ok := true
		for j := uint(0); j < n; j++ {
			if honest(i, j) == 0 {
				continue
			}
			for k := 0; k < a[j]; k++ {
				if honest(i, uint(x[j][k])) != y[j][k] {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			ans = max(ans, honestsNum(i))
		}
	}

	fmt.Println(ans)
}
