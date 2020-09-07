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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w := nextInt(), nextInt()
	s := make([]string, h)
	for i := range s {
		s[i] = next()
	}

	checked := make([][]bool, h)
	for i := range checked {
		checked[i] = make([]bool, w)
	}
	di := []int{0, 1, 0, -1}
	dj := []int{1, 0, -1, 0}
	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if checked[i][j] {
				continue
			}

			black, white := 0, 0
			// BFS
			que := make([][2]int, 0)
			que = append(que, [2]int{i, j})
			for len(que) > 0 {
				I, J := que[0][0], que[0][1]
				que = que[1:]
				if checked[I][J] {
					continue
				}
				checked[I][J] = true
				if s[I][J] == '#' {
					black++
				} else {
					white++
				}
				for k := 0; k < 4; k++ {
					nextI, nextJ := I+di[k], J+dj[k]
					if nextI < 0 || h <= nextI || nextJ < 0 || w <= nextJ {
						continue
					}
					if checked[nextI][nextJ] {
						continue
					}
					if s[I][J] != s[nextI][nextJ] {
						que = append(que, [2]int{nextI, nextJ})
					}
				}
			}

			ans += black * white
		}
	}

	puts(ans)
}
