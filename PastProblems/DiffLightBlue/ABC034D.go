// 嘘解法
// 一応AC

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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
}

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

// 混合後の体積，濃度
func mix(w1, p1, w2, p2 float64) (float64, float64) {
	return w1 + w2, (p1*w1 + p2*w2) / (w1 + w2)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt2()
	w, p := make([]float64, n), make([]float64, n)
	for i := 0; i < n; i++ {
		w[i], p[i] = nextFloat64(), nextFloat64()/100.0
	}

	ans, weight := 0.0, 0.0
	selected := make([]bool, n)
	for i := 0; i < k; i++ {
		var sel int
		nextW, nextP := 0.0, 0.0
		for j := 0; j < n; j++ {
			if selected[j] {
				continue
			}
			nw, np := mix(weight, ans, w[j], p[j])
			if np > nextP {
				sel = j
				nextW = nw
				nextP = np
			}
		}
		selected[sel] = true
		ans, weight = nextP, nextW
	}

	puts(ans * 100)
}
