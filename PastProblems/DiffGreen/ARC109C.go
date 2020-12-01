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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func win(a, b byte) byte {
	// hand[x]: xを出して勝てる相手の手
	hand := map[byte]byte{
		'R': 'S',
		'S': 'P',
		'P': 'R',
	}
	if hand[b] == a {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := getInt(), getInt()
	s := gets()

	for i := 0; i < k; i++ {
		t := s + s
		s = ""
		for j := 0; j < n*2; j += 2 {
			s += string(win(t[j], t[j+1]))
		}
	}

	puts(string(s[0]))
}
