// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

// 先手の得点
func score(t, b, c [][]int) int {
	s := 0
	for i := 0; i < 3-1; i++ {
		for j := 0; j < 3; j++ {
			if t[i][j] == t[i+1][j] {
				s += b[i][j]
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3-1; j++ {
			if t[i][j] == t[i][j+1] {
				s += c[i][j]
			}
		}
	}
	return s
}

func dfs(t, b, c [][]int, turn int) int {
	done := true
	for k := 0; k < 9; k++ {
		i, j := k/3, k%3
		done = done && t[i][j] > 0
	}
	if done {
		return score(t, b, c)
	}
	s := (turn-1)*10000 + (turn-2)*10000
	for k := 0; k < 9; k++ {
		i, j := k/3, k%3
		if t[i][j] > 0 {
			continue
		}
		t[i][j] = turn
		if turn == 1 {
			s = max(s, dfs(t, b, c, 2))
		} else {
			s = min(s, dfs(t, b, c, 1))
		}
		t[i][j] = 0
	}
	return s
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	b, c := make([][]int, 2), make([][]int, 3)
	mx := 0
	for i := range b {
		b[i] = make([]int, 3)
		for j := range b[i] {
			b[i][j] = nextInt()
			mx += b[i][j]
		}
	}
	for i := range c {
		c[i] = make([]int, 2)
		for j := range c[i] {
			c[i][j] = nextInt()
			mx += c[i][j]
		}
	}

	t := make([][]int, 3)
	for i := range t {
		t[i] = make([]int, 3)
	}
	choku := dfs(t, b, c, 1)
	puts(choku)
	puts(mx - choku)
}
