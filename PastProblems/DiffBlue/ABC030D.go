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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, a := nextInt(), nextInt()-1
	k := next()
	b := make([]int, n)
	for i := range b {
		b[i] = nextInt() - 1
	}

	// ループの開始地点
	loopBegin := a
	for vis := make([]bool, n); !vis[loopBegin]; {
		vis[loopBegin] = true
		loopBegin = b[loopBegin]
	}

	// 単語aからループに入るまでの経由単語
	before := make([]int, 0)
	for x := a; x != loopBegin; {
		before = append(before, x)
		x = b[x]
	}
	// ループ
	loop := make([]int, 0)
	for x := loopBegin; ; {
		loop = append(loop, x)
		x = b[x]
		if x == loopBegin {
			break
		}
	}

	notLoop := true
	x := 0
	for i := 0; notLoop && i < len(k); i++ {
		x = x*10 + int(k[i]-'0')
		notLoop = x <= len(before)
	}

	// kステップ目に意味を知りたい単語
	var w int
	// ループに入らない場合
	if notLoop {
		// aからx(=k,1-indexed)ステップ目の単語
		w = before[x-1]
	} else {
		L, kModL := len(loop), 0
		for i := 0; i < len(k); i++ {
			kModL = 10*kModL + int(k[i]-'0')
			kModL %= L
		}
		// loopBeginの単語から，(k-len(before))%L (1-indexded)ステップ目の単語
		w = loop[(kModL-len(before)%L+L-1)%L]
	}

	// wの意味を知るために調べる単語が答え
	puts(b[w] + 1)
}
