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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w, _ := nextInt(), nextInt(), nextInt()
	s := make([]string, h)
	for i := range s {
		s[i] = next()
	}

	cake := make([][]int, h)
	for i := range cake {
		cake[i] = make([]int, w)
	}

	id := 0
	for top := 0; top < h; {
		// 1段目の状態で初期化
		exist := make([]bool, w)
		for i := 0; i < w; i++ {
			exist[i] = s[top][i] == '#'
		}
		// ヨコに切れるところまでbottomを下げる
		bottom := top + 1
		for ok := true; ok && bottom < h; {
			line := make([]bool, w)
			for i := 0; i < w; i++ {
				line[i] = s[bottom][i] == '#'
				if line[i] && exist[i] {
					ok = false
					break
				}
			}
			if ok {
				for i := 0; i < w; i++ {
					exist[i] = exist[i] || line[i]
				}
				bottom++
			}
		}

		// 左から順にidを増やしながら切る
		cnt := 0
		for l := 0; l < w; {
			if exist[l] {
				cnt++
			}
			r := l
			for ; r+1 < w; r++ {
				if exist[r+1] {
					cnt++
					if cnt > 1 {
						break
					}
				}
			}
			id++
			for i := l; i <= r; i++ {
				for j := top; j < bottom; j++ {
					cake[j][i] = id
				}
			}
			l = r + 1
		}

		top = bottom
	}

	for i := range cake {
		for j := 0; j < w; j++ {
			if j < w-1 {
				putf("%d ", cake[i][j])
			} else {
				putf("%d\n", cake[i][j])
			}
		}
	}
}
