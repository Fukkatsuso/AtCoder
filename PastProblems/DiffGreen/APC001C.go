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
	fmt.Println(a...)
}

func vacantExists(left, right int, sl, sr string) bool {
	if (right-left)%2 == 0 && sl != sr {
		return true
	} else if (right-left)%2 == 1 && sl == sr {
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)

	n := nextInt()

	left, right := 0, n-1
	var sl, sr, sm string
	puts(left)
	if sl = next(); sl == "Vacant" {
		return
	}
	puts(right)
	if sr = next(); sr == "Vacant" {
		return
	}

	for cnt := 2; cnt <= 20; {
		mid := (left + right) / 2
		puts(mid)
		cnt++
		if sm = next(); sm == "Vacant" {
			return
		}
		if vacantExists(left, mid, sl, sm) {
			right = mid
			sr = sm
		} else if vacantExists(mid, right, sm, sr) {
			left = mid
			sl = sm
		}
	}
}
