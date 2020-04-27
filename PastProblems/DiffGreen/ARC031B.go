package main

import (
	"bufio"
	"fmt"
	"os"
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

// 島が1つだけか?
func isSingle(n int, a [][]byte, iBegin, jBegin int) bool {
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// [(i, j)]
	q := make([][2]int, 0)
	q = append(q, [2]int{iBegin, jBegin})
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	// bfs
	for len(q) > 0 {
		x, y := q[0][1], q[0][0]
		visited[y][x] = true
		for i := 0; i < 4; i++ {
			nextX, nextY := x+dx[i], y+dy[i]
			if nextX < 0 || n <= nextX || nextY < 0 || n <= nextY {
				continue
			}
			if a[nextY][nextX] == 'o' && !visited[nextY][nextX] {
				q = append(q, [2]int{nextY, nextX})
			}
		}
		q = q[1:]
	}

	// 陸地を全て訪れたか?
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 'o' && !visited[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := 10
	a := make([][]byte, n)
	for i := range a {
		a[i] = []byte(next())
	}

	// 左上から探して最初に'o'が見つかる場所
	var iBegin, jBegin int
	for iBegin, jBegin = 0, 0; a[iBegin][jBegin] != 'o'; iBegin, jBegin = iBegin+(jBegin+1)/n, (jBegin+1)%n {
	}

	ok := false
	for i := 0; i < n && !ok; i++ {
		for j := 0; j < n && !ok; j++ {
			if a[i][j] == 'o' {
				continue
			}
			// a[i][j]を埋め立てて，1つの島になるか試す
			a[i][j] = 'o'
			ok = isSingle(n, a, iBegin, jBegin)
			// 元に戻す
			a[i][j] = 'x'
		}
	}

	if ok {
		puts("YES")
	} else {
		puts("NO")
	}
}
