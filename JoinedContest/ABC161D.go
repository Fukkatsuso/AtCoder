// NG

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

	k := nextInt()

	s := make([][10][10]int, 20)
	sum := 0
	for i := 1; i <= 9; i++ {
		s[1][i][i] = 1
		if sum++; sum == k {
			puts(i)
			return
		}
	}
	for m := 2; m <= 20; m++ {
		for i := 1; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				if j == 0 {
					s[m][i][j] = s[m-1][i][1] + s[m-1][i][0]
				} else if j == 9 {
					s[m][i][j] = s[m-1][i][8] + s[m-1][i][9]
				} else {
					s[m][i][j] = s[m-1][i][j-1] + s[m-1][i][j] + s[m-1][i][j+1]
				}

				sum += s[m][i][j]
				if sum+s[m][i][j] >= k {
					puts(m, i, j)

					return
				}
			}
		}
	}
}
