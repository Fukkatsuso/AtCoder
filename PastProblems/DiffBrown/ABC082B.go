// []rune型としてソートした文字列同士を比較するほうが簡単
// 一応AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func main() {
	sc.Split(bufio.ScanWords)
	s, t := next(), next()
	sTable := make([]int, 26)
	tTable := make([]int, 26)
	for i := 0; i < len(s); i++ {
		sTable[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		tTable[t[i]-'a']++
	}

	si, ti := 0, 25
	for si <= ti {
		for si < 26 && sTable[si] == 0 {
			si++
		}
		for ti >= 0 && tTable[ti] == 0 {
			ti--
		}
		if si == 26 {
			fmt.Println("Yes")
			return
		}
		if ti == -1 {
			fmt.Println("No")
			return
		}
		if si < ti {
			fmt.Println("Yes")
			return
		}
		if si == ti && sTable[si] == tTable[ti] {
			fmt.Println("No")
			return
		}
		sTable[si]--
		tTable[ti]--
	}
	fmt.Println("No")
}
