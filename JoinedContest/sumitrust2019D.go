package main

import (
	"fmt"
)

// x^y
func pow(x, y int) int {
	ret := 1
	for i := 0; i < y; i++ {
		ret *= x
	}
	return ret
}

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	ans := 0
	for i := 0; i < 1000; i++ {
		can := true
		digit := 0 // ラッキーナンバーのインデックス

		for j := 2; j >= 0 && can; j-- { // iの調査する桁数
			checkNum := (i / pow(10, j)) % 10
			for ; ; digit++ {
				if digit == n {
					can = false
					break
				}
				if checkNum == int(s[digit]-48) {
					digit++
					break
				}
			}
		}

		if can {
			ans++
		}
	}

	fmt.Println(ans)
}
