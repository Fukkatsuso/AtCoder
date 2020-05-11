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

func dir(deg int) string {
	deg *= 10
	if deg < 1125 || 34875 <= deg {
		return "N"
	}
	d := []string{"NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}
	return d[(deg-1125)/(36000/16)]
}

func w(dis int) int {
	// 秒速変換後，少数第2位を四捨五入した値の10倍
	v := ((dis*100)/60 + 5) / 10
	switch {
	case v <= 2:
		return 0
	case v <= 15:
		return 1
	case v <= 33:
		return 2
	case v <= 54:
		return 3
	case v <= 79:
		return 4
	case v <= 107:
		return 5
	case v <= 138:
		return 6
	case v <= 171:
		return 7
	case v <= 207:
		return 8
	case v <= 244:
		return 9
	case v <= 284:
		return 10
	case v <= 326:
		return 11
	}
	return 12
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	deg, dis := nextInt(), nextInt()

	ansW := w(dis)
	ansDir := dir(deg)
	if ansW == 0 {
		ansDir = "C"
	}
	puts(ansDir, ansW)
}
