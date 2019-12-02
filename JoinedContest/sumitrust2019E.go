package main

import (
	"fmt"
)

func main() {
	const mod int = 1000000007
	var n int
	var a int
	fmt.Scan(&n)

	ans := 1
	table := make([][]int, n+1)
	table[0] = make([]int, 3)
	for i := 1; i <= n; i++ {
		table[i] = make([]int, 3)
		fmt.Scan(&a)
		cnt := 0
		for j := 0; j < 3; j++ {
			table[i][j] = table[i-1][j]
			if a == table[i-1][j] {
				cnt++
				if cnt == 1 {
					table[i][j]++
				}
			}
		}
		ans = (ans * cnt) % mod
	}

	fmt.Println(ans)
}
