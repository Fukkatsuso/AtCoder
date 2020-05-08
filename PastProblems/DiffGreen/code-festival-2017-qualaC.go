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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func YesOrNo(b bool) {
	if b {
		puts("Yes")
	} else {
		puts("No")
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w := nextInt2()
	cnt := map[byte]int{}
	for i := 0; i < h; i++ {
		a := next()
		for j := 0; j < w; j++ {
			cnt[a[j]]++
		}
	}

	odd := 0
	set2, set4 := 0, 0
	for _, v := range cnt {
		if v%2 == 1 {
			odd++
		}
		set2 += v / 2
		set4 += v / 4
	}

	if h%2 == 0 && w%2 == 0 {
		// 各文字の個数が4の倍数ならOK
		for _, v := range cnt {
			if v%4 != 0 {
				puts("No")
				return
			}
		}
		puts("Yes")
	} else if h%2 == 1 && w%2 == 1 {
		// 文字数奇数個の文字が1つで
		// 4つセットが一定数以上作れるならOK
		YesOrNo(odd == 1 && 4*set4 >= (h*w-(h+w-1)))
	} else if h%2 == 0 {
		YesOrNo(odd == 0 && 4*set4 >= (h*(w-1)))
	} else {
		YesOrNo(odd == 0 && 4*set4 >= (w*(h-1)))
	}
}
