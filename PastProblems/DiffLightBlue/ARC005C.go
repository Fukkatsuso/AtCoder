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

type Point struct {
	i, j int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const inf = 1 << 50

	h, w := nextInt(), nextInt()
	c := make([]string, h)
	var si, sj, gi, gj int
	for i := range c {
		c[i] = next()
		for j := range c[i] {
			if c[i][j] == 's' {
				si, sj = i, j
			} else if c[i][j] == 'g' {
				gi, gj = i, j
			}
		}
	}

	dist := make([][][3]int, h)
	for i := range dist {
		dist[i] = make([][3]int, w)
		for j := range dist[i] {
			dist[i][j][0], dist[i][j][1], dist[i][j][2] = inf, inf, inf
		}
	}

	que := make([]Point, 0)
	dist[si][sj][0] = 0
	que = append(que, Point{si, sj})
	di := []int{1, 0, -1, 0}
	dj := []int{0, 1, 0, -1}
	for len(que) > 0 {
		p := que[0]
		for k := 0; k < 4; k++ {
			nextI, nextJ := p.i+di[k], p.j+dj[k]
			if nextI < 0 || h <= nextI || nextJ < 0 || w <= nextJ {
				continue
			}
			for x := 0; x < 2; x++ {
				if c[nextI][nextJ] == '.' || c[nextI][nextJ] == 'g' {
					for i := 0; i < 3; i++ {
						if dist[p.i][p.j][i]+1 < dist[nextI][nextJ][i] {
							dist[nextI][nextJ][i] = dist[p.i][p.j][i] + 1
							que = append(que, Point{nextI, nextJ})
						}
					}
				} else if c[nextI][nextJ] == '#' {
					for i := 0; i < 2; i++ {
						if dist[p.i][p.j][i]+1 < dist[nextI][nextJ][i+1] {
							dist[nextI][nextJ][i+1] = dist[p.i][p.j][i] + 1
							que = append(que, Point{nextI, nextJ})
						}
					}
				}
			}
		}
		que = que[1:]
	}

	if d := dist[gi][gj]; d[0] < inf || d[1] < inf || d[2] < inf {
		puts("YES")
	} else {
		puts("NO")
	}
}
