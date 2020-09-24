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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, d := nextInt(), nextInt()
	x, y := abs(nextInt()), abs(nextInt())

	if x%d != 0 || y%d != 0 {
		puts(0)
		return
	}
	x /= d
	y /= d

	if r := (n - x - y); r < 0 || r%2 == 1 {
		puts(0)
		return
	}

	nck := make([][1001]float64, 1001)
	nck[0][0] = 1
	for i := 1; i < 1001; i++ {
		nck[i][0] = nck[i-1][0] / 2.0
	}
	for i := 1; i < 1001; i++ {
		for j := 1; j < 1001; j++ {
			nck[i][j] = (nck[i-1][j-1] + nck[i-1][j]) / 2.0
		}
	}

	ans := 0.0
	for k := 0; k <= n; k++ {
		if k < x || (k+x)%2 != 0 || (n-k) < y || (n-k+y)%2 != 0 {
			continue
		}
		// n回の移動のうちk回，X軸方向に移動する確率
		// k回移動してxに到達する(xの方向に(k+x)/2回の移動が必要)確率
		// Y軸方向にn-k回移動してyに到達する確率
		// の積
		add := nck[n][k] * nck[k][(k+x)/2] * nck[n-k][(n-k+y)/2]
		ans += add
	}

	puts(ans)
}
