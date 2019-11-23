// x

package main

import (
	"fmt"
)

func main() {
	var h, w, k int
	fmt.Scan(&h, &w, &k)
	s := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&s[i])
	}

	ans := make([][]int, h)
	for i := 0; i < h; i++ {
		ans[i] = make([]int, w)
		for j := 0; j < w; j++ {
			ans[i][j] = -1
		}
	}

	num := 1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				dfs(s, num, i, j, ans)
				num++
			}
		}
	}
}

func dfs(s []string, num, i, j int, ans [][]int) {
	ans[i][j] = num
	// 横いっぱい広げる
	if j+1 < len(s) && ans[i][j+1] == -1 && s[i][j+1] != '#' {
		dfs(s, num, i, j+1, ans)
	}
	if 0 <= j-1 && ans[i][j-1] == -1 && s[i][j-1] != '#' {
		dfs(s, num, i, j-1, ans)
	}
	// 縦いっぱい広げる
}
